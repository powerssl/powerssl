package internal

import (
	"context"

	"github.com/go-kit/kit/log"
	"golang.org/x/sync/errgroup"

	"powerssl.dev/common/transport"
	"powerssl.dev/common/util"
)

func Run(cfg *Config) error {
	_, logger := util.NewZapAndKitLogger()

	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return util.InterruptHandler(ctx, logger)
	})

	if cfg.Metrics.Addr != "" {
		g.Go(func() error {
			return transport.ServeMetrics(ctx, cfg.Metrics.Addr, log.With(logger, "component", "metrics"))
		})
	}

	g.Go(func() error {
		return ServeHTTP(ctx, cfg.Addr, log.With(logger, "component", "http"), cfg.Auth.URI, cfg.APIServer.Addr, cfg.GRPCWeb.URI)
	})

	if err := g.Wait(); err != nil {
		switch err.(type) {
		case util.InterruptError:
		default:
			return err
		}
	}
	return nil
}
