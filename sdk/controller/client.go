package controller // import "powerssl.dev/sdk/controller"

import (
	"context"

	stdopentracing "github.com/opentracing/opentracing-go"
	"go.uber.org/zap"

	apiv1 "powerssl.dev/api/controller/v1"
	"powerssl.dev/common/transport"
)

type GRPCClient struct {
	ACME        apiv1.ACMEServiceClient
	Integration apiv1.IntegrationServiceClient
}

func NewGRPCClient(ctx context.Context, cfg Config, logger *zap.SugaredLogger, tracer stdopentracing.Tracer) (*GRPCClient, error) {
	conn, err := transport.New(ctx, cfg.Client)
	if err != nil {
		return nil, err
	}
	return &GRPCClient{
		ACME:        apiv1.NewACMEServiceClient(conn),
		Integration: apiv1.NewIntegrationServiceClient(conn),
	}, nil
}
