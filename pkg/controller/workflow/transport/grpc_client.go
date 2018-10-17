package transport // import "powerssl.io/pkg/controller/workflow/transport"

import (
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"

	apiv1 "powerssl.io/pkg/controller/api/v1"
	"powerssl.io/pkg/controller/workflow/endpoint"
	service "powerssl.io/pkg/controller/workflow/service"
)

const serviceName = "powerssl.controller.v1.WorkflowService"

func NewGRPCClient(conn *grpc.ClientConn, logger log.Logger) service.Service {
	options := []grpctransport.ClientOption{}

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
	}

	return endpoint.Endpoints{
		CreateEndpoint: createEndpoint,
	}
}
