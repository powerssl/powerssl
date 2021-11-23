package server

import (
	"context"

	"github.com/google/wire"
	"google.golang.org/grpc"

	"powerssl.dev/common/log"
)

var Provider = wire.NewSet(
	Provide,
)

type F func() error

func Provide(ctx context.Context, cfg Config, logger log.Logger, conn *grpc.ClientConn) F {
	logger = logger.With("component", "server")

	return func() error { return ServeHTTP(ctx, cfg, logger, conn) }
}
