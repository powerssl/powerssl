package client // import "powerssl.io/pkg/controller/client"

import (
	"github.com/go-kit/kit/log"
	stdopentracing "github.com/opentracing/opentracing-go"

	acmetransport "powerssl.io/internal/app/controller/acme/transport"
	intregrationtransport "powerssl.io/internal/app/controller/integration/transport"
	workflowtransport "powerssl.io/internal/app/controller/workflow/transport"
	"powerssl.io/internal/pkg/auth"
	"powerssl.io/internal/pkg/util"
	"powerssl.io/pkg/controller/acme"
	"powerssl.io/pkg/controller/integration"
	"powerssl.io/pkg/controller/workflow"
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
