package transport // import "powerssl.dev/backend/transport"

import (
	"context"

	"github.com/google/wire"
	"go.uber.org/zap"
)

var Provider = wire.NewSet(
	Provide,
)

type F func() error
type Register func(srv *Server)

func Provide(ctx context.Context, cfg Config, logger *zap.SugaredLogger, f Register) (F, error) {
	logger = logger.With("component", "grpcServer")
	srv, err := New(cfg, logger)
	if err != nil {
		return nil, err
	}
	f(srv)

	return func() error { return srv.Serve(ctx) }, nil
}
