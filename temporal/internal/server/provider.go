package server

import (
	"context"

	"github.com/google/wire"

	"powerssl.dev/common/log"
)

var Provider = wire.NewSet(
	Provide,
)

type F func() error

func Provide(ctx context.Context, cfg *Config, logger log.Logger) F {
	return func() error {
		return Run(ctx, cfg, logger)
	}
}
