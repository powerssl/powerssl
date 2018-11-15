package integration

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"golang.org/x/sync/errgroup"

	acmetransport "powerssl.io/pkg/controller/acme/transport" // TODO: Wrong package
	"powerssl.io/pkg/controller/api"
	apiv1 "powerssl.io/pkg/controller/api/v1"
	controllerclient "powerssl.io/pkg/controller/client"
	integrationacme "powerssl.io/pkg/integration/acme"
	// integrationdns "powerssl.io/pkg/integration/dns"
	"powerssl.io/pkg/util"
	"powerssl.io/pkg/util/tracing"
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
	client      *controllerclient.GRPCClient
	logger      log.Logger
	kind        kind
	name        string
	handler     Integration
	metricsAddr string
}

func New(addr, certFile, serverNameOverride string, insecure, insecureSkipTLSVerify bool, metricsAddr, tracerImpl string, kind kind, name string, handler interface{}) *integration {
	var logger log.Logger
	{
		logger = util.NewLogger(os.Stdout)
	}

	tracer, closer, err := tracing.Init(fmt.Sprintf("powerssl-integration-%s", kind), tracerImpl, log.With(logger, "component", "tracing"))
	if err != nil {
		logger.Log("component", "tracing", "err", err)
		os.Exit(1)
	}
	// defer closer.Close()
	var _ = closer

	var authToken string
	client, err := controllerclient.NewGRPCClient(addr, certFile, serverNameOverride, insecure, insecureSkipTLSVerify, authToken, logger, tracer)
	if err != nil {
		logger.Log("transport", "gRPC", "err", err)
		os.Exit(1)
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

	return &integration{
		client:      client,
		logger:      logger,
		kind:        kind,
		name:        name,
		handler:     integrationhandler,
		metricsAddr: metricsAddr,
	}
}

func (i *integration) Run() {
	logger := i.logger
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return util.InterruptHandler(ctx, logger)
	})

	if i.metricsAddr != "" {
		g.Go(func() error {
			return util.ServeMetrics(ctx, i.metricsAddr, log.With(logger, "component", "metrics"))
		})
	}

	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				i.logger.Log("err", i.run(ctx))
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
		if err == io.EOF {
			i.logger.Log("err", "EOF")
			break
		}
		if err != nil {
			return err
		}
		{
			activity, err := acmetransport.DecodeGRPCActivity(activity)
			if err != nil {
				return err
			}
			go i.handleActivity(activity)
		}
	}
	return nil
}

func (i *integration) handleActivity(activity *api.Activity) {
	i.logger.Log("activity", activity.Token, "name", activity.Name)
	wireContext, err := tracing.WireContextFromJSON(activity.Signature) // TODO: Do not use Signature for Span
	if err != nil {
		panic(err)
		// TODO
	}
	activitySpan := opentracing.StartSpan(activity.Name.String(), ext.RPCServerOption(wireContext))
	activitySpan.SetTag("token", activity.Token)
	ctx := opentracing.ContextWithSpan(context.Background(), activitySpan)
	err = i.handler.HandleActivity(ctx, activity)
	activitySpan.Finish()
	i.logger.Log("activity", activity.Token, "err", err)
}
