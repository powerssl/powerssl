package user

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	apiv1 "powerssl.dev/api/apiserver/v1"
)

var ServiceDesc = &apiv1.UserService_ServiceDesc

type Service struct {
	apiv1.UnimplementedUserServiceServer
}

func New() *Service {
	return &Service{}
}

func (s Service) Create(ctx context.Context, request *apiv1.CreateUserRequest) (*apiv1.User, error) {
	panic("implement me")
}

func (s Service) Delete(ctx context.Context, request *apiv1.DeleteUserRequest) (*emptypb.Empty, error) {
	panic("implement me")
}

func (s Service) Get(ctx context.Context, request *apiv1.GetUserRequest) (*apiv1.User, error) {
	panic("implement me")
}

func (s Service) List(ctx context.Context, request *apiv1.ListUsersRequest) (*apiv1.ListUsersResponse, error) {
	panic("implement me")
}

func (s Service) Update(ctx context.Context, request *apiv1.UpdateUserRequest) (*apiv1.User, error) {
	panic("implement me")
}
