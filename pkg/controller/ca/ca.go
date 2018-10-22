package ca // import "powerssl.io/pkg/controller/ca"

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"google.golang.org/grpc"

	apiv1 "powerssl.io/pkg/controller/api/v1"
	"powerssl.io/pkg/controller/ca/endpoint"
	service "powerssl.io/pkg/controller/ca/service"
	"powerssl.io/pkg/controller/ca/transport"
	workflowengine "powerssl.io/pkg/controller/workflow/engine"
	resource "powerssl.io/pkg/resource"
)

type CA struct {
	endpoints endpoint.Endpoints
	logger    log.Logger
}

func New(logger log.Logger, duration metrics.Histogram, workflowengine *workflowengine.Engine) resource.Resource {
	svc := service.New(logger, workflowengine)
	endpoints := endpoint.NewEndpoints(svc, logger, duration)

	return &CA{
		endpoints: endpoints,
		logger:    logger,
	}
}

func (controller *CA) RegisterGRPCServer(baseServer *grpc.Server) {
	grpcServer := transport.NewGRPCServer(controller.endpoints, controller.logger)
	apiv1.RegisterCAServiceServer(baseServer, grpcServer)
}
