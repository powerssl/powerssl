package apiserver // import "powerssl.dev/sdk/apiserver"

import (
	"context"

	stdopentracing "github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"

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
	opts := []grpc.DialOption{
		grpc.WithUnaryInterceptor(internal.AuthInterceptor()),
		grpc.WithUnaryInterceptor(internal.LoggerInterceptor(logger)),
		grpc.WithUnaryInterceptor(internal.TracerInterceptor(tracer)),
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
