package transport // import "powerssl.dev/sdk/apiserver/user/transport"

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
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

func encodeGRPCUsers(users []*api.User) ([]*apiv1.User, error) {
	grpcUsers := make([]*apiv1.User, len(users))
	for i, user := range users {
		grpcUser, err := encodeGRPCUser(user)
		if err != nil {
			return nil, err
		}
		grpcUsers[i] = grpcUser
	}
	return grpcUsers, nil
}

func decodeGRPCCreateRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.CreateUserRequest)
	user, err := decodeGRPCUser(req.GetUser())
	if err != nil {
		return nil, err
	}
	return endpoint.CreateRequest{
		User: user,
	}, nil
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

func encodeGRPCCreateResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.CreateResponse)
	return encodeGRPCUser(resp.User)
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

func decodeGRPCDeleteRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.DeleteUserRequest)
	return endpoint.DeleteRequest{
		Name: req.GetName(),
	}, nil
}

func decodeGRPCDeleteResponse(_ context.Context, _ interface{}) (interface{}, error) {
	return endpoint.DeleteResponse{}, nil
}

func encodeGRPCDeleteResponse(_ context.Context, _ interface{}) (interface{}, error) {
	return &emptypb.Empty{}, nil
}

func encodeGRPCDeleteRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.DeleteRequest)
	return &apiv1.DeleteUserRequest{
		Name: req.Name,
	}, nil
}

func decodeGRPCGetRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.GetUserRequest)
	return endpoint.GetRequest{
		Name: req.GetName(),
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

func encodeGRPCGetResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.GetResponse)
	return encodeGRPCUser(resp.User)
}

func encodeGRPCGetRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.GetRequest)
	return &apiv1.GetUserRequest{
		Name: req.Name,
	}, nil
}

func decodeGRPCListRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.ListUsersRequest)
	return endpoint.ListRequest{
		PageSize:  int(req.GetPageSize()),
		PageToken: req.GetPageToken(),
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

func encodeGRPCListResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.ListResponse)
	users, err := encodeGRPCUsers(resp.Users)
	if err != nil {
		return nil, err
	}
	return &apiv1.ListUsersResponse{
		Users:         users,
		NextPageToken: resp.NextPageToken,
	}, nil
}

func encodeGRPCListRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.ListRequest)
	return &apiv1.ListUsersRequest{
		PageSize:  int32(req.PageSize),
		PageToken: req.PageToken,
	}, nil
}

func decodeGRPCUpdateRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.UpdateUserRequest)
	user, err := decodeGRPCUser(req.GetUser())
	if err != nil {
		return nil, err
	}
	return endpoint.UpdateRequest{
		Name: req.GetName(),
		User: user,
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

func encodeGRPCUpdateResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.UpdateResponse)
	return encodeGRPCUser(resp.User)
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
