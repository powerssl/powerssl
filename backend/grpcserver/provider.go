package grpcserver

import (
	"context"

	"github.com/google/wire"

	"powerssl.dev/common/log"
)

var Provider = wire.NewSet(
	Provide,
)

type F func() error

type Register func(srv *Server)

func Provide(ctx context.Context, cfg Config, logger log.Logger, f Register) (F, error) {
	logger = logger.With("component", "grpcServer")
	srv, err := New(cfg, logger)
	if err != nil {
		return nil, err
	}
	f(srv)
	return func() error { return srv.Serve(ctx) }, nil
}
