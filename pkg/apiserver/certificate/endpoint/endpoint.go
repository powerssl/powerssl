package endpoint // import "powerssl.io/pkg/apiserver/certificate/endpoint"

import (
	"context"

	"github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/tracing/opentracing"
	stdopentracing "github.com/opentracing/opentracing-go"

	"powerssl.io/pkg/apiserver/api"
	"powerssl.io/pkg/apiserver/certificate/service"
	"powerssl.io/pkg/util/auth"
	"powerssl.io/pkg/util/middleware"
)

type Endpoints struct {
	CreateEndpoint endpoint.Endpoint
	DeleteEndpoint endpoint.Endpoint
	GetEndpoint    endpoint.Endpoint
	ListEndpoint   endpoint.Endpoint
	UpdateEndpoint endpoint.Endpoint
}

func NewEndpoints(svc service.Service, logger log.Logger, tracer stdopentracing.Tracer, duration metrics.Histogram) Endpoints {
	jwtParser := jwt.NewParser(auth.KeyFunc, auth.Method, jwt.StandardClaimsFactory)

	var createEndpoint endpoint.Endpoint
	{
		createEndpoint = makeCreateEndpoint(svc)
		createEndpoint = jwtParser(createEndpoint)
		createEndpoint = opentracing.TraceServer(tracer, "Create")(createEndpoint)
		createEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "Create"))(createEndpoint)
		createEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "Create"))(createEndpoint)
	}

	var deleteEndpoint endpoint.Endpoint
	{
		deleteEndpoint = makeDeleteEndpoint(svc)
		deleteEndpoint = jwtParser(deleteEndpoint)
		deleteEndpoint = opentracing.TraceServer(tracer, "Delete")(deleteEndpoint)
		deleteEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "Delete"))(deleteEndpoint)
		deleteEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "Delete"))(deleteEndpoint)
	}

	var getEndpoint endpoint.Endpoint
	{
		getEndpoint = makeGetEndpoint(svc)
		getEndpoint = jwtParser(getEndpoint)
		getEndpoint = opentracing.TraceServer(tracer, "Get")(getEndpoint)
		getEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "Get"))(getEndpoint)
		getEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "Get"))(getEndpoint)
	}

	var listEndpoint endpoint.Endpoint
	{
		listEndpoint = makeListEndpoint(svc)
		listEndpoint = jwtParser(listEndpoint)
		listEndpoint = opentracing.TraceServer(tracer, "List")(listEndpoint)
		listEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "List"))(listEndpoint)
		listEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "List"))(listEndpoint)
	}

	var updateEndpoint endpoint.Endpoint
	{
		updateEndpoint = makeUpdateEndpoint(svc)
		updateEndpoint = jwtParser(updateEndpoint)
		updateEndpoint = opentracing.TraceServer(tracer, "Update")(updateEndpoint)
		updateEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "Update"))(updateEndpoint)
		updateEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "Update"))(updateEndpoint)
	}

	return Endpoints{
		CreateEndpoint: createEndpoint,
		DeleteEndpoint: deleteEndpoint,
		GetEndpoint:    getEndpoint,
		ListEndpoint:   listEndpoint,
		UpdateEndpoint: updateEndpoint,
	}
}

func (e Endpoints) Create(ctx context.Context, certificate *api.Certificate) (*api.Certificate, error) {
	resp, err := e.CreateEndpoint(ctx, CreateRequest{
		Certificate: certificate,
	})
	if err != nil {
		return nil, err
	}
	response := resp.(CreateResponse)
	return response.Certificate, nil
}

func (e Endpoints) Delete(ctx context.Context, name string) error {
	_, err := e.DeleteEndpoint(ctx, DeleteRequest{
		Name: name,
	})
	if err != nil {
		return err
	}
	return nil
}

func (e Endpoints) Get(ctx context.Context, name string) (*api.Certificate, error) {
	resp, err := e.GetEndpoint(ctx, GetRequest{
		Name: name,
	})
	if err != nil {
		return nil, err
	}
	response := resp.(GetResponse)
	return response.Certificate, nil
}

func (e Endpoints) List(ctx context.Context, pageSize int, pageToken string) ([]*api.Certificate, string, error) {
	resp, err := e.ListEndpoint(ctx, ListRequest{
		PageSize:  pageSize,
		PageToken: pageToken,
	})
	if err != nil {
		return nil, "", err
	}
	response := resp.(ListResponse)
	return response.Certificates, response.NextPageToken, nil
}

func (e Endpoints) Update(ctx context.Context, name string, certificate *api.Certificate) (*api.Certificate, error) {
	resp, err := e.UpdateEndpoint(ctx, UpdateRequest{
		Name:        name,
		Certificate: certificate,
	})
	if err != nil {
		return nil, err
	}
	response := resp.(UpdateResponse)
	return response.Certificate, nil
}

type CreateRequest struct {
	Certificate *api.Certificate
}

type CreateResponse struct {
	Certificate *api.Certificate
}

func makeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		certificate, err := s.Create(ctx, req.Certificate)
		if err != nil {
			return nil, err
		}
		return CreateResponse{
			Certificate: certificate,
		}, nil
	}
}

type DeleteRequest struct {
	Name string
}

type DeleteResponse struct{}

func makeDeleteEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRequest)
		if err := s.Delete(ctx, req.Name); err != nil {
			return nil, err
		}
		return DeleteResponse{}, nil
	}
}

type GetRequest struct {
	Name string
}

type GetResponse struct {
	Certificate *api.Certificate
}

func makeGetEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRequest)
		certificate, err := s.Get(ctx, req.Name)
		if err != nil {
			return nil, err
		}
		return GetResponse{
			Certificate: certificate,
		}, nil
	}
}

type ListRequest struct {
	PageSize  int
	PageToken string
}

type ListResponse struct {
	Certificates  []*api.Certificate
	NextPageToken string
}

func makeListEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ListRequest)
		certificates, nextPageToken, err := s.List(ctx, req.PageSize, req.PageToken)
		if err != nil {
			return nil, err
		}
		return ListResponse{
			Certificates:  certificates,
			NextPageToken: nextPageToken,
		}, nil
	}
}

type UpdateRequest struct {
	Name        string
	Certificate *api.Certificate
}

type UpdateResponse struct {
	Certificate *api.Certificate
}

func makeUpdateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateRequest)
		certificate, err := s.Update(ctx, req.Name, req.Certificate)
		if err != nil {
			return nil, err
		}
		return UpdateResponse{
			Certificate: certificate,
		}, nil
	}
}
