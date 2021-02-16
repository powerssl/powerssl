package certificate

import (
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/tracing/opentracing"
	stdopentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"powerssl.dev/backend/middleware"
	"powerssl.dev/sdk/apiserver/certificate"

	"powerssl.dev/apiserver/internal/certificate/service"
	"powerssl.dev/apiserver/internal/repository"
	apiv1 "powerssl.dev/sdk/apiserver/api/v1"
	"powerssl.dev/sdk/apiserver/certificate/endpoint"
	"powerssl.dev/sdk/apiserver/certificate/transport"
)

type Certificate struct {
	endpoints endpoint.Endpoints
	logger    log.Logger
	tracer    stdopentracing.Tracer
}

func New(repositories *repository.Repositories, logger log.Logger, tracer stdopentracing.Tracer, duration metrics.Histogram, auth kitendpoint.Middleware) *Certificate {
	svc := service.New(repositories, logger)
	endpoints := makeEndpoints(svc, logger, tracer, duration, auth)

	return &Certificate{
		endpoints: endpoints,
		logger:    logger,
		tracer:    tracer,
	}
}

func (certificate *Certificate) RegisterGRPCServer(baseServer *grpc.Server) {
	grpcServer := transport.NewGRPCServer(certificate.endpoints, certificate.logger, certificate.tracer)
	apiv1.RegisterCertificateServiceServer(baseServer, grpcServer)
}

func (*Certificate) ServiceName() string {
	return "powerssl.apiserver.v1.CertificateService"
}

func makeEndpoints(svc certificate.Service, logger log.Logger, tracer stdopentracing.Tracer, duration metrics.Histogram, auth kitendpoint.Middleware) endpoint.Endpoints {
	var createEndpoint kitendpoint.Endpoint
	{
		createEndpoint = endpoint.MakeCreateEndpoint(svc)
		createEndpoint = auth(createEndpoint)
		createEndpoint = opentracing.TraceServer(tracer, "Create")(createEndpoint)
		createEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "Create"))(createEndpoint)
		createEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "Create"))(createEndpoint)
	}

	var deleteEndpoint kitendpoint.Endpoint
	{
		deleteEndpoint = endpoint.MakeDeleteEndpoint(svc)
		deleteEndpoint = auth(deleteEndpoint)
		deleteEndpoint = opentracing.TraceServer(tracer, "Delete")(deleteEndpoint)
		deleteEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "Delete"))(deleteEndpoint)
		deleteEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "Delete"))(deleteEndpoint)
	}

	var getEndpoint kitendpoint.Endpoint
	{
		getEndpoint = endpoint.MakeGetEndpoint(svc)
		getEndpoint = auth(getEndpoint)
		getEndpoint = opentracing.TraceServer(tracer, "Get")(getEndpoint)
		getEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "Get"))(getEndpoint)
		getEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "Get"))(getEndpoint)
	}

	var listEndpoint kitendpoint.Endpoint
	{
		listEndpoint = endpoint.MakeListEndpoint(svc)
		listEndpoint = auth(listEndpoint)
		listEndpoint = opentracing.TraceServer(tracer, "List")(listEndpoint)
		listEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "List"))(listEndpoint)
		listEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "List"))(listEndpoint)
	}

	var updateEndpoint kitendpoint.Endpoint
	{
		updateEndpoint = endpoint.MakeUpdateEndpoint(svc)
		updateEndpoint = auth(updateEndpoint)
		updateEndpoint = opentracing.TraceServer(tracer, "Update")(updateEndpoint)
		updateEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "Update"))(updateEndpoint)
		updateEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "Update"))(updateEndpoint)
	}

	return endpoint.Endpoints{
		CreateEndpoint: createEndpoint,
		DeleteEndpoint: deleteEndpoint,
		GetEndpoint:    getEndpoint,
		ListEndpoint:   listEndpoint,
		UpdateEndpoint: updateEndpoint,
	}
}
