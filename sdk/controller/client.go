package controller // import "powerssl.dev/sdk/controller"

import (
	"context"

	stdopentracing "github.com/opentracing/opentracing-go"

	"powerssl.dev/common/log"
	"powerssl.dev/common/transport"

	"powerssl.dev/sdk/controller/acme"
	acmetransport "powerssl.dev/sdk/controller/acme/transport"
	"powerssl.dev/sdk/controller/integration"
	intregrationtransport "powerssl.dev/sdk/controller/integration/transport"
	"powerssl.dev/sdk/internal"
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
	authSigner := internal.NewSigner(authToken)
	var _ = authSigner
	var _ = conn
	return &GRPCClient{
		ACME:        acmetransport.NewGRPCClient(conn, logger, tracer),
		Integration: intregrationtransport.NewGRPCClient(conn),
	}, nil
}
