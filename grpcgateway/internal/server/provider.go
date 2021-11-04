package server

import (
	"context"

	"github.com/google/wire"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var Provider = wire.NewSet(
	Provide,
)

type F func() error

func Provide(ctx context.Context, cfg *Config, logger *zap.SugaredLogger, conn *grpc.ClientConn) F {
	logger = logger.With("component", "server")

	return func() error { return ServeHTTP(ctx, cfg, logger, conn) }
}
