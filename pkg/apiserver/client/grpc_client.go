package client // import "powerssl.io/pkg/apiserver/client"

import (
	"github.com/go-kit/kit/log"
	stdopentracing "github.com/opentracing/opentracing-go"

	acmeaccounttransport "powerssl.io/internal/app/apiserver/acmeaccount/transport"
	acmeservertransport "powerssl.io/internal/app/apiserver/acmeserver/transport"
	certificatetransport "powerssl.io/internal/app/apiserver/certificate/transport"
	usertransport "powerssl.io/internal/app/apiserver/user/transport"
	"powerssl.io/internal/pkg/auth"
	"powerssl.io/internal/pkg/util"
	"powerssl.io/pkg/apiserver/acmeaccount"
	"powerssl.io/pkg/apiserver/acmeserver"
	"powerssl.io/pkg/apiserver/certificate"
	"powerssl.io/pkg/apiserver/user"
)

type GRPCClient struct {
	ACMEAccount acmeaccount.Service
	ACMEServer  acmeserver.Service
	Certificate certificate.Service
	User        user.Service
}

func NewGRPCClient(addr, certFile, serverNameOverride string, insecure, insecureSkipTLSVerify bool, authToken string, logger log.Logger, tracer stdopentracing.Tracer) (*GRPCClient, error) {
	conn, err := util.NewClientConn(addr, certFile, serverNameOverride, insecure, insecureSkipTLSVerify)
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
