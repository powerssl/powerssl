package client // import "powerssl.dev/sdk/apiserver/client"

import (
	"context"

	"github.com/go-kit/kit/log"
	stdopentracing "github.com/opentracing/opentracing-go"

	"powerssl.dev/common/auth"
	"powerssl.dev/common/transport"
	"powerssl.dev/sdk/apiserver/acmeaccount"
	acmeaccounttransport "powerssl.dev/sdk/apiserver/acmeaccount/transport"
	"powerssl.dev/sdk/apiserver/acmeserver"
	acmeservertransport "powerssl.dev/sdk/apiserver/acmeserver/transport"
	"powerssl.dev/sdk/apiserver/certificate"
	certificatetransport "powerssl.dev/sdk/apiserver/certificate/transport"
	"powerssl.dev/sdk/apiserver/user"
	usertransport "powerssl.dev/sdk/apiserver/user/transport"
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
