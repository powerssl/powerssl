package apiserver // import "powerssl.dev/sdk/apiserver"

import (
	"context"

	stdopentracing "github.com/opentracing/opentracing-go"
	"go.uber.org/zap"

	apiv1 "powerssl.dev/api/apiserver/v1"
	"powerssl.dev/common/transport"

	"powerssl.dev/sdk/internal"
)

type Client struct {
	ACMEAccount apiv1.ACMEAccountServiceClient
	ACMEServer  apiv1.ACMEServerServiceClient
	Certificate apiv1.CertificateServiceClient
	User        apiv1.UserServiceClient
}

func NewClient(ctx context.Context, cfg Config, logger *zap.SugaredLogger, tracer stdopentracing.Tracer) (*Client, error) {
	conn, err := transport.New(ctx, cfg.Client)
	if err != nil {
		return nil, err
	}
	_ = internal.NewSigner(cfg.AuthToken)
	return &Client{
		ACMEAccount: apiv1.NewACMEAccountServiceClient(conn),
		ACMEServer:  apiv1.NewACMEServerServiceClient(conn),
		Certificate: apiv1.NewCertificateServiceClient(conn),
		User:        apiv1.NewUserServiceClient(conn),
	}, nil
}
