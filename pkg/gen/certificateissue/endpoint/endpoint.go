// Code generated by protoc-gen-gotemplate. DO NOT EDIT.

package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"

	"powerssl.io/pkg/api"
	"powerssl.io/pkg/gen/certificateissue/service"
	"powerssl.io/pkg/resources/endpoints"
)

type Endpoints struct {
	CreateEndpoint endpoint.Endpoint
	DeleteEndpoint endpoint.Endpoint
	GetEndpoint    endpoint.Endpoint
	ListEndpoint   endpoint.Endpoint
	UpdateEndpoint endpoint.Endpoint
}

func New(svc service.Service, logger log.Logger, duration metrics.Histogram) Endpoints {
	var createEndpoint endpoint.Endpoint
	{
		createEndpoint = makeCreateEndpoint(svc)
		createEndpoint = endpoints.LoggingMiddleware(log.With(logger, "method", "Create"))(createEndpoint)
		createEndpoint = endpoints.InstrumentingMiddleware(duration.With("method", "Create"))(createEndpoint)
	}

	var deleteEndpoint endpoint.Endpoint
	{
		deleteEndpoint = makeDeleteEndpoint(svc)
		deleteEndpoint = endpoints.LoggingMiddleware(log.With(logger, "method", "Delete"))(deleteEndpoint)
		deleteEndpoint = endpoints.InstrumentingMiddleware(duration.With("method", "Delete"))(deleteEndpoint)
	}

	var getEndpoint endpoint.Endpoint
	{
		getEndpoint = makeGetEndpoint(svc)
		getEndpoint = endpoints.LoggingMiddleware(log.With(logger, "method", "Get"))(getEndpoint)
		getEndpoint = endpoints.InstrumentingMiddleware(duration.With("method", "Get"))(getEndpoint)
	}

	var listEndpoint endpoint.Endpoint
	{
		listEndpoint = makeListEndpoint(svc)
		listEndpoint = endpoints.LoggingMiddleware(log.With(logger, "method", "List"))(listEndpoint)
		listEndpoint = endpoints.InstrumentingMiddleware(duration.With("method", "List"))(listEndpoint)
	}

	var updateEndpoint endpoint.Endpoint
	{
		updateEndpoint = makeUpdateEndpoint(svc)
		updateEndpoint = endpoints.LoggingMiddleware(log.With(logger, "method", "Update"))(updateEndpoint)
		updateEndpoint = endpoints.InstrumentingMiddleware(duration.With("method", "Update"))(updateEndpoint)
	}

	return Endpoints{
		CreateEndpoint: createEndpoint,
		DeleteEndpoint: deleteEndpoint,
		GetEndpoint:    getEndpoint,
		ListEndpoint:   listEndpoint,
		UpdateEndpoint: updateEndpoint,
	}
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

type DeleteResponse struct {
}

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
}

type ListResponse struct {
	CertificateIssues []*api.CertificateIssue
}

func makeListEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		certificateIssues, err := s.List(ctx)
		if err != nil {
			return nil, err
		}
		return ListResponse{
			CertificateIssues: certificateIssues,
		}, nil
	}
}

type UpdateRequest struct {
	CertificateIssue *api.CertificateIssue
}

type UpdateResponse struct {
	CertificateIssue *api.CertificateIssue
}

func makeUpdateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateRequest)
		certificateIssue, err := s.Update(ctx, req.CertificateIssue)
		if err != nil {
			return nil, err
		}
		return UpdateResponse{
			CertificateIssue: certificateIssue,
		}, nil
	}
}
