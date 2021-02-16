package internal

import (
	"context"
	"io"

	"github.com/go-kit/kit/log"
	"github.com/opentracing/opentracing-go"
	temporalactivity "go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/worker"
	temporalworkflow "go.temporal.io/sdk/workflow"
	"golang.org/x/sync/errgroup"

	"powerssl.dev/backend/apiserver"
	"powerssl.dev/backend/temporal"
	temporalclient "powerssl.dev/backend/temporal/client"
	"powerssl.dev/backend/vault"
	"powerssl.dev/common/tracing"
	"powerssl.dev/common/transport"
	"powerssl.dev/common/util"
	apiserverclient "powerssl.dev/sdk/apiserver/client"
	"powerssl.dev/worker/internal/activity"
	"powerssl.dev/worker/internal/workflow"
	sharedworkflow "powerssl.dev/workflow"
	sharedactivity "powerssl.dev/workflow/activity"
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
		w.RegisterActivityWithOptions(activity.CreateVaultTransitKey, temporalactivity.RegisterOptions{
			Name: sharedactivity.CreateVaultTransitKey,
		})
		w.RegisterActivityWithOptions(activity.UpdateAccount, temporalactivity.RegisterOptions{
			Name: sharedactivity.UpdateAccount,
		})
		w.RegisterWorkflowWithOptions(workflow.CreateAccount, temporalworkflow.RegisterOptions{
			Name: sharedworkflow.CreateAccount,
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
