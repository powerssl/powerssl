package workflow // import "powerssl.io/pkg/controller/workflow"

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"google.golang.org/grpc"

	apiv1 "powerssl.io/pkg/controller/api/v1"
	"powerssl.io/pkg/controller/workflow/endpoint"
	workflowengine "powerssl.io/pkg/controller/workflow/engine"
	service "powerssl.io/pkg/controller/workflow/service"
	"powerssl.io/pkg/controller/workflow/transport"
	resource "powerssl.io/pkg/resource"
)

type Workflow struct {
	endpoints endpoint.Endpoints
	logger    log.Logger
}

func New(logger log.Logger, duration metrics.Histogram, workflowengine *workflowengine.Engine) resource.Resource {
	svc := service.New(logger, workflowengine)
	endpoints := endpoint.NewEndpoints(svc, logger, duration)

	return &Workflow{
		endpoints: endpoints,
		logger:    logger,
	}
}

func (controller *Workflow) RegisterGRPCServer(baseServer *grpc.Server) {
	grpcServer := transport.NewGRPCServer(controller.endpoints, controller.logger)
	apiv1.RegisterWorkflowServiceServer(baseServer, grpcServer)
}
