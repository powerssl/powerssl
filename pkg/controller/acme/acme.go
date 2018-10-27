package acme // import "powerssl.io/pkg/controller/acme"

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"google.golang.org/grpc"

	"powerssl.io/pkg/controller/acme/endpoint"
	service "powerssl.io/pkg/controller/acme/service"
	"powerssl.io/pkg/controller/acme/transport"
	apiv1 "powerssl.io/pkg/controller/api/v1"
	workflowengine "powerssl.io/pkg/controller/workflow/engine"
	resource "powerssl.io/pkg/resource"
)

type ACME struct {
	endpoints endpoint.Endpoints
	logger    log.Logger
}

func New(logger log.Logger, duration metrics.Histogram, workflowengine *workflowengine.Engine) resource.Resource {
	svc := service.New(logger, workflowengine)
	endpoints := endpoint.NewEndpoints(svc, logger, duration)

	return &ACME{
		endpoints: endpoints,
		logger:    logger,
	}
}

func (controller *ACME) RegisterGRPCServer(baseServer *grpc.Server) {
	grpcServer := transport.NewGRPCServer(controller.endpoints, controller.logger)
	apiv1.RegisterACMEServiceServer(baseServer, grpcServer)
}
