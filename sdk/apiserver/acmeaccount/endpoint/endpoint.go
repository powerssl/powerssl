package endpoint // import "powerssl.dev/sdk/apiserver/acmeaccount/endpoint"

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"powerssl.dev/sdk/apiserver/acmeaccount"
	"powerssl.dev/sdk/apiserver/api"
)

type Endpoints struct {
	CreateEndpoint endpoint.Endpoint
	DeleteEndpoint endpoint.Endpoint
	GetEndpoint    endpoint.Endpoint
	ListEndpoint   endpoint.Endpoint
	UpdateEndpoint endpoint.Endpoint
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

func (e Endpoints) Update(ctx context.Context, name string, updateMask []string, acmeAccount *api.ACMEAccount) (*api.ACMEAccount, error) {
	resp, err := e.UpdateEndpoint(ctx, UpdateRequest{
		Name:        name,
		UpdateMask:  updateMask,
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

type DeleteRequest struct {
	Name string
}

type DeleteResponse struct{}

type GetRequest struct {
	Name string
}

type GetResponse struct {
	ACMEAccount *api.ACMEAccount
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

type UpdateRequest struct {
	Name        string
	UpdateMask  []string
	ACMEAccount *api.ACMEAccount
}

type UpdateResponse struct {
	ACMEAccount *api.ACMEAccount
}

func MakeCreateEndpoint(s acmeaccount.Service) endpoint.Endpoint {
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

func MakeDeleteEndpoint(s acmeaccount.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRequest)
		if err := s.Delete(ctx, req.Name); err != nil {
			return nil, err
		}
		return DeleteResponse{}, nil
	}
}

func MakeGetEndpoint(s acmeaccount.Service) endpoint.Endpoint {
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

func MakeListEndpoint(s acmeaccount.Service) endpoint.Endpoint {
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

func MakeUpdateEndpoint(s acmeaccount.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateRequest)
		acmeAccount, err := s.Update(ctx, req.Name, req.UpdateMask, req.ACMEAccount)
		if err != nil {
			return nil, err
		}
		return UpdateResponse{
			ACMEAccount: acmeAccount,
		}, nil
	}
}
