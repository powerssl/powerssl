package transport

import (
	"context"

	"github.com/gogo/protobuf/types"

	"powerssl.dev/powerssl/internal/app/apiserver/acmeserver/endpoint"
	apiv1 "powerssl.dev/powerssl/internal/pkg/apiserver/api/v1"
	"powerssl.dev/powerssl/pkg/apiserver/api"
)

func decodeGRPCACMEServer(acmeServer *apiv1.ACMEServer) (*api.ACMEServer, error) {
	createTime, err := types.TimestampFromProto(acmeServer.GetCreateTime())
	if err != nil {
		return nil, err
	}
	updateTime, err := types.TimestampFromProto(acmeServer.GetUpdateTime())
	if err != nil {
		return nil, err
	}
	return &api.ACMEServer{
		Name:            acmeServer.GetName(),
		CreateTime:      createTime,
		UpdateTime:      updateTime,
		DisplayName:     acmeServer.GetDisplayName(),
		DirectoryURL:    acmeServer.GetDirectoryUrl(),
		IntegrationName: acmeServer.GetIntegrationName(),
	}, nil
}

func encodeGRPCACMEServer(acmeServer *api.ACMEServer) (*apiv1.ACMEServer, error) {
	createTime, err := types.TimestampProto(acmeServer.CreateTime)
	if err != nil {
		return nil, err
	}
	updateTime, err := types.TimestampProto(acmeServer.UpdateTime)
	if err != nil {
		return nil, err
	}
	return &apiv1.ACMEServer{
		Name:            acmeServer.Name,
		CreateTime:      createTime,
		UpdateTime:      updateTime,
		DisplayName:     acmeServer.DisplayName,
		DirectoryUrl:    acmeServer.DirectoryURL,
		IntegrationName: acmeServer.IntegrationName,
	}, nil
}

func decodeGRPCACMEServers(grpcACMEServers []*apiv1.ACMEServer) ([]*api.ACMEServer, error) {
	acmeServers := make([]*api.ACMEServer, len(grpcACMEServers))
	for i, grpcACMEServer := range grpcACMEServers {
		acmeServer, err := decodeGRPCACMEServer(grpcACMEServer)
		if err != nil {
			return nil, err
		}
		acmeServers[i] = acmeServer
	}
	return acmeServers, nil
}

func encodeGRPCACMEServers(acmeServers []*api.ACMEServer) ([]*apiv1.ACMEServer, error) {
	grpcACMEServers := make([]*apiv1.ACMEServer, len(acmeServers))
	for i, acmeServer := range acmeServers {
		grpcACMEServer, err := encodeGRPCACMEServer(acmeServer)
		if err != nil {
			return nil, err
		}
		grpcACMEServers[i] = grpcACMEServer
	}
	return grpcACMEServers, nil
}

func decodeGRPCCreateRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.CreateACMEServerRequest)
	acmeServer, err := decodeGRPCACMEServer(req.GetAcmeServer())
	if err != nil {
		return nil, err
	}
	return endpoint.CreateRequest{
		ACMEServer: acmeServer,
	}, nil
}

func decodeGRPCCreateResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.ACMEServer)
	acmeServer, err := decodeGRPCACMEServer(reply)
	if err != nil {
		return nil, err
	}
	return endpoint.CreateResponse{
		ACMEServer: acmeServer,
	}, nil
}

func encodeGRPCCreateResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.CreateResponse)
	return encodeGRPCACMEServer(resp.ACMEServer)
}

func encodeGRPCCreateRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.CreateRequest)
	acmeServer, err := encodeGRPCACMEServer(req.ACMEServer)
	if err != nil {
		return nil, err
	}
	return &apiv1.CreateACMEServerRequest{
		AcmeServer: acmeServer,
	}, nil
}

func decodeGRPCDeleteRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.DeleteACMEServerRequest)
	return endpoint.DeleteRequest{
		Name: req.GetName(),
	}, nil
}

func decodeGRPCDeleteResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	return endpoint.DeleteResponse{}, nil
}

func encodeGRPCDeleteResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &types.Empty{}, nil
}

func encodeGRPCDeleteRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.DeleteRequest)
	return &apiv1.DeleteACMEServerRequest{
		Name: req.Name,
	}, nil
}

func decodeGRPCGetRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.GetACMEServerRequest)
	return endpoint.GetRequest{
		Name: req.GetName(),
	}, nil
}

func decodeGRPCGetResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.ACMEServer)
	acmeServer, err := decodeGRPCACMEServer(reply)
	if err != nil {
		return nil, err
	}
	return endpoint.GetResponse{
		ACMEServer: acmeServer,
	}, nil
}

func encodeGRPCGetResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.GetResponse)
	return encodeGRPCACMEServer(resp.ACMEServer)
}

func encodeGRPCGetRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.GetRequest)
	return &apiv1.GetACMEServerRequest{
		Name: req.Name,
	}, nil
}

func decodeGRPCListRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.ListACMEServersRequest)
	return endpoint.ListRequest{
		PageSize:  int(req.GetPageSize()),
		PageToken: req.GetPageToken(),
	}, nil
}

func decodeGRPCListResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.ListACMEServersResponse)
	acmeServers, err := decodeGRPCACMEServers(reply.GetAcmeServers())
	if err != nil {
		return nil, err
	}
	return endpoint.ListResponse{
		ACMEServers:   acmeServers,
		NextPageToken: reply.GetNextPageToken(),
	}, nil
}

func encodeGRPCListResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.ListResponse)
	acmeServers, err := encodeGRPCACMEServers(resp.ACMEServers)
	if err != nil {
		return nil, err
	}
	return &apiv1.ListACMEServersResponse{
		AcmeServers:   acmeServers,
		NextPageToken: resp.NextPageToken,
	}, nil
}

func encodeGRPCListRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.ListRequest)
	return &apiv1.ListACMEServersRequest{
		PageSize:  int32(req.PageSize),
		PageToken: req.PageToken,
	}, nil
}

func decodeGRPCUpdateRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.UpdateACMEServerRequest)
	acmeServer, err := decodeGRPCACMEServer(req.GetAcmeServer())
	if err != nil {
		return nil, err
	}
	return endpoint.UpdateRequest{
		Name:       req.GetName(),
		ACMEServer: acmeServer,
	}, nil
}

func decodeGRPCUpdateResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.ACMEServer)
	acmeServer, err := decodeGRPCACMEServer(reply)
	if err != nil {
		return nil, err
	}
	return endpoint.UpdateResponse{
		ACMEServer: acmeServer,
	}, nil
}

func encodeGRPCUpdateResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.UpdateResponse)
	return encodeGRPCACMEServer(resp.ACMEServer)
}

func encodeGRPCUpdateRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.UpdateRequest)
	acmeServer, err := encodeGRPCACMEServer(req.ACMEServer)
	if err != nil {
		return nil, err
	}
	return &apiv1.UpdateACMEServerRequest{
		Name:       req.Name,
		AcmeServer: acmeServer,
	}, nil
}
