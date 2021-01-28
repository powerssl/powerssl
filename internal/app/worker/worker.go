package worker

import (
	"context"
	"os"

	"github.com/go-kit/kit/log"
	temporalworker "go.temporal.io/sdk/worker"
	"golang.org/x/sync/errgroup"

	"powerssl.dev/powerssl/internal/pkg/auth"
	"powerssl.dev/powerssl/internal/pkg/temporal"
	"powerssl.dev/powerssl/internal/pkg/temporal/activity"
	temporalclient "powerssl.dev/powerssl/internal/pkg/temporal/client"
	"powerssl.dev/powerssl/internal/pkg/temporal/workflow"
	"powerssl.dev/powerssl/internal/pkg/tracing"
	"powerssl.dev/powerssl/internal/pkg/transport"
	"powerssl.dev/powerssl/internal/pkg/util"
	apiserverclient "powerssl.dev/powerssl/pkg/apiserver/client"
)

const component = "powerssl-controller"

func Run(cfg *Config) {
	logger := util.NewLogger(os.Stdout)

	util.ValidateConfig(cfg, logger)

	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return util.InterruptHandler(ctx, logger)
	})

	tracer, closer, err := tracing.Init(component, cfg.Tracer, log.With(logger, "component", "tracing"))
	if err != nil {
		logger.Log("component", "tracing", "err", err)
		os.Exit(1)
	}
	defer closer.Close()

	var client *apiserverclient.GRPCClient
	{
		token, err := auth.NewServiceToken(cfg.AuthToken)
		if err != nil {
			logger.Log("err", err)
			os.Exit(1)
		}
		if client, err = apiserverclient.NewGRPCClient(ctx, cfg.APIServerClientConfig, token, logger, tracer); err != nil {
			logger.Log("transport", "gRPC", "during", "Connect", "err", err)
			os.Exit(1)
		}
	}
	var _ = client

	var temporalClient temporalclient.Client
	{
		var err error
		if temporalClient, err = temporalclient.NewClient(temporalclient.Config{
			CAFile:    cfg.VaultClientConfig.CAFile, // TODO Wrong cfg path
			HostPort:  cfg.TemporalClientConfig.HostPort,
			Namespace: cfg.TemporalClientConfig.Namespace,
		}, nil, tracer); err != nil {
			logger.Log("err", err)
			os.Exit(1)
		}
		defer temporalClient.Close()
	}

	if cfg.MetricsAddr != "" {
		g.Go(func() error {
			return transport.ServeMetrics(ctx, cfg.MetricsAddr, log.With(logger, "component", "metrics"))
		})
	}

	g.Go(func() error {
		worker := temporalworker.New(temporalClient, temporal.TaskQueue, temporalworker.Options{})
		worker.RegisterActivity(activity.CreateACMEAccount)
		worker.RegisterActivity(activity.UpdateAccount)
		worker.RegisterWorkflow(workflow.CreateAccount)
		if err := worker.Start(); err != nil {
			return err
		}
		select {
		case <-ctx.Done():
			worker.Stop()
			return ctx.Err()
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
