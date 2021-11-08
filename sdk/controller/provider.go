package controller // import "powerssl.dev/sdk/controller"

import (
	"context"

	"github.com/google/wire"
	stdopentracing "github.com/opentracing/opentracing-go"
	"go.uber.org/zap"

	"powerssl.dev/common/transport"
)

var Provider = wire.NewSet(
	Provide,
)

type AuthToken string

func Provide(ctx context.Context, cfg *transport.ClientConfig, authToken AuthToken, logger *zap.SugaredLogger, tracer stdopentracing.Tracer) (*GRPCClient, error) {
	client, err := NewGRPCClient(ctx, cfg, string(authToken), logger, tracer)
	if err != nil {
		return nil, err
	}
	return client, nil
}
