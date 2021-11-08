package endpoint // import "powerssl.dev/sdk/apiserver/acmeserver/endpoint"

import (
	"context"

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
	UpdateMask []string
	ACMEServer *api.ACMEServer
}

type UpdateResponse struct {
	ACMEServer *api.ACMEServer
}
