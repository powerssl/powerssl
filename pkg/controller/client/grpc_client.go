package client // import "powerssl.io/powerssl/pkg/controller/client"

import (
	"github.com/go-kit/kit/log"
	stdopentracing "github.com/opentracing/opentracing-go"

	acmetransport "powerssl.io/powerssl/internal/app/controller/acme/transport"
	intregrationtransport "powerssl.io/powerssl/internal/app/controller/integration/transport"
	workflowtransport "powerssl.io/powerssl/internal/app/controller/workflow/transport"
	"powerssl.io/powerssl/internal/pkg/auth"
	"powerssl.io/powerssl/internal/pkg/util"
	"powerssl.io/powerssl/pkg/controller/acme"
	"powerssl.io/powerssl/pkg/controller/integration"
	"powerssl.io/powerssl/pkg/controller/workflow"
)

type GRPCClient struct {
	ACME        acme.Service
	Integration integration.Service
	Workflow    workflow.Service
}

func NewGRPCClient(addr, certFile, serverNameOverride string, insecure, insecureSkipTLSVerify bool, authToken string, logger log.Logger, tracer stdopentracing.Tracer) (*GRPCClient, error) {
	conn, err := util.NewClientConn(addr, certFile, serverNameOverride, insecure, insecureSkipTLSVerify)
	if err != nil {
		return nil, err
	}
	authSigner := auth.NewSigner(authToken)
	return &GRPCClient{
		ACME:        acmetransport.NewGRPCClient(conn, logger, tracer),
		Integration: intregrationtransport.NewGRPCClient(conn, logger),
		Workflow:    workflowtransport.NewGRPCClient(conn, logger, tracer, authSigner),
	}, nil
}
