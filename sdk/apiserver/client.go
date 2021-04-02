package apiserver // import "powerssl.dev/sdk/apiserver"

import (
	"context"

	"github.com/go-kit/kit/log"
	stdopentracing "github.com/opentracing/opentracing-go"

	"powerssl.dev/common/transport"

	"powerssl.dev/sdk/apiserver/acmeaccount"
	acmeaccounttransport "powerssl.dev/sdk/apiserver/acmeaccount/transport"
	"powerssl.dev/sdk/apiserver/acmeserver"
	acmeservertransport "powerssl.dev/sdk/apiserver/acmeserver/transport"
	"powerssl.dev/sdk/apiserver/certificate"
	certificatetransport "powerssl.dev/sdk/apiserver/certificate/transport"
	"powerssl.dev/sdk/apiserver/user"
	usertransport "powerssl.dev/sdk/apiserver/user/transport"
	"powerssl.dev/sdk/internal"
)

type Client struct {
	ACMEAccount acmeaccount.Service
	ACMEServer  acmeserver.Service
	Certificate certificate.Service
	User        user.Service
}

func NewClient(ctx context.Context, cfg *transport.ClientConfig, authToken string, logger log.Logger, tracer stdopentracing.Tracer) (*Client, error) {
	conn, err := transport.NewClientConn(ctx, cfg)
	if err != nil {
		return nil, err
	}
	authSigner := internal.NewSigner(authToken)
	return &Client{
		ACMEAccount: acmeaccounttransport.NewGRPCClient(conn, logger, tracer, authSigner),
		ACMEServer:  acmeservertransport.NewGRPCClient(conn, logger, tracer, authSigner),
		Certificate: certificatetransport.NewGRPCClient(conn, logger, tracer, authSigner),
		User:        usertransport.NewGRPCClient(conn, logger, tracer, authSigner),
	}, nil
}
