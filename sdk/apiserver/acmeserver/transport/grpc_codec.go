package transport // import "powerssl.dev/sdk/apiserver/acmeserver/transport"

import (
	"context"

	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	apiv1 "powerssl.dev/api/apiserver/v1"

	"powerssl.dev/sdk/apiserver/acmeserver/endpoint"
	"powerssl.dev/sdk/apiserver/api"
)

func decodeGRPCACMEServer(acmeServer *apiv1.ACMEServer) (*api.ACMEServer, error) {
	return &api.ACMEServer{
		Name:            acmeServer.GetName(),
		CreateTime:      acmeServer.GetCreateTime().AsTime(),
		UpdateTime:      acmeServer.GetUpdateTime().AsTime(),
		DisplayName:     acmeServer.GetDisplayName(),
		DirectoryURL:    acmeServer.GetDirectoryUrl(),
		IntegrationName: acmeServer.GetIntegrationName(),
	}, nil
}

func encodeGRPCACMEServer(acmeServer *api.ACMEServer) (*apiv1.ACMEServer, error) {
	return &apiv1.ACMEServer{
		Name:            acmeServer.Name,
		CreateTime:      timestamppb.New(acmeServer.CreateTime),
		UpdateTime:      timestamppb.New(acmeServer.UpdateTime),
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

func decodeGRPCDeleteResponse(_ context.Context, _ interface{}) (interface{}, error) {
	return endpoint.DeleteResponse{}, nil
}

func encodeGRPCDeleteRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.DeleteRequest)
	return &apiv1.DeleteACMEServerRequest{
		Name: req.Name,
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

func encodeGRPCGetRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.GetRequest)
	return &apiv1.GetACMEServerRequest{
		Name: req.Name,
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

func encodeGRPCListRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.ListRequest)
	return &apiv1.ListACMEServersRequest{
		PageSize:  int32(req.PageSize),
		PageToken: req.PageToken,
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

func encodeGRPCUpdateRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.UpdateRequest)
	acmeServer, err := encodeGRPCACMEServer(req.ACMEServer)
	if err != nil {
		return nil, err
	}
	var messageType *apiv1.ACMEServer
	updateMask, err := fieldmaskpb.New(messageType, req.UpdateMask...)
	if err != nil {
		return nil, err
	}
	return &apiv1.UpdateACMEServerRequest{
		Name:       req.Name,
		UpdateMask: updateMask,
		AcmeServer: acmeServer,
	}, nil
}
