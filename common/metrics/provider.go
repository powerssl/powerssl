package metrics

import (
	"context"

	"github.com/google/wire"

	"powerssl.dev/common/log"
)

var Provider = wire.NewSet(
	Provide,
)

type F func() error

func Provide(ctx context.Context, cfg Config, logger log.Logger) F {
	logger = logger.With("component", "metricsServer")
	return func() error {
		if cfg.Addr == "" {
			return nil
		}
		return ServeMetrics(ctx, cfg, logger)
	}
}
