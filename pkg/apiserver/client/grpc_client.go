package client

import (
	"github.com/go-kit/kit/log"
	stdopentracing "github.com/opentracing/opentracing-go"

	acmeaccountmeta "powerssl.io/pkg/apiserver/acmeaccount/meta"
	acmeaccounttransport "powerssl.io/pkg/apiserver/acmeaccount/transport"
	acmeservermeta "powerssl.io/pkg/apiserver/acmeserver/meta"
	acmeservertransport "powerssl.io/pkg/apiserver/acmeserver/transport"
	certificatemeta "powerssl.io/pkg/apiserver/certificate/meta"
	certificatetransport "powerssl.io/pkg/apiserver/certificate/transport"
	"powerssl.io/pkg/util"
	"powerssl.io/pkg/util/auth"
)

type GRPCClient struct {
	ACMEAccount acmeaccountmeta.Service
	ACMEServer  acmeservermeta.Service
	Certificate certificatemeta.Service
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
	}, nil
}
