package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/tracing/opentracing"
	stdopentracing "github.com/opentracing/opentracing-go"

	"powerssl.dev/powerssl/internal/pkg/middleware"
	"powerssl.dev/powerssl/pkg/controller/api"
	"powerssl.dev/powerssl/pkg/controller/workflow"
)

type Endpoints struct {
	CreateEndpoint endpoint.Endpoint
}

func NewEndpoints(svc workflow.Service, logger log.Logger, tracer stdopentracing.Tracer, duration metrics.Histogram, auth endpoint.Middleware) Endpoints {
	var createEndpoint endpoint.Endpoint
	{
		createEndpoint = makeCreateEndpoint(svc)
		createEndpoint = auth(createEndpoint)
		createEndpoint = opentracing.TraceServer(tracer, "Create")(createEndpoint)
		createEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "Create"))(createEndpoint)
		createEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "Create"))(createEndpoint)
	}

	return Endpoints{
		CreateEndpoint: createEndpoint,
	}
}

func (e Endpoints) Create(ctx context.Context, workflow *api.Workflow) (*api.Workflow, error) {
	resp, err := e.CreateEndpoint(ctx, CreateRequest{
		Workflow: workflow,
	})
	if err != nil {
		return nil, err
	}
	response := resp.(CreateResponse)
	return response.Workflow, nil
}

type CreateRequest struct {
	Workflow *api.Workflow
}

type CreateResponse struct {
	Workflow *api.Workflow
}

func makeCreateEndpoint(s workflow.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		workflow, err := s.Create(ctx, req.Workflow)
		if err != nil {
			return nil, err
		}
		return CreateResponse{
			Workflow: workflow,
		}, nil
	}
}
