package endpoint // import "powerssl.io/pkg/apiserver/acmeaccount/endpoint"

import (
	"context"

	"github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/tracing/opentracing"
	stdopentracing "github.com/opentracing/opentracing-go"

	"powerssl.io/pkg/apiserver/acmeaccount/meta"
	"powerssl.io/pkg/apiserver/api"
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

func NewEndpoints(svc meta.Service, logger log.Logger, tracer stdopentracing.Tracer, duration metrics.Histogram) Endpoints {
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

func (e Endpoints) Create(ctx context.Context, parent string, acmeAccount *api.ACMEAccount) (*api.ACMEAccount, error) {
	resp, err := e.CreateEndpoint(ctx, CreateRequest{
		Parent:      parent,
		ACMEAccount: acmeAccount,
	})
	if err != nil {
		return nil, err
	}
	response := resp.(CreateResponse)
	return response.ACMEAccount, nil
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

func (e Endpoints) Get(ctx context.Context, name string) (*api.ACMEAccount, error) {
	resp, err := e.GetEndpoint(ctx, GetRequest{
		Name: name,
	})
	if err != nil {
		return nil, err
	}
	response := resp.(GetResponse)
	return response.ACMEAccount, nil
}

func (e Endpoints) List(ctx context.Context, parent string, pageSize int, pageToken string) ([]*api.ACMEAccount, string, error) {
	resp, err := e.ListEndpoint(ctx, ListRequest{
		Parent:    parent,
		PageSize:  pageSize,
		PageToken: pageToken,
	})
	if err != nil {
		return nil, "", err
	}
	response := resp.(ListResponse)
	return response.ACMEAccounts, response.NextPageToken, nil
}

func (e Endpoints) Update(ctx context.Context, name string, acmeAccount *api.ACMEAccount) (*api.ACMEAccount, error) {
	resp, err := e.UpdateEndpoint(ctx, UpdateRequest{
		Name:        name,
		ACMEAccount: acmeAccount,
	})
	if err != nil {
		return nil, err
	}
	response := resp.(UpdateResponse)
	return response.ACMEAccount, nil
}

type CreateRequest struct {
	Parent      string
	ACMEAccount *api.ACMEAccount
}

type CreateResponse struct {
	ACMEAccount *api.ACMEAccount
}

func makeCreateEndpoint(s meta.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		acmeAccount, err := s.Create(ctx, req.Parent, req.ACMEAccount)
		if err != nil {
			return nil, err
		}
		return CreateResponse{
			ACMEAccount: acmeAccount,
		}, nil
	}
}

type DeleteRequest struct {
	Name string
}

type DeleteResponse struct{}

func makeDeleteEndpoint(s meta.Service) endpoint.Endpoint {
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
	ACMEAccount *api.ACMEAccount
}

func makeGetEndpoint(s meta.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRequest)
		acmeAccount, err := s.Get(ctx, req.Name)
		if err != nil {
			return nil, err
		}
		return GetResponse{
			ACMEAccount: acmeAccount,
		}, nil
	}
}

type ListRequest struct {
	Parent    string
	PageSize  int
	PageToken string
}

type ListResponse struct {
	ACMEAccounts  []*api.ACMEAccount
	NextPageToken string
}

func makeListEndpoint(s meta.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ListRequest)
		acmeAccounts, nextPageToken, err := s.List(ctx, req.Parent, req.PageSize, req.PageToken)
		if err != nil {
			return nil, err
		}
		return ListResponse{
			ACMEAccounts:  acmeAccounts,
			NextPageToken: nextPageToken,
		}, nil
	}
}

type UpdateRequest struct {
	Name        string
	ACMEAccount *api.ACMEAccount
}

type UpdateResponse struct {
	ACMEAccount *api.ACMEAccount
}

func makeUpdateEndpoint(s meta.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateRequest)
		acmeAccount, err := s.Update(ctx, req.Name, req.ACMEAccount)
		if err != nil {
			return nil, err
		}
		return UpdateResponse{
			ACMEAccount: acmeAccount,
		}, nil
	}
}
