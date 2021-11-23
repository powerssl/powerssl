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
	logger = logger.With("component", "server")

	return func() error { return ServeHTTP(ctx, cfg, logger) }
}
