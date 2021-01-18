package controller

import (
	"context"
	"os"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/uber-go/tally"
	temporalclient "go.temporal.io/sdk/client"
	temporalworker "go.temporal.io/sdk/worker"
	"golang.org/x/sync/errgroup"

	"powerssl.dev/powerssl/internal/app/controller/activity"
	"powerssl.dev/powerssl/internal/app/controller/workflow"
	workflowengine "powerssl.dev/powerssl/internal/app/controller/workflow/engine"
	"powerssl.dev/powerssl/internal/pkg/auth"
	"powerssl.dev/powerssl/internal/pkg/temporal"
	"powerssl.dev/powerssl/internal/pkg/tracing"
	"powerssl.dev/powerssl/internal/pkg/transport"
	"powerssl.dev/powerssl/internal/pkg/util"
	apiserverclient "powerssl.dev/powerssl/pkg/apiserver/client"
)

const component = "powerssl-controller"

func Run(cfg *Config) {
	logger := util.NewLogger(os.Stdout)

	cfg.ServerConfig.VaultRole = component
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

	duration := prometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "powerssl_io",
		Subsystem: "controller",
		Name:      "request_duration_seconds",
		Help:      "Request duration in seconds.",
	}, []string{"method", "success"})

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

	var temporalClient temporalclient.Client
	{
		scope, _ := tally.NewRootScope(tally.ScopeOptions{Separator: "_"}, time.Second)
		var err error
		if temporalClient, err = temporalclient.NewClient(temporalclient.Options{
			HostPort:     cfg.TemporalClientConfig.HostPort,
			Namespace:    cfg.TemporalClientConfig.Namespace,
			MetricsScope: scope,
			Tracer:       tracer,
			//Logger:            logger,
			//Identity:          "",
			//ConnectionOptions: temporalclient.ConnectionOptions{},
		}); err != nil {
			logger.Log("err", err)
			os.Exit(1)
		}
		defer temporalClient.Close()
	}

	engine := workflowengine.New()

	services, err := makeServices(logger, tracer, duration, client, cfg.JWKSURL)
	if err != nil {
		logger.Log("err", err)
		os.Exit(1)
	}

	g.Go(func() error {
		return engine.Run(ctx)
	})

	if cfg.MetricsAddr != "" {
		g.Go(func() error {
			return transport.ServeMetrics(ctx, cfg.MetricsAddr, log.With(logger, "component", "metrics"))
		})
	}

	g.Go(func() error {
		return transport.ServeGRPC(ctx, cfg.ServerConfig, log.With(logger, "transport", "gRPC"), services)
	})

	g.Go(func() error {
		worker := temporalworker.New(temporalClient, temporal.TaskQueue, temporalworker.Options{})
		worker.RegisterActivity(activity.CreateAccount)
		worker.RegisterWorkflow(workflow.CreateAccount)
		interruptCh := make(chan interface{}, 1)
		go func() {
			s := <-ctx.Done()
			interruptCh <- s
			close(interruptCh)
		}()
		return worker.Run(interruptCh)
	})

	if err := g.Wait(); err != nil {
		switch err.(type) {
		case util.InterruptError:
		default:
			logger.Log("err", err)
		}
	}
}
