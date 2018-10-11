// Code generated by protoc-gen-gotemplate. DO NOT EDIT.

package endpoint // import "powerssl.io/pkg/resource/generated/certificateissue/endpoint"

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"

	"powerssl.io/pkg/api"
	resource "powerssl.io/pkg/resource"
	service "powerssl.io/pkg/resource/certificateissue"
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

func (e Endpoints) Create(ctx context.Context, certificateIssue *api.CertificateIssue) (*api.CertificateIssue, error) {
	resp, err := e.CreateEndpoint(ctx, CreateRequest{
		CertificateIssue: certificateIssue,
	})
	if err != nil {
		return nil, err
	}
	response := resp.(CreateResponse)
	return response.CertificateIssue, nil
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

func (e Endpoints) Get(ctx context.Context, name string) (*api.CertificateIssue, error) {
	resp, err := e.GetEndpoint(ctx, GetRequest{
		Name: name,
	})
	if err != nil {
		return nil, err
	}
	response := resp.(GetResponse)
	return response.CertificateIssue, nil
}

func (e Endpoints) List(ctx context.Context, pageSize int, pageToken string) ([]*api.CertificateIssue, string, error) {
	resp, err := e.ListEndpoint(ctx, ListRequest{
		PageSize:  pageSize,
		PageToken: pageToken,
	})
	if err != nil {
		return nil, "", err
	}
	response := resp.(ListResponse)
	return response.CertificateIssues, response.NextPageToken, nil
}

func (e Endpoints) Update(ctx context.Context, name string, certificateIssue *api.CertificateIssue) (*api.CertificateIssue, error) {
	resp, err := e.UpdateEndpoint(ctx, UpdateRequest{
		Name:             name,
		CertificateIssue: certificateIssue,
	})
	if err != nil {
		return nil, err
	}
	response := resp.(UpdateResponse)
	return response.CertificateIssue, nil
}

type CreateRequest struct {
	CertificateIssue *api.CertificateIssue
}

type CreateResponse struct {
	CertificateIssue *api.CertificateIssue
}

func makeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		certificateIssue, err := s.Create(ctx, req.CertificateIssue)
		if err != nil {
			return nil, err
		}
		return CreateResponse{
			CertificateIssue: certificateIssue,
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
	CertificateIssue *api.CertificateIssue
}

func makeGetEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRequest)
		certificateIssue, err := s.Get(ctx, req.Name)
		if err != nil {
			return nil, err
		}
		return GetResponse{
			CertificateIssue: certificateIssue,
		}, nil
	}
}

type ListRequest struct {
	PageSize  int
	PageToken string
}

type ListResponse struct {
	CertificateIssues []*api.CertificateIssue
	NextPageToken     string
}

func makeListEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ListRequest)
		certificateIssues, nextPageToken, err := s.List(ctx, req.PageSize, req.PageToken)
		if err != nil {
			return nil, err
		}
		return ListResponse{
			CertificateIssues: certificateIssues,
			NextPageToken:     nextPageToken,
		}, nil
	}
}

type UpdateRequest struct {
	Name             string
	CertificateIssue *api.CertificateIssue
}

type UpdateResponse struct {
	CertificateIssue *api.CertificateIssue
}

func makeUpdateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateRequest)
		certificateIssue, err := s.Update(ctx, req.Name, req.CertificateIssue)
		if err != nil {
			return nil, err
		}
		return UpdateResponse{
			CertificateIssue: certificateIssue,
		}, nil
	}
}
