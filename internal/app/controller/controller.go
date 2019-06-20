package controller

import (
	"context"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"golang.org/x/sync/errgroup"

	workflowengine "powerssl.io/powerssl/internal/app/controller/workflow/engine"
	"powerssl.io/powerssl/internal/pkg/auth"
	"powerssl.io/powerssl/internal/pkg/tracing"
	"powerssl.io/powerssl/internal/pkg/transport"
	"powerssl.io/powerssl/internal/pkg/util"
	apiserverclient "powerssl.io/powerssl/pkg/apiserver/client"
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

	if err := g.Wait(); err != nil {
		switch err.(type) {
		case util.InterruptError:
		default:
			logger.Log("err", err)
		}
	}
}
