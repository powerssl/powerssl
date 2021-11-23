package controller // import "powerssl.dev/sdk/controller"

import (
	"context"

	stdopentracing "github.com/opentracing/opentracing-go"

	apiv1 "powerssl.dev/api/controller/v1"
	"powerssl.dev/common/log"
	"powerssl.dev/common/transport"
)

type GRPCClient struct {
	ACME        apiv1.ACMEServiceClient
	Integration apiv1.IntegrationServiceClient
}

func NewGRPCClient(ctx context.Context, cfg Config, logger log.Logger, tracer stdopentracing.Tracer) (*GRPCClient, error) {
	conn, err := transport.New(ctx, cfg.Client)
	if err != nil {
		return nil, err
	}
	return &GRPCClient{
		ACME:        apiv1.NewACMEServiceClient(conn),
		Integration: apiv1.NewIntegrationServiceClient(conn),
	}, nil
}
