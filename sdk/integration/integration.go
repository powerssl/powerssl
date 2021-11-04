package integration // import "powerssl.dev/sdk/integration"

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/opentracing/opentracing-go"
	"golang.org/x/sync/errgroup"

	apiv1 "powerssl.dev/api/controller/v1"
	error2 "powerssl.dev/common/error"
	"powerssl.dev/common/interrupthandler"
	"powerssl.dev/common/log"
	"powerssl.dev/common/metrics"
	"powerssl.dev/common/tracing"
	"powerssl.dev/sdk/controller"
	"powerssl.dev/sdk/controller/api"
	integrationacme "powerssl.dev/sdk/integration/acme"
	// integrationdns "powerssl.dev/sdk/integration/dns"
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
	client  *controller.GRPCClient
	logger  log.Logger
	kind    kind
	name    string
	handler Integration
}

func Run(cfg *Config, kind kind, name string, handler interface{}) (err error) {
	var logger log.Logger
	if logger, err = log.NewLogger(false); err != nil {
		return err
	}
	defer error2.ErrWrapSync(logger, &err)

	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return interrupthandler.InterruptHandler(ctx, logger)
	})

	var tracer opentracing.Tracer
	{
		var closer io.Closer
		if tracer, closer, err = tracing.Init(fmt.Sprintf("powerssl-integration-%s", kind), cfg.Tracer, logger.With("component", "tracing")); err != nil {
			return err
		}
		defer error2.ErrWrapCloser(closer, &err)
	}

	var controllerClient *controller.GRPCClient
	{
		if controllerClient, err = controller.NewGRPCClient(ctx, &cfg.ControllerClientConfig, cfg.AuthToken, logger, tracer); err != nil {
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
			return metrics.ServeMetrics(ctx, cfg.Metrics.Addr, logger.With("component", "metrics"))
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
				logger.Error(i.run(ctx))
				time.Sleep(time.Second)
			}
		}
	})

	if err = g.Wait(); err != nil {
		switch err.(type) {
		case interrupthandler.InterruptError:
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
	i.logger.Info("connected")
	for {
		activity, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				i.logger.Error("EOF")
				break
			}
			return err
		}
		go func() {
			_ = i.handleActivity(ctx, &api.Activity{
				Token: activity.GetToken(),
				Name:  api.ActivityName(activity.GetName()),
			})
		}()
	}
	return nil
}

func (i *integration) loggingMiddleware(ctx context.Context, activity *api.Activity) (err error) {
	defer func() {
		i.logger.Infow("Received activity", "activity", activity.Token, "name", activity.Name, "err", err)
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
