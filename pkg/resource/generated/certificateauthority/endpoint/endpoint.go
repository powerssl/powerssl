// Code generated by protoc-gen-gotemplate. DO NOT EDIT.

package endpoint // import "powerssl.io/pkg/resource/generated/certificateauthority/endpoint"

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"

	"powerssl.io/pkg/api"
	resource "powerssl.io/pkg/resource"
	service "powerssl.io/pkg/resource/certificateauthority"
)

type Endpoints struct {
	CreateEndpoint endpoint.Endpoint
	DeleteEndpoint endpoint.Endpoint
	GetEndpoint    endpoint.Endpoint
	ListEndpoint   endpoint.Endpoint
	UpdateEndpoint endpoint.Endpoint
}

func NewEndpoints(svc service.Service, logger log.Logger, duration metrics.Histogram) Endpoints {
	var createEndpoint endpoint.Endpoint
	{
		createEndpoint = makeCreateEndpoint(svc)
		createEndpoint = resource.LoggingMiddleware(log.With(logger, "method", "Create"))(createEndpoint)
		createEndpoint = resource.InstrumentingMiddleware(duration.With("method", "Create"))(createEndpoint)
	}

	var deleteEndpoint endpoint.Endpoint
	{
		deleteEndpoint = makeDeleteEndpoint(svc)
		deleteEndpoint = resource.LoggingMiddleware(log.With(logger, "method", "Delete"))(deleteEndpoint)
		deleteEndpoint = resource.InstrumentingMiddleware(duration.With("method", "Delete"))(deleteEndpoint)
	}

	var getEndpoint endpoint.Endpoint
	{
		getEndpoint = makeGetEndpoint(svc)
		getEndpoint = resource.LoggingMiddleware(log.With(logger, "method", "Get"))(getEndpoint)
		getEndpoint = resource.InstrumentingMiddleware(duration.With("method", "Get"))(getEndpoint)
	}

	var listEndpoint endpoint.Endpoint
	{
		listEndpoint = makeListEndpoint(svc)
		listEndpoint = resource.LoggingMiddleware(log.With(logger, "method", "List"))(listEndpoint)
		listEndpoint = resource.InstrumentingMiddleware(duration.With("method", "List"))(listEndpoint)
	}

	var updateEndpoint endpoint.Endpoint
	{
		updateEndpoint = makeUpdateEndpoint(svc)
		updateEndpoint = resource.LoggingMiddleware(log.With(logger, "method", "Update"))(updateEndpoint)
		updateEndpoint = resource.InstrumentingMiddleware(duration.With("method", "Update"))(updateEndpoint)
	}

	return Endpoints{
		CreateEndpoint: createEndpoint,
		DeleteEndpoint: deleteEndpoint,
		GetEndpoint:    getEndpoint,
		ListEndpoint:   listEndpoint,
		UpdateEndpoint: updateEndpoint,
	}
}

func (e Endpoints) Create(ctx context.Context, certificateAuthority *api.CertificateAuthority) (*api.CertificateAuthority, error) {
	resp, err := e.CreateEndpoint(ctx, CreateRequest{
		CertificateAuthority: certificateAuthority,
	})
	if err != nil {
		return nil, err
	}
	response := resp.(CreateResponse)
	return response.CertificateAuthority, nil
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

func (e Endpoints) Get(ctx context.Context, name string) (*api.CertificateAuthority, error) {
	resp, err := e.GetEndpoint(ctx, GetRequest{
		Name: name,
	})
	if err != nil {
		return nil, err
	}
	response := resp.(GetResponse)
	return response.CertificateAuthority, nil
}

func (e Endpoints) List(ctx context.Context) ([]*api.CertificateAuthority, error) {
	resp, err := e.ListEndpoint(ctx, ListRequest{})
	if err != nil {
		return nil, err
	}
	response := resp.(ListResponse)
	return response.CertificateAuthorities, nil
}

func (e Endpoints) Update(ctx context.Context, certificateAuthority *api.CertificateAuthority) (*api.CertificateAuthority, error) {
	resp, err := e.UpdateEndpoint(ctx, UpdateRequest{
		CertificateAuthority: certificateAuthority,
	})
	if err != nil {
		return nil, err
	}
	response := resp.(UpdateResponse)
	return response.CertificateAuthority, nil
}

type CreateRequest struct {
	CertificateAuthority *api.CertificateAuthority
}

type CreateResponse struct {
	CertificateAuthority *api.CertificateAuthority
}

func makeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		certificateAuthority, err := s.Create(ctx, req.CertificateAuthority)
		if err != nil {
			return nil, err
		}
		return CreateResponse{
			CertificateAuthority: certificateAuthority,
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
	CertificateAuthority *api.CertificateAuthority
}

func makeGetEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRequest)
		certificateAuthority, err := s.Get(ctx, req.Name)
		if err != nil {
			return nil, err
		}
		return GetResponse{
			CertificateAuthority: certificateAuthority,
		}, nil
	}
}

type ListRequest struct{}

type ListResponse struct {
	CertificateAuthorities []*api.CertificateAuthority
}

func makeListEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		certificateAuthorities, err := s.List(ctx)
		if err != nil {
			return nil, err
		}
		return ListResponse{
			CertificateAuthorities: certificateAuthorities,
		}, nil
	}
}

type UpdateRequest struct {
	CertificateAuthority *api.CertificateAuthority
}

type UpdateResponse struct {
	CertificateAuthority *api.CertificateAuthority
}

func makeUpdateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateRequest)
		certificateAuthority, err := s.Update(ctx, req.CertificateAuthority)
		if err != nil {
			return nil, err
		}
		return UpdateResponse{
			CertificateAuthority: certificateAuthority,
		}, nil
	}
}
