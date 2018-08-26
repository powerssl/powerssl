package domainendpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"

	domainmodel "powerssl.io/pkg/domain/model"
	domainservice "powerssl.io/pkg/domain/service"
)

type Set struct {
	CreateEndpoint endpoint.Endpoint
	DeleteEndpoint endpoint.Endpoint
	GetEndpoint    endpoint.Endpoint
	ListEndpoint   endpoint.Endpoint
	UpdateEndpoint endpoint.Endpoint
}

func New(svc domainservice.Service, logger log.Logger) Set {
	var createEndpoint endpoint.Endpoint
	{
		createEndpoint = MakeCreateEndpoint(svc)
		createEndpoint = LoggingMiddleware(log.With(logger, "method", "Create"))(createEndpoint)
	}
	var deleteEndpoint endpoint.Endpoint
	{
		deleteEndpoint = MakeDeleteEndpoint(svc)
		deleteEndpoint = LoggingMiddleware(log.With(logger, "method", "Delete"))(deleteEndpoint)
	}
	var getEndpoint endpoint.Endpoint
	{
		getEndpoint = MakeGetEndpoint(svc)
		getEndpoint = LoggingMiddleware(log.With(logger, "method", "Get"))(getEndpoint)
	}
	var listEndpoint endpoint.Endpoint
	{
		listEndpoint = MakeListEndpoint(svc)
		listEndpoint = LoggingMiddleware(log.With(logger, "method", "List"))(listEndpoint)
	}
	var updateEndpoint endpoint.Endpoint
	{
		updateEndpoint = MakeUpdateEndpoint(svc)
		updateEndpoint = LoggingMiddleware(log.With(logger, "method", "Update"))(updateEndpoint)
	}
	return Set{
		CreateEndpoint: createEndpoint,
		DeleteEndpoint: deleteEndpoint,
		GetEndpoint:    getEndpoint,
		ListEndpoint:   listEndpoint,
		UpdateEndpoint: updateEndpoint,
	}
}

func (s Set) Create(ctx context.Context, domain domainmodel.Domain) (domainmodel.Domain, error) {
	resp, err := s.CreateEndpoint(ctx, CreateRequest{Domain: domain})
	if err != nil {
		return domainmodel.Domain{}, err
	}
	response := resp.(CreateResponse)
	return response.Domain, response.Err
}

func (s Set) Delete(ctx context.Context, name string) error {
	resp, err := s.DeleteEndpoint(ctx, DeleteRequest{Name: name})
	if err != nil {
		return err
	}
	response := resp.(DeleteResponse)
	return response.Err
}

func (s Set) Get(ctx context.Context, name string) (domainmodel.Domain, error) {
	resp, err := s.GetEndpoint(ctx, GetRequest{Name: name})
	if err != nil {
		return domainmodel.Domain{}, err
	}
	response := resp.(GetResponse)
	return response.Domain, response.Err
}

func (s Set) List(ctx context.Context, pageSize int, pageToken string) ([]domainmodel.Domain, error) {
	resp, err := s.ListEndpoint(ctx, ListRequest{
		PageSize:  pageSize,
		PageToken: pageToken,
	})
	if err != nil {
		return []domainmodel.Domain{}, err
	}
	response := resp.(ListResponse)
	return response.Domains, response.Err
}

func (s Set) Update(ctx context.Context, domain domainmodel.Domain, updateMask string) (domainmodel.Domain, error) { // TODO: Update Mask
	resp, err := s.UpdateEndpoint(ctx, UpdateRequest{
		Domain:     domain,
		UpdateMask: updateMask,
	})
	if err != nil {
		return domainmodel.Domain{}, err
	}
	response := resp.(UpdateResponse)
	return response.Domain, response.Err
}

func MakeCreateEndpoint(s domainservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CreateRequest)
		domain, err := s.Create(ctx, req.Domain)
		return CreateResponse{Domain: domain, Err: err}, nil
	}
}

func MakeDeleteEndpoint(s domainservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(DeleteRequest)
		err = s.Delete(ctx, req.Name)
		return DeleteResponse{Err: err}, nil
	}
}

func MakeGetEndpoint(s domainservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetRequest)
		domain, err := s.Get(ctx, req.Name)
		return GetResponse{Domain: domain, Err: err}, nil
	}
}

func MakeListEndpoint(s domainservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(ListRequest)
		domains, err := s.List(ctx, req.PageSize, req.PageToken)
		return ListResponse{Domains: domains, Err: err}, nil
	}
}

func MakeUpdateEndpoint(s domainservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UpdateRequest)
		domain, err := s.Update(ctx, req.Domain, req.UpdateMask)
		return UpdateResponse{Domain: domain, Err: err}, nil
	}
}

type Failer interface {
	Failed() error
}

type CreateRequest struct {
	Domain domainmodel.Domain
}

type CreateResponse struct {
	Domain domainmodel.Domain
	Err    error
}

func (r CreateResponse) Failed() error { return r.Err }

type DeleteRequest struct {
	Name string
}

type DeleteResponse struct {
	Err error
}

func (r DeleteResponse) Failed() error { return r.Err }

type GetRequest struct {
	Name string
}

type GetResponse struct {
	Domain domainmodel.Domain
	Err    error
}

func (r GetResponse) Failed() error { return r.Err }

type ListRequest struct {
	PageSize  int
	PageToken string
}

type ListResponse struct {
	Domains []domainmodel.Domain
	Err     error
}

func (r ListResponse) Failed() error { return r.Err }

type UpdateRequest struct {
	Domain     domainmodel.Domain
	UpdateMask string // TODO: Update Mask
}

type UpdateResponse struct {
	Domain domainmodel.Domain
	Err    error
}

func (r UpdateResponse) Failed() error { return r.Err }
