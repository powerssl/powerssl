package endpoint // import "powerssl.dev/sdk/apiserver/certificate/endpoint"

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"powerssl.dev/sdk/apiserver/api"
	"powerssl.dev/sdk/apiserver/certificate"
)

type Endpoints struct {
	CreateEndpoint endpoint.Endpoint
	DeleteEndpoint endpoint.Endpoint
	GetEndpoint    endpoint.Endpoint
	ListEndpoint   endpoint.Endpoint
	UpdateEndpoint endpoint.Endpoint
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

type DeleteRequest struct {
	Name string
}

type DeleteResponse struct{}

type GetRequest struct {
	Name string
}

type GetResponse struct {
	Certificate *api.Certificate
}

type ListRequest struct {
	PageSize  int
	PageToken string
}

type ListResponse struct {
	Certificates  []*api.Certificate
	NextPageToken string
}

type UpdateRequest struct {
	Name        string
	Certificate *api.Certificate
}

type UpdateResponse struct {
	Certificate *api.Certificate
}

func MakeCreateEndpoint(s certificate.Service) endpoint.Endpoint {
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

func MakeDeleteEndpoint(s certificate.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRequest)
		if err := s.Delete(ctx, req.Name); err != nil {
			return nil, err
		}
		return DeleteResponse{}, nil
	}
}

func MakeGetEndpoint(s certificate.Service) endpoint.Endpoint {
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

func MakeListEndpoint(s certificate.Service) endpoint.Endpoint {
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

func MakeUpdateEndpoint(s certificate.Service) endpoint.Endpoint {
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
