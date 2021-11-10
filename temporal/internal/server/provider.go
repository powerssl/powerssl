package server

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
	return func() error {
		return Run(ctx, cfg, logger)
	}
}
