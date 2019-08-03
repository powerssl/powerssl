package signer

import (
	"context"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"golang.org/x/sync/errgroup"

	"powerssl.dev/powerssl/internal/pkg/tracing"
	"powerssl.dev/powerssl/internal/pkg/transport"
	"powerssl.dev/powerssl/internal/pkg/util"
)

const component = "powerssl-signer"

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

	var _ = tracer // TODO

	duration := prometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "powerssl_io",
		Subsystem: "signer",
		Name:      "request_duration_seconds",
		Help:      "Request duration in seconds.",
	}, []string{"method", "success"})

	var _ = duration // TODO

	var services []transport.Service

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
