package internal

import (
	"context"
	"io"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics/prometheus"
	"github.com/opentracing/opentracing-go"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	temporalactivity "go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/worker"
	"golang.org/x/sync/errgroup"

	"powerssl.dev/backend/apiserver"
	"powerssl.dev/backend/temporal"
	temporalclient "powerssl.dev/backend/temporal/client"
	backendtransport "powerssl.dev/backend/transport"
	"powerssl.dev/backend/vault"
	"powerssl.dev/common/tracing"
	"powerssl.dev/common/transport"
	"powerssl.dev/common/util"
	"powerssl.dev/controller/internal/acme"
	"powerssl.dev/controller/internal/integration"
	apiserverclient "powerssl.dev/sdk/apiserver/client"
	"powerssl.dev/workflow/activity"
	sharedactivity "powerssl.dev/workflow/activity"
)

const component = "powerssl-controller"

func Run(cfg *Config) (err error) {
	zapLogger, logger := util.NewZapAndKitLogger()

	cfg.ServerConfig.VaultToken = cfg.VaultClientConfig.Token
	cfg.ServerConfig.VaultURL = cfg.VaultClientConfig.URL
	cfg.ServerConfig.VaultRole = component

	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return util.InterruptHandler(ctx, logger)
	})

	var tracer opentracing.Tracer
	{
		var closer io.Closer
		if tracer, closer, err = tracing.Init(component, cfg.Tracer, log.With(logger, "component", "tracing")); err != nil {
			return err
		}
		defer func() {
			err = closer.Close()
		}()
	}

	duration := prometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "powerssl_io",
		Subsystem: "controller",
		Name:      "request_duration_seconds",
		Help:      "Request duration in seconds.",
	}, []string{"method", "success"})

	var apiserverClient *apiserverclient.GRPCClient
	{
		if apiserverClient, err = apiserverclient.NewGRPCClient(ctx, &cfg.APIServerClientConfig, cfg.AuthToken, logger, tracer); err != nil {
			return err
		}
	}

	var temporalClient temporalclient.Client
	{
		var closer io.Closer
		if temporalClient, closer, err = temporalclient.NewClient(cfg.TemporalClientConfig, zapLogger, tracer, component); err != nil {
			return err
		}
		defer func() {
			temporalClient.Close()
			err = closer.Close()
		}()
	}

	var vaultClient *vault.Client
	{
		if vaultClient, err = vault.New(cfg.VaultClientConfig); err != nil {
			return err
		}
	}

	if cfg.Metrics.Addr != "" {
		g.Go(func() error {
			return transport.ServeMetrics(ctx, cfg.Metrics.Addr, log.With(logger, "component", "metrics"))
		})
	}

	g.Go(func() error {
		return backendtransport.ServeGRPC(ctx, &cfg.ServerConfig, log.With(logger, "transport", "gRPC"), []backendtransport.Service{
			acme.New(logger, tracer, duration, temporalClient),
			integration.New(ctx, logger, duration), // TODO: tracing
		})
	})

	g.Go(func() error {
		worker.EnableVerboseLogging(true)
		backgroundActivityContext := context.Background()
		backgroundActivityContext = apiserver.SetClient(backgroundActivityContext, apiserverClient)
		backgroundActivityContext = vault.SetClient(backgroundActivityContext, vaultClient)
		w := worker.New(temporalClient, temporal.ControllerTaskQueue, worker.Options{
			BackgroundActivityContext: backgroundActivityContext,
		})
		w.RegisterActivityWithOptions(activity.CreateACMEAccount, temporalactivity.RegisterOptions{
			Name: sharedactivity.CreateACMEAccount,
		})
		if err = w.Start(); err != nil {
			return err
		}
		select {
		case <-ctx.Done():
			w.Stop()
			return ctx.Err()
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