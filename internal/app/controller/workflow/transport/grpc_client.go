package transport

import (
	"fmt"

	"github.com/go-kit/kit/auth/jwt"
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	stdopentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"

	"powerssl.io/powerssl/internal/app/controller/workflow/endpoint"
	apiv1 "powerssl.io/powerssl/internal/pkg/controller/api/v1"
	"powerssl.io/powerssl/pkg/controller/workflow"
)

const serviceName = "powerssl.controller.v1.WorkflowService"

func NewGRPCClient(conn *grpc.ClientConn, logger log.Logger, tracer stdopentracing.Tracer, authSigner kitendpoint.Middleware) workflow.Service {
	options := []grpctransport.ClientOption{
		grpctransport.ClientBefore(jwt.ContextToGRPC()),
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	}

	var createEndpoint kitendpoint.Endpoint
	{
		createEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			"Create",
			encodeGRPCCreateRequest,
			decodeGRPCCreateResponse,
			apiv1.Workflow{},
			options...,
		).Endpoint()
		createEndpoint = authSigner(createEndpoint)
		createEndpoint = opentracing.TraceClient(tracer, fmt.Sprintf("/%s/Create", serviceName))(createEndpoint)
	}

	return endpoint.Endpoints{
		CreateEndpoint: createEndpoint,
	}
}
