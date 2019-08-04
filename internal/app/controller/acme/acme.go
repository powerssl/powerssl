package acme

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"google.golang.org/grpc"

	stdopentracing "github.com/opentracing/opentracing-go"
	"powerssl.dev/powerssl/internal/app/controller/acme/endpoint"
	service "powerssl.dev/powerssl/internal/app/controller/acme/service"
	"powerssl.dev/powerssl/internal/app/controller/acme/transport"
	apiv1 "powerssl.dev/powerssl/internal/pkg/controller/api/v1"
)

type ACME struct {
	endpoints endpoint.Endpoints
	logger    log.Logger
	tracer    stdopentracing.Tracer
}

func New(logger log.Logger, tracer stdopentracing.Tracer, duration metrics.Histogram) *ACME {
	svc := service.New(logger)
	endpoints := endpoint.NewEndpoints(svc, logger, tracer, duration)

	return &ACME{
		endpoints: endpoints,
		logger:    logger,
		tracer:    tracer,
	}
}

func (a *ACME) RegisterGRPCServer(baseServer *grpc.Server) {
	grpcServer := transport.NewGRPCServer(a.endpoints, a.logger, a.tracer)
	apiv1.RegisterACMEServiceServer(baseServer, grpcServer)
}

func (*ACME) ServiceName() string {
	return "powerssl.controller.v1.ACMEService"
}
