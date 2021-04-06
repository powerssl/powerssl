package user

import (
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/tracing/opentracing"
	stdopentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"

	apiv1 "powerssl.dev/api/apiserver/v1"
	"powerssl.dev/backend/middleware"
	backendtransport "powerssl.dev/backend/transport"
	"powerssl.dev/common/log"
	"powerssl.dev/sdk/apiserver/user"
	"powerssl.dev/sdk/apiserver/user/endpoint"
	"powerssl.dev/sdk/apiserver/user/transport"

	"powerssl.dev/apiserver/internal/repository"
	"powerssl.dev/apiserver/internal/user/service"
)

func NewService(repositories *repository.Repositories, logger log.Logger, tracer stdopentracing.Tracer, duration metrics.Histogram, auth kitendpoint.Middleware) backendtransport.Service {
	svc := service.New(repositories, logger)
	endpoints := makeEndpoints(svc, logger, tracer, duration, auth)

	return backendtransport.Service{
		ServiceName: "powerssl.apiserver.v1.UserService",
		RegisterGRPCServer: func(baseServer *grpc.Server) {
			grpcServer := transport.NewGRPCServer(endpoints, logger, tracer)
			apiv1.RegisterUserServiceServer(baseServer, grpcServer)
		},
	}
}

func makeEndpoints(svc user.Service, logger log.Logger, tracer stdopentracing.Tracer, duration metrics.Histogram, auth kitendpoint.Middleware) endpoint.Endpoints {
	var createEndpoint kitendpoint.Endpoint
	{
		createEndpoint = endpoint.MakeCreateEndpoint(svc)
		createEndpoint = auth(createEndpoint)
		createEndpoint = opentracing.TraceServer(tracer, "Create")(createEndpoint)
		createEndpoint = middleware.LoggingMiddleware(logger.With("method", "Create"))(createEndpoint)
		createEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "Create"))(createEndpoint)
	}

	var deleteEndpoint kitendpoint.Endpoint
	{
		deleteEndpoint = endpoint.MakeDeleteEndpoint(svc)
		deleteEndpoint = auth(deleteEndpoint)
		deleteEndpoint = opentracing.TraceServer(tracer, "Delete")(deleteEndpoint)
		deleteEndpoint = middleware.LoggingMiddleware(logger.With("method", "Delete"))(deleteEndpoint)
		deleteEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "Delete"))(deleteEndpoint)
	}

	var getEndpoint kitendpoint.Endpoint
	{
		getEndpoint = endpoint.MakeGetEndpoint(svc)
		getEndpoint = auth(getEndpoint)
		getEndpoint = opentracing.TraceServer(tracer, "Get")(getEndpoint)
		getEndpoint = middleware.LoggingMiddleware(logger.With("method", "Get"))(getEndpoint)
		getEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "Get"))(getEndpoint)
	}

	var listEndpoint kitendpoint.Endpoint
	{
		listEndpoint = endpoint.MakeListEndpoint(svc)
		listEndpoint = auth(listEndpoint)
		listEndpoint = opentracing.TraceServer(tracer, "List")(listEndpoint)
		listEndpoint = middleware.LoggingMiddleware(logger.With("method", "List"))(listEndpoint)
		listEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "List"))(listEndpoint)
	}

	var updateEndpoint kitendpoint.Endpoint
	{
		updateEndpoint = endpoint.MakeUpdateEndpoint(svc)
		updateEndpoint = auth(updateEndpoint)
		updateEndpoint = opentracing.TraceServer(tracer, "Update")(updateEndpoint)
		updateEndpoint = middleware.LoggingMiddleware(logger.With("method", "Update"))(updateEndpoint)
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
