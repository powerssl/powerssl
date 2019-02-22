package client // import "powerssl.io/pkg/controller/client"

import (
	"github.com/go-kit/kit/log"
	stdopentracing "github.com/opentracing/opentracing-go"

	acmemeta "powerssl.io/internal/app/controller/acme/meta"
	acmetransport "powerssl.io/internal/app/controller/acme/transport"
	integrationmeta "powerssl.io/internal/app/controller/integration/meta"
	intregrationtransport "powerssl.io/internal/app/controller/integration/transport"
	workflowmeta "powerssl.io/internal/app/controller/workflow/meta"
	workflowtransport "powerssl.io/internal/app/controller/workflow/transport"
	"powerssl.io/internal/pkg/util"
	"powerssl.io/internal/pkg/util/auth"
)

type GRPCClient struct {
	ACME        acmemeta.Service
	Integration integrationmeta.Service
	Workflow    workflowmeta.Service
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
