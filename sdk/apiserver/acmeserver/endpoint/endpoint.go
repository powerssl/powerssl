package endpoint // import "powerssl.dev/sdk/apiserver/acmeserver/endpoint"

import (
	"context"
	"powerssl.dev/sdk/apiserver/acmeserver"

	"github.com/go-kit/kit/endpoint"

	"powerssl.dev/sdk/apiserver/api"
)

type Endpoints struct {
	CreateEndpoint endpoint.Endpoint
	DeleteEndpoint endpoint.Endpoint
	GetEndpoint    endpoint.Endpoint
	ListEndpoint   endpoint.Endpoint
	UpdateEndpoint endpoint.Endpoint
}

func (e Endpoints) Create(ctx context.Context, acmeServer *api.ACMEServer) (*api.ACMEServer, error) {
	resp, err := e.CreateEndpoint(ctx, CreateRequest{
		ACMEServer: acmeServer,
	})
	if err != nil {
		return nil, err
	}
	response := resp.(CreateResponse)
	return response.ACMEServer, nil
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

func (e Endpoints) Get(ctx context.Context, name string) (*api.ACMEServer, error) {
	resp, err := e.GetEndpoint(ctx, GetRequest{
		Name: name,
	})
	if err != nil {
		return nil, err
	}
	response := resp.(GetResponse)
	return response.ACMEServer, nil
}

func (e Endpoints) List(ctx context.Context, pageSize int, pageToken string) ([]*api.ACMEServer, string, error) {
	resp, err := e.ListEndpoint(ctx, ListRequest{
		PageSize:  pageSize,
		PageToken: pageToken,
	})
	if err != nil {
		return nil, "", err
	}
	response := resp.(ListResponse)
	return response.ACMEServers, response.NextPageToken, nil
}

func (e Endpoints) Update(ctx context.Context, name string, updateMask []string, acmeServer *api.ACMEServer) (*api.ACMEServer, error) {
	resp, err := e.UpdateEndpoint(ctx, UpdateRequest{
		Name:       name,
		UpdateMask: updateMask,
		ACMEServer: acmeServer,
	})
	if err != nil {
		return nil, err
	}
	response := resp.(UpdateResponse)
	return response.ACMEServer, nil
}

type CreateRequest struct {
	ACMEServer *api.ACMEServer
}

type CreateResponse struct {
	ACMEServer *api.ACMEServer
}

type DeleteRequest struct {
	Name string
}

type DeleteResponse struct{}

type GetRequest struct {
	Name string
}

type GetResponse struct {
	ACMEServer *api.ACMEServer
}

type ListRequest struct {
	PageSize  int
	PageToken string
}

type ListResponse struct {
	ACMEServers   []*api.ACMEServer
	NextPageToken string
}

type UpdateRequest struct {
	Name       string
	UpdateMask  []string
	ACMEServer *api.ACMEServer
}

type UpdateResponse struct {
	ACMEServer *api.ACMEServer
}

func MakeCreateEndpoint(s acmeserver.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		acmeServer, err := s.Create(ctx, req.ACMEServer)
		if err != nil {
			return nil, err
		}
		return CreateResponse{
			ACMEServer: acmeServer,
		}, nil
	}
}

func MakeDeleteEndpoint(s acmeserver.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRequest)
		if err := s.Delete(ctx, req.Name); err != nil {
			return nil, err
		}
		return DeleteResponse{}, nil
	}
}

func MakeGetEndpoint(s acmeserver.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRequest)
		acmeServer, err := s.Get(ctx, req.Name)
		if err != nil {
			return nil, err
		}
		return GetResponse{
			ACMEServer: acmeServer,
		}, nil
	}
}

func MakeListEndpoint(s acmeserver.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ListRequest)
		acmeServers, nextPageToken, err := s.List(ctx, req.PageSize, req.PageToken)
		if err != nil {
			return nil, err
		}
		return ListResponse{
			ACMEServers:   acmeServers,
			NextPageToken: nextPageToken,
		}, nil
	}
}

func MakeUpdateEndpoint(s acmeserver.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateRequest)
		acmeServer, err := s.Update(ctx, req.Name, req.UpdateMask, req.ACMEServer)
		if err != nil {
			return nil, err
		}
		return UpdateResponse{
			ACMEServer: acmeServer,
		}, nil
	}
}
