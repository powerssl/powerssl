package client // import "powerssl.dev/powerssl/pkg/apiserver/client"

import (
	"context"

	"github.com/go-kit/kit/log"
	stdopentracing "github.com/opentracing/opentracing-go"

	acmeaccounttransport "powerssl.dev/powerssl/internal/app/apiserver/acmeaccount/transport"
	acmeservertransport "powerssl.dev/powerssl/internal/app/apiserver/acmeserver/transport"
	certificatetransport "powerssl.dev/powerssl/internal/app/apiserver/certificate/transport"
	usertransport "powerssl.dev/powerssl/internal/app/apiserver/user/transport"
	"powerssl.dev/powerssl/internal/pkg/auth"
	"powerssl.dev/powerssl/internal/pkg/transport"
	"powerssl.dev/powerssl/pkg/apiserver/acmeaccount"
	"powerssl.dev/powerssl/pkg/apiserver/acmeserver"
	"powerssl.dev/powerssl/pkg/apiserver/certificate"
	"powerssl.dev/powerssl/pkg/apiserver/user"
)

type GRPCClient struct {
	ACMEAccount acmeaccount.Service
	ACMEServer  acmeserver.Service
	Certificate certificate.Service
	User        user.Service
}

func NewGRPCClient(ctx context.Context, cfg *transport.ClientConfig, authToken string, logger log.Logger, tracer stdopentracing.Tracer) (*GRPCClient, error) {
	conn, err := transport.NewClientConn(ctx, cfg)
	if err != nil {
		return nil, err
	}
	authSigner := auth.NewSigner(authToken)
	return &GRPCClient{
		ACMEAccount: acmeaccounttransport.NewGRPCClient(conn, logger, tracer, authSigner),
		ACMEServer:  acmeservertransport.NewGRPCClient(conn, logger, tracer, authSigner),
		Certificate: certificatetransport.NewGRPCClient(conn, logger, tracer, authSigner),
		User:        usertransport.NewGRPCClient(conn, logger, tracer, authSigner),
	}, nil
}
