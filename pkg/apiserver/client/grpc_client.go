package client // import "powerssl.io/powerssl/pkg/apiserver/client"

import (
	"context"

	"github.com/go-kit/kit/log"
	stdopentracing "github.com/opentracing/opentracing-go"

	acmeaccounttransport "powerssl.io/powerssl/internal/app/apiserver/acmeaccount/transport"
	acmeservertransport "powerssl.io/powerssl/internal/app/apiserver/acmeserver/transport"
	certificatetransport "powerssl.io/powerssl/internal/app/apiserver/certificate/transport"
	usertransport "powerssl.io/powerssl/internal/app/apiserver/user/transport"
	"powerssl.io/powerssl/internal/pkg/auth"
	"powerssl.io/powerssl/internal/pkg/transport"
	"powerssl.io/powerssl/pkg/apiserver/acmeaccount"
	"powerssl.io/powerssl/pkg/apiserver/acmeserver"
	"powerssl.io/powerssl/pkg/apiserver/certificate"
	"powerssl.io/powerssl/pkg/apiserver/user"
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
