package server

import (
	"context"

	"github.com/google/wire"
	"go.uber.org/zap"

	"powerssl.dev/auth/internal/oauth2"
)

var Provider = wire.NewSet(
	Provide,
)

type F func() error

func Provide(ctx context.Context, cfg *Config, logger *zap.SugaredLogger, auth2 *oauth2.OAuth2) F {
	logger = logger.With("component", "server")

	return func() error { return ServeHTTP(ctx, cfg, logger, auth2) }
}
