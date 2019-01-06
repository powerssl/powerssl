package workflow // import "powerssl.io/pkg/controller/workflow"

import (
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"google.golang.org/grpc"

	stdopentracing "github.com/opentracing/opentracing-go"
	apiserverclient "powerssl.io/pkg/apiserver/client"
	apiv1 "powerssl.io/pkg/controller/api/v1"
	"powerssl.io/pkg/controller/workflow/endpoint"
	service "powerssl.io/pkg/controller/workflow/service"
	"powerssl.io/pkg/controller/workflow/transport"
)

type Workflow struct {
	endpoints endpoint.Endpoints
	logger    log.Logger
	tracer    stdopentracing.Tracer
}

func New(logger log.Logger, tracer stdopentracing.Tracer, duration metrics.Histogram, client *apiserverclient.GRPCClient, auth kitendpoint.Middleware) *Workflow {
	svc := service.New(logger, client)
	endpoints := endpoint.NewEndpoints(svc, logger, tracer, duration, auth)

	return &Workflow{
		endpoints: endpoints,
		logger:    logger,
		tracer:    tracer,
	}
}

func (w *Workflow) RegisterGRPCServer(baseServer *grpc.Server) {
	grpcServer := transport.NewGRPCServer(w.endpoints, w.logger, w.tracer)
	apiv1.RegisterWorkflowServiceServer(baseServer, grpcServer)
}

func (*Workflow) ServiceName() string {
	return "powerssl.controller.v1.WorkflowService"
}
