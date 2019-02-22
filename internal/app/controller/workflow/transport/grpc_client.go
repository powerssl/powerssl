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

	"powerssl.io/internal/app/controller/workflow/endpoint"
	"powerssl.io/internal/app/controller/workflow/meta"
	apiv1 "powerssl.io/pkg/controller/api/v1"
)

const serviceName = "powerssl.controller.v1.WorkflowService"

func NewGRPCClient(conn *grpc.ClientConn, logger log.Logger, tracer stdopentracing.Tracer, authSigner kitendpoint.Middleware) meta.Service {
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
