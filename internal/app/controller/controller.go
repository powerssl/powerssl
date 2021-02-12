package controller

import (
	"context"
	"io"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics/prometheus"
	"github.com/opentracing/opentracing-go"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"go.temporal.io/sdk/worker"
	"golang.org/x/sync/errgroup"

	"powerssl.dev/powerssl/internal/app/controller/acme"
	"powerssl.dev/powerssl/internal/app/controller/integration"
	"powerssl.dev/powerssl/internal/pkg/apiserver"
	"powerssl.dev/powerssl/internal/pkg/temporal"
	"powerssl.dev/powerssl/internal/pkg/temporal/activity"
	temporalclient "powerssl.dev/powerssl/internal/pkg/temporal/client"
	"powerssl.dev/powerssl/internal/pkg/tracing"
	"powerssl.dev/powerssl/internal/pkg/transport"
	"powerssl.dev/powerssl/internal/pkg/util"
	"powerssl.dev/powerssl/internal/pkg/vault"
	apiserverclient "powerssl.dev/powerssl/pkg/apiserver/client"
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
		return transport.ServeGRPC(ctx, &cfg.ServerConfig, log.With(logger, "transport", "gRPC"), []transport.Service{
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
		w.RegisterActivity(activity.CreateACMEAccount)
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
