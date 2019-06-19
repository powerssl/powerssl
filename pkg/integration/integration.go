package integration

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"golang.org/x/sync/errgroup"

	acmetransport "powerssl.io/powerssl/internal/app/controller/acme/transport" // TODO: Wrong package
	apiv1 "powerssl.io/powerssl/internal/pkg/controller/api/v1"
	"powerssl.io/powerssl/pkg/controller/api"
	controllerclient "powerssl.io/powerssl/pkg/controller/client"
	integrationacme "powerssl.io/powerssl/pkg/integration/acme"
	// integrationdns "powerssl.io/powerssl/pkg/integration/dns"
	"powerssl.io/powerssl/internal/pkg/tracing"
	"powerssl.io/powerssl/internal/pkg/util"
)

type kind string

const (
	KindACME kind = "acme"
	KindDNS  kind = "dns"
)

type Integration interface {
	HandleActivity(ctx context.Context, activity *api.Activity) error
}

type integration struct {
	client  *controllerclient.GRPCClient
	logger  log.Logger
	kind    kind
	name    string
	handler Integration
}

func Run(cfg *Config, kind kind, name string, handler interface{}) {
	logger := util.NewLogger(os.Stdout)

	util.ValidateConfig(cfg, logger)

	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return util.InterruptHandler(ctx, logger)
	})

	tracer, closer, err := tracing.Init(fmt.Sprintf("powerssl-integration-%s", kind), cfg.Tracer, log.With(logger, "component", "tracing"))
	if err != nil {
		logger.Log("component", "tracing", "err", err)
		os.Exit(1)
	}
	defer closer.Close()

	var client *controllerclient.GRPCClient
	{
		var err error
		if client, err = controllerclient.NewGRPCClient(ctx, cfg.ControllerClientConfig, cfg.AuthToken, logger, tracer); err != nil {
			logger.Log("transport", "gRPC", "during", "Connect", "err", err)
			os.Exit(1)
		}
	}

	var integrationhandler Integration
	switch kind {
	case KindACME:
		integrationhandler = integrationacme.New(client.ACME, handler.(integrationacme.Integration))
	case KindDNS:
		logger.Log("kind", "DNS", "err", "Not yet supported")
		os.Exit(1)
		// integrationhandler = integrationdns.New(client.DNS, handler.(integrationdns.Integration))
	}

	if cfg.MetricsAddr != "" {
		g.Go(func() error {
			return util.ServeMetrics(ctx, cfg.MetricsAddr, log.With(logger, "component", "metrics"))
		})
	}

	i := &integration{
		client:  client,
		logger:  logger,
		kind:    kind,
		name:    name,
		handler: integrationhandler,
	}

	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				logger.Log("err", i.run(ctx))
				time.Sleep(time.Second)
			}
		}
	})

	if err := g.Wait(); err != nil {
		switch err.(type) {
		case util.InterruptError:
		default:
			logger.Log("err", err)
		}
	}
}

func (i *integration) run(ctx context.Context) error {
	var kind apiv1.IntegrationKind
	switch i.kind {
	case KindACME:
		kind = apiv1.IntegrationKind_ACME
	case KindDNS:
		kind = apiv1.IntegrationKind_DNS
	}

	stream, err := i.client.Integration.Register(ctx, &apiv1.RegisterIntegrationRequest{
		Kind: kind,
		Name: i.name,
	})
	if err != nil {
		return err
	}
	i.logger.Log("connected", true)
	for {
		activity, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				i.logger.Log("err", "EOF")
				break
			}
			return err
		}
		apiActivity, err := acmetransport.DecodeGRPCActivity(activity)
		if err != nil {
			return err
		}
		go i.handleActivity(ctx, apiActivity)
	}
	return nil
}

func (i *integration) loggingMiddleware(ctx context.Context, activity *api.Activity) (err error) {
	defer func() {
		i.logger.Log("activity", activity.Token, "name", activity.Name, "err", err)
	}()

	return i.tracingMiddleware(ctx, activity)
}

func (i *integration) tracingMiddleware(ctx context.Context, activity *api.Activity) error {
	wireContext, err := tracing.WireContextFromJSON(activity.Signature) // TODO: Do not use Signature for Span
	if err != nil && !strings.Contains(err.Error(), "not found") {
		i.logger.Log("activity", activity.Token, "err", err)
	}
	activitySpan := opentracing.StartSpan(activity.Name.String(), ext.RPCServerOption(wireContext))
	defer activitySpan.Finish()
	activitySpan.SetTag("token", activity.Token)
	ctx = opentracing.ContextWithSpan(ctx, activitySpan)

	return i.handler.HandleActivity(ctx, activity)
}

func (i *integration) handleActivity(ctx context.Context, activity *api.Activity) error {
	return i.loggingMiddleware(ctx, activity)
}
