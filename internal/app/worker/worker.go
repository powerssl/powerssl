package worker

import (
	"context"
	"io"

	"github.com/go-kit/kit/log"
	"github.com/opentracing/opentracing-go"
	"go.temporal.io/sdk/worker"
	"golang.org/x/sync/errgroup"

	"powerssl.dev/powerssl/internal/pkg/apiserver"
	"powerssl.dev/powerssl/internal/pkg/temporal"
	"powerssl.dev/powerssl/internal/pkg/temporal/activity"
	temporalclient "powerssl.dev/powerssl/internal/pkg/temporal/client"
	"powerssl.dev/powerssl/internal/pkg/temporal/workflow"
	"powerssl.dev/powerssl/internal/pkg/tracing"
	"powerssl.dev/powerssl/internal/pkg/transport"
	"powerssl.dev/powerssl/internal/pkg/util"
	"powerssl.dev/powerssl/internal/pkg/vault"
	apiserverclient "powerssl.dev/powerssl/pkg/apiserver/client"
)

const component = "powerssl-worker"

func Run(cfg *Config) (err error) {
	zapLogger, logger := util.NewZapAndKitLogger()

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
		worker.EnableVerboseLogging(true)
		backgroundActivityContext := context.Background()
		backgroundActivityContext = apiserver.SetClient(backgroundActivityContext, apiserverClient)
		backgroundActivityContext = vault.SetClient(backgroundActivityContext, vaultClient)
		w := worker.New(temporalClient, temporal.WorkerTaskQueue, worker.Options{
			BackgroundActivityContext: backgroundActivityContext,
		})
		w.RegisterActivity(activity.CreateVaultTransitKey)
		w.RegisterActivity(activity.UpdateAccount)
		w.RegisterWorkflow(workflow.CreateAccount)
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
