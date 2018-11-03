package endpoint // import "powerssl.io/pkg/controller/workflow/endpoint"

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"

	"github.com/go-kit/kit/tracing/opentracing"
	stdopentracing "github.com/opentracing/opentracing-go"
	"powerssl.io/pkg/controller/api"
	service "powerssl.io/pkg/controller/workflow/service"
	"powerssl.io/pkg/util/middleware"
)

type Endpoints struct {
	CreateEndpoint endpoint.Endpoint
}

func NewEndpoints(svc service.Service, logger log.Logger, tracer stdopentracing.Tracer, duration metrics.Histogram) Endpoints {
	var createEndpoint endpoint.Endpoint
	{
		createEndpoint = makeCreateEndpoint(svc)
		createEndpoint = opentracing.TraceServer(tracer, "Create")(createEndpoint)
		createEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "Create"))(createEndpoint)
		createEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "Create"))(createEndpoint)
	}

	return Endpoints{
		CreateEndpoint: createEndpoint,
	}
}

func (e Endpoints) Create(ctx context.Context, kind string) (*api.Workflow, error) {
	resp, err := e.CreateEndpoint(ctx, CreateRequest{
		Kind: kind,
	})
	if err != nil {
		return nil, err
	}
	response := resp.(CreateResponse)
	return response.Workflow, nil
}

type CreateRequest struct {
	Kind string
}

type CreateResponse struct {
	Workflow *api.Workflow
}

func makeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		workflow, err := s.Create(ctx, req.Kind)
		if err != nil {
			return nil, err
		}
		return CreateResponse{
			Workflow: workflow,
		}, nil
	}
}
