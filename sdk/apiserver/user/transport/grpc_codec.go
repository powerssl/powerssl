package transport // import "powerssl.dev/sdk/apiserver/user/transport"

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	apiv1 "powerssl.dev/api/apiserver/v1"

	"powerssl.dev/sdk/apiserver/api"
	"powerssl.dev/sdk/apiserver/user/endpoint"
)

func decodeGRPCUser(user *apiv1.User) (*api.User, error) {
	return &api.User{
		Name:        user.GetName(),
		CreateTime:  user.GetCreateTime().AsTime(),
		UpdateTime:  user.GetUpdateTime().AsTime(),
		DisplayName: user.GetDisplayName(),
		UserName:    user.GetUserName(),
	}, nil
}

func encodeGRPCUser(user *api.User) (*apiv1.User, error) {
	return &apiv1.User{
		Name:        user.Name,
		CreateTime:  timestamppb.New(user.CreateTime),
		UpdateTime:  timestamppb.New(user.UpdateTime),
		DisplayName: user.DisplayName,
		UserName:    user.UserName,
	}, nil
}

func decodeGRPCUsers(grpcUsers []*apiv1.User) ([]*api.User, error) {
	users := make([]*api.User, len(grpcUsers))
	for i, grpcUser := range grpcUsers {
		user, err := decodeGRPCUser(grpcUser)
		if err != nil {
			return nil, err
		}
		users[i] = user
	}
	return users, nil
}

func decodeGRPCCreateResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.User)
	user, err := decodeGRPCUser(reply)
	if err != nil {
		return nil, err
	}
	return endpoint.CreateResponse{
		User: user,
	}, nil
}

func encodeGRPCCreateRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.CreateRequest)
	user, err := encodeGRPCUser(req.User)
	if err != nil {
		return nil, err
	}
	return &apiv1.CreateUserRequest{
		User: user,
	}, nil
}

func decodeGRPCDeleteResponse(_ context.Context, _ interface{}) (interface{}, error) {
	return endpoint.DeleteResponse{}, nil
}

func encodeGRPCDeleteRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.DeleteRequest)
	return &apiv1.DeleteUserRequest{
		Name: req.Name,
	}, nil
}

func decodeGRPCGetResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.User)
	user, err := decodeGRPCUser(reply)
	if err != nil {
		return nil, err
	}
	return endpoint.GetResponse{
		User: user,
	}, nil
}

func encodeGRPCGetRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.GetRequest)
	return &apiv1.GetUserRequest{
		Name: req.Name,
	}, nil
}

func decodeGRPCListResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.ListUsersResponse)
	users, err := decodeGRPCUsers(reply.GetUsers())
	if err != nil {
		return nil, err
	}
	return endpoint.ListResponse{
		Users:         users,
		NextPageToken: reply.GetNextPageToken(),
	}, nil
}

func encodeGRPCListRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.ListRequest)
	return &apiv1.ListUsersRequest{
		PageSize:  int32(req.PageSize),
		PageToken: req.PageToken,
	}, nil
}

func decodeGRPCUpdateResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.User)
	user, err := decodeGRPCUser(reply)
	if err != nil {
		return nil, err
	}
	return endpoint.UpdateResponse{
		User: user,
	}, nil
}

func encodeGRPCUpdateRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.UpdateRequest)
	user, err := encodeGRPCUser(req.User)
	if err != nil {
		return nil, err
	}
	return &apiv1.UpdateUserRequest{
		Name: req.Name,
		User: user,
	}, nil
}
