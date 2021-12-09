package server

import (
	"context"

	"github.com/google/wire"
	"google.golang.org/grpc"

	"powerssl.dev/common/log"
	"powerssl.dev/common/transport"
)

var Provider = wire.NewSet(
	Provide,
	transport.NoDialOptions,
	transport.Provider,
)

type F func() error

func Provide(ctx context.Context, cfg Config, logger log.Logger, conn *grpc.ClientConn) F {
	logger = logger.With("component", "server")
	server := New(cfg, logger, conn)
	return func() error {
		return server.ServeHTTP(ctx)
	}
}
