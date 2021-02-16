package endpoint // import "powerssl.dev/sdk/apiserver/user/endpoint"

import (
	"context"
	"powerssl.dev/sdk/apiserver/user"

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

func (e Endpoints) Create(ctx context.Context, user *api.User) (*api.User, error) {
	resp, err := e.CreateEndpoint(ctx, CreateRequest{
		User: user,
	})
	if err != nil {
		return nil, err
	}
	response := resp.(CreateResponse)
	return response.User, nil
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

func (e Endpoints) Get(ctx context.Context, name string) (*api.User, error) {
	resp, err := e.GetEndpoint(ctx, GetRequest{
		Name: name,
	})
	if err != nil {
		return nil, err
	}
	response := resp.(GetResponse)
	return response.User, nil
}

func (e Endpoints) List(ctx context.Context, pageSize int, pageToken string) ([]*api.User, string, error) {
	resp, err := e.ListEndpoint(ctx, ListRequest{
		PageSize:  pageSize,
		PageToken: pageToken,
	})
	if err != nil {
		return nil, "", err
	}
	response := resp.(ListResponse)
	return response.Users, response.NextPageToken, nil
}

func (e Endpoints) Update(ctx context.Context, name string, user *api.User) (*api.User, error) {
	resp, err := e.UpdateEndpoint(ctx, UpdateRequest{
		Name: name,
		User: user,
	})
	if err != nil {
		return nil, err
	}
	response := resp.(UpdateResponse)
	return response.User, nil
}

type CreateRequest struct {
	User *api.User
}

type CreateResponse struct {
	User *api.User
}

type DeleteRequest struct {
	Name string
}

type DeleteResponse struct{}

type GetRequest struct {
	Name string
}

type GetResponse struct {
	User *api.User
}

type ListRequest struct {
	PageSize  int
	PageToken string
}

type ListResponse struct {
	Users         []*api.User
	NextPageToken string
}

type UpdateRequest struct {
	Name string
	User *api.User
}

type UpdateResponse struct {
	User *api.User
}

func MakeCreateEndpoint(s user.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		user, err := s.Create(ctx, req.User)
		if err != nil {
			return nil, err
		}
		return CreateResponse{
			User: user,
		}, nil
	}
}

func MakeDeleteEndpoint(s user.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRequest)
		if err := s.Delete(ctx, req.Name); err != nil {
			return nil, err
		}
		return DeleteResponse{}, nil
	}
}

func MakeGetEndpoint(s user.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRequest)
		user, err := s.Get(ctx, req.Name)
		if err != nil {
			return nil, err
		}
		return GetResponse{
			User: user,
		}, nil
	}
}

func MakeListEndpoint(s user.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ListRequest)
		users, nextPageToken, err := s.List(ctx, req.PageSize, req.PageToken)
		if err != nil {
			return nil, err
		}
		return ListResponse{
			Users:         users,
			NextPageToken: nextPageToken,
		}, nil
	}
}

func MakeUpdateEndpoint(s user.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateRequest)
		user, err := s.Update(ctx, req.Name, req.User)
		if err != nil {
			return nil, err
		}
		return UpdateResponse{
			User: user,
		}, nil
	}
}
