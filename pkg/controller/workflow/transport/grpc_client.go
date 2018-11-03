package transport // import "powerssl.io/pkg/controller/workflow/transport"

import (
	"fmt"

	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	stdopentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"

	apiv1 "powerssl.io/pkg/controller/api/v1"
	"powerssl.io/pkg/controller/workflow/endpoint"
	service "powerssl.io/pkg/controller/workflow/service"
)

const serviceName = "powerssl.controller.v1.WorkflowService"

func NewGRPCClient(conn *grpc.ClientConn, logger log.Logger, tracer stdopentracing.Tracer) service.Service {
	options := []grpctransport.ClientOption{
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
		createEndpoint = opentracing.TraceClient(tracer, fmt.Sprintf("/%s/Create", serviceName))(createEndpoint)
	}

	return endpoint.Endpoints{
		CreateEndpoint: createEndpoint,
	}
}
