package metrics

import (
	"context"

	"github.com/google/wire"
	"go.uber.org/zap"
)

var Provider = wire.NewSet(
	Provide,
)

type F func() error

func Provide(ctx context.Context, cfg *Config, logger *zap.SugaredLogger) F {
	logger = logger.With("component", "metricsServer")
	return func() error {
		if cfg.Addr == "" {
			return nil
		}
		return ServeMetrics(ctx, cfg, logger)
	}
}
