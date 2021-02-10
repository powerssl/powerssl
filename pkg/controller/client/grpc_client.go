package client // import "powerssl.dev/powerssl/pkg/controller/client"

import (
	"context"

	"github.com/go-kit/kit/log"
	stdopentracing "github.com/opentracing/opentracing-go"

	acmetransport "powerssl.dev/powerssl/internal/app/controller/acme/transport"
	intregrationtransport "powerssl.dev/powerssl/internal/app/controller/integration/transport"
	"powerssl.dev/powerssl/internal/pkg/auth"
	"powerssl.dev/powerssl/internal/pkg/transport"
	"powerssl.dev/powerssl/pkg/controller/acme"
	"powerssl.dev/powerssl/pkg/controller/integration"
)

type GRPCClient struct {
	ACME        acme.Service
	Integration integration.Service
}

func NewGRPCClient(ctx context.Context, cfg *transport.ClientConfig, authToken string, logger log.Logger, tracer stdopentracing.Tracer) (*GRPCClient, error) {
	conn, err := transport.NewClientConn(ctx, cfg)
	if err != nil {
		return nil, err
	}
	authSigner := auth.NewSigner(authToken)
	var _ = authSigner
	return &GRPCClient{
		ACME:        acmetransport.NewGRPCClient(conn, logger, tracer),
		Integration: intregrationtransport.NewGRPCClient(conn, logger),
	}, nil
}
