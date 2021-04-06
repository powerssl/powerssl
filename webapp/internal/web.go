package internal

import (
	"context"

	"golang.org/x/sync/errgroup"

	"powerssl.dev/common"
	"powerssl.dev/common/log"
	"powerssl.dev/common/transport"
)

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

	if cfg.Metrics.Addr != "" {
		g.Go(func() error {
			return transport.ServeMetrics(ctx, cfg.Metrics.Addr, logger.With("component", "metrics"))
		})
	}

	g.Go(func() error {
		return ServeHTTP(ctx, cfg.Addr, cfg.Insecure, cfg.TLS.CertFile, cfg.TLS.PrivateKeyFile, logger.With("component", "http"), cfg.Auth.URI, cfg.APIServer.Addr, cfg.GRPCWeb.URI)
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
