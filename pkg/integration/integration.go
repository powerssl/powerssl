package integration // import "powerssl.dev/powerssl/pkg/integration"

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/opentracing/opentracing-go"
	"golang.org/x/sync/errgroup"

	acmetransport "powerssl.dev/powerssl/internal/app/controller/acme/transport" // TODO: Wrong package
	apiv1 "powerssl.dev/powerssl/internal/pkg/controller/api/v1"
	"powerssl.dev/powerssl/pkg/controller/api"
	controllerclient "powerssl.dev/powerssl/pkg/controller/client"
	integrationacme "powerssl.dev/powerssl/pkg/integration/acme"
	// integrationdns "powerssl.dev/powerssl/pkg/integration/dns"
	"powerssl.dev/powerssl/internal/pkg/tracing"
	"powerssl.dev/powerssl/internal/pkg/transport"
	"powerssl.dev/powerssl/internal/pkg/util"
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

func Run(cfg *Config, kind kind, name string, handler interface{}) (err error) {
	_, logger := util.NewZapAndKitLogger()

	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return util.InterruptHandler(ctx, logger)
	})

	var tracer opentracing.Tracer
	{
		var closer io.Closer
		if tracer, closer, err = tracing.Init(fmt.Sprintf("powerssl-integration-%s", kind), cfg.Tracer, log.With(logger, "component", "tracing")); err != nil {
			return err
		}
		defer func() {
			err = closer.Close()
		}()
	}

	var controllerClient *controllerclient.GRPCClient
	{
		if controllerClient, err = controllerclient.NewGRPCClient(ctx, &cfg.ControllerClientConfig, cfg.AuthToken, logger, tracer); err != nil {
			return err
		}
	}

	var integrationhandler Integration
	switch kind {
	case KindACME:
		integrationhandler = integrationacme.New(controllerClient.ACME, handler.(integrationacme.Integration))
	case KindDNS:
		return fmt.Errorf("not yet supported")
		// integrationhandler = integrationdns.New(client.DNS, handler.(integrationdns.Integration))
	}

	if cfg.Metrics.Addr != "" {
		g.Go(func() error {
			return transport.ServeMetrics(ctx, cfg.Metrics.Addr, log.With(logger, "component", "metrics"))
		})
	}

	i := &integration{
		client:  controllerClient,
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
				_ = logger.Log("err", i.run(ctx))
				time.Sleep(time.Second)
			}
		}
	})

	if err = g.Wait(); err != nil {
		switch err.(type) {
		case util.InterruptError:
		default:
			return err
		}
	}
	return nil
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
	_ = i.logger.Log("connected", true)
	for {
		activity, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				_ = i.logger.Log("err", "EOF")
				break
			}
			return err
		}
		apiActivity, err := acmetransport.DecodeGRPCActivity(activity)
		if err != nil {
			return err
		}
		go func() {
			_ = i.handleActivity(ctx, apiActivity)
		}()
	}
	return nil
}

func (i *integration) loggingMiddleware(ctx context.Context, activity *api.Activity) (err error) {
	defer func() {
		_ = i.logger.Log("activity", activity.Token, "name", activity.Name, "err", err)
	}()

	return i.tracingMiddleware(ctx, activity)
}

func (i *integration) tracingMiddleware(ctx context.Context, activity *api.Activity) error {
	//wireContext, err := tracing.WireContextFromJSON(activity.Signature) // TODO: Do not use Signature for Span
	//if err != nil && !strings.Contains(err.Error(), "not found") {
	//	_ = i.logger.Log("activity", activity.Token, "err", err)
	//}
	//activitySpan := opentracing.StartSpan(activity.Name.String(), ext.RPCServerOption(wireContext))
	//defer activitySpan.Finish()
	//activitySpan.SetTag("token", activity.Token)
	//ctx = opentracing.ContextWithSpan(ctx, activitySpan)

	return i.handler.HandleActivity(ctx, activity)
}

func (i *integration) handleActivity(ctx context.Context, activity *api.Activity) error {
	return i.loggingMiddleware(ctx, activity)
}
