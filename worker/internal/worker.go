package internal

import (
	"context"
	"io"

	"github.com/opentracing/opentracing-go"
	temporalactivity "go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/worker"
	temporalworkflow "go.temporal.io/sdk/workflow"
	"golang.org/x/sync/errgroup"

	backendapiserver "powerssl.dev/backend/apiserver"
	"powerssl.dev/backend/temporal"
	temporalclient "powerssl.dev/backend/temporal/client"
	"powerssl.dev/backend/vault"
	"powerssl.dev/common"
	"powerssl.dev/common/log"
	"powerssl.dev/common/tracing"
	"powerssl.dev/common/transport"
	"powerssl.dev/sdk/apiserver"
	sharedworkflow "powerssl.dev/workflow"
	sharedactivity "powerssl.dev/workflow/activity"

	"powerssl.dev/worker/internal/activity"
	"powerssl.dev/worker/internal/workflow"
)

const component = "powerssl-worker"

func Run(cfg *Config) (err error) {
	var logger log.Logger
	if logger, err = log.NewLogger(false); err != nil {
		return err
	}
	defer common.ErrWrapSync(logger, &err)

	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return common.InterruptHandler(ctx, logger)
	})

	var tracer opentracing.Tracer
	{
		var closer io.Closer
		if tracer, closer, err = tracing.Init(component, cfg.Tracer, logger.With("component", "tracing")); err != nil {
			return err
		}
		defer common.ErrWrapCloser(closer, &err)
	}

	var apiserverClient *apiserver.Client
	{
		if apiserverClient, err = apiserver.NewClient(ctx, &cfg.APIServerClientConfig, cfg.AuthToken, logger, tracer); err != nil {
			return err
		}
	}

	var temporalClient temporalclient.Client
	{
		var closer io.Closer
		if temporalClient, closer, err = temporalclient.NewClient(cfg.TemporalClientConfig, logger, tracer, component); err != nil {
			return err
		}
		defer func() {
			temporalClient.Close()
			common.ErrWrapCloser(closer, &err)
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
			return transport.ServeMetrics(ctx, cfg.Metrics.Addr, logger.With("component", "metrics"))
		})
	}

	g.Go(func() error {
		worker.EnableVerboseLogging(true)
		backgroundActivityContext := context.Background()
		backgroundActivityContext = backendapiserver.SetClient(backgroundActivityContext, apiserverClient)
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
		case common.InterruptError:
		default:
			return err
		}
	}
	return nil
}
