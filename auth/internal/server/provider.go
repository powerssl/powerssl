package server

import (
	"context"

	"github.com/google/wire"

	"powerssl.dev/auth/internal/oauth2"
	"powerssl.dev/common/log"
)

var Provider = wire.NewSet(
	Provide,
)

type F func() error

func Provide(ctx context.Context, cfg *Config, logger log.Logger, auth2 *oauth2.OAuth2) F {
	logger = logger.With("component", "server")

	return func() error { return ServeHTTP(ctx, cfg, logger, auth2) }
}
