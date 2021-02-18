package acmeaccount

import (
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/tracing/opentracing"
	stdopentracing "github.com/opentracing/opentracing-go"
	temporalclient "go.temporal.io/sdk/client"
	"google.golang.org/grpc"
	"powerssl.dev/backend/middleware"
	"powerssl.dev/sdk/apiserver/acmeaccount"

	"powerssl.dev/apiserver/internal/acmeaccount/service"
	"powerssl.dev/apiserver/internal/repository"
	backendtransport "powerssl.dev/backend/transport"
	"powerssl.dev/sdk/apiserver/acmeaccount/endpoint"
	"powerssl.dev/sdk/apiserver/acmeaccount/transport"
	apiv1 "powerssl.dev/sdk/apiserver/api/v1"
)

func NewService(repositories *repository.Repositories, logger log.Logger, tracer stdopentracing.Tracer, duration metrics.Histogram, temporalClient temporalclient.Client, auth kitendpoint.Middleware) backendtransport.Service {
	svc := service.New(repositories, logger, temporalClient)
	endpoints := makeEndpoints(svc, logger, tracer, duration, auth)

	return backendtransport.Service{
		ServiceName: "powerssl.apiserver.v1.ACMEAccountService",
		RegisterGRPCServer: func(baseServer *grpc.Server) {
			grpcServer := transport.NewGRPCServer(endpoints, logger, tracer)
			apiv1.RegisterACMEAccountServiceServer(baseServer, grpcServer)
		},
	}
}

func makeEndpoints(svc acmeaccount.Service, logger log.Logger, tracer stdopentracing.Tracer, duration metrics.Histogram, auth kitendpoint.Middleware) endpoint.Endpoints {
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
