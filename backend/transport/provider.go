package transport

import (
	"context"

	"github.com/google/wire"
	"go.uber.org/zap"
)

var Provider = wire.NewSet(
	ProvideServer,
)

type ServerF func() error
type RegisterF func(srv *Server)

func ProvideServer(ctx context.Context, cfg ServerConfig, logger *zap.SugaredLogger, registerF RegisterF) (ServerF, error) {
	logger = logger.With("component", "grpcServer")
	srv, err := New(cfg, logger)
	if err != nil {
		return nil, err
	}
	registerF(srv)

	return func() error { return srv.Serve(ctx) }, nil
}
