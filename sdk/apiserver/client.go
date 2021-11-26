package apiserver // import "powerssl.dev/sdk/apiserver"

import (
	"context"

	"google.golang.org/grpc"

	apiv1 "powerssl.dev/api/apiserver/v1"
	"powerssl.dev/common/log"
	"powerssl.dev/common/telemetry"
	"powerssl.dev/common/transport"

	"powerssl.dev/sdk/internal"
)

type Client struct {
	ACMEAccount apiv1.ACMEAccountServiceClient
	ACMEServer  apiv1.ACMEServerServiceClient
	Certificate apiv1.CertificateServiceClient
	User        apiv1.UserServiceClient
}

func NewClient(ctx context.Context, cfg Config, logger log.Logger, telemetry *telemetry.Telemeter) (*Client, error) {
	opts := []grpc.DialOption{
		grpc.WithUnaryInterceptor(internal.AuthInterceptor()),
		grpc.WithUnaryInterceptor(internal.LoggerInterceptor(logger)),
		grpc.WithUnaryInterceptor(internal.TelemetryInterceptor(telemetry)),
	}
	conn, err := transport.New(ctx, cfg.Client, opts...)
	if err != nil {
		return nil, err
	}
	return &Client{
		ACMEAccount: apiv1.NewACMEAccountServiceClient(conn),
		ACMEServer:  apiv1.NewACMEServerServiceClient(conn),
		Certificate: apiv1.NewCertificateServiceClient(conn),
		User:        apiv1.NewUserServiceClient(conn),
	}, nil
}
