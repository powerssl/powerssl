package transport // import "powerssl.dev/sdk/apiserver/acmeaccount/transport"

import (
	"context"

	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	apiv1 "powerssl.dev/api/apiserver/v1"

	"powerssl.dev/sdk/apiserver/acmeaccount/endpoint"
	"powerssl.dev/sdk/apiserver/api"
)

func decodeGRPCACMEAccount(acmeAccount *apiv1.ACMEAccount) (*api.ACMEAccount, error) {
	return &api.ACMEAccount{
		Name:                 acmeAccount.GetName(),
		CreateTime:           acmeAccount.GetCreateTime().AsTime(),
		UpdateTime:           acmeAccount.GetUpdateTime().AsTime(),
		DisplayName:          acmeAccount.GetDisplayName(),
		Title:                acmeAccount.GetTitle(),
		Description:          acmeAccount.GetDescription(),
		Labels:               acmeAccount.GetLabels(),
		TermsOfServiceAgreed: acmeAccount.GetTermsOfServiceAgreed(),
		Contacts:             acmeAccount.GetContacts(),
		AccountURL:           acmeAccount.GetAccountUrl(),
	}, nil
}

func encodeGRPCACMEAccount(acmeAccount *api.ACMEAccount) (*apiv1.ACMEAccount, error) {
	return &apiv1.ACMEAccount{
		Name:                 acmeAccount.Name,
		CreateTime:           timestamppb.New(acmeAccount.CreateTime),
		UpdateTime:           timestamppb.New(acmeAccount.UpdateTime),
		DisplayName:          acmeAccount.DisplayName,
		Title:                acmeAccount.Title,
		Description:          acmeAccount.Description,
		Labels:               acmeAccount.Labels,
		TermsOfServiceAgreed: acmeAccount.TermsOfServiceAgreed,
		Contacts:             acmeAccount.Contacts,
		AccountUrl:           acmeAccount.AccountURL,
	}, nil
}

func decodeGRPCACMEAccounts(grpcACMEAccounts []*apiv1.ACMEAccount) ([]*api.ACMEAccount, error) {
	acmeAccounts := make([]*api.ACMEAccount, len(grpcACMEAccounts))
	for i, grpcACMEAccount := range grpcACMEAccounts {
		acmeAccount, err := decodeGRPCACMEAccount(grpcACMEAccount)
		if err != nil {
			return nil, err
		}
		acmeAccounts[i] = acmeAccount
	}
	return acmeAccounts, nil
}

func decodeGRPCCreateResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.ACMEAccount)
	acmeAccount, err := decodeGRPCACMEAccount(reply)
	if err != nil {
		return nil, err
	}
	return endpoint.CreateResponse{
		ACMEAccount: acmeAccount,
	}, nil
}

func encodeGRPCCreateRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.CreateRequest)
	acmeAccount, err := encodeGRPCACMEAccount(req.ACMEAccount)
	if err != nil {
		return nil, err
	}
	return &apiv1.CreateACMEAccountRequest{
		Parent:      req.Parent,
		AcmeAccount: acmeAccount,
	}, nil
}

func decodeGRPCDeleteResponse(_ context.Context, _ interface{}) (interface{}, error) {
	return endpoint.DeleteResponse{}, nil
}

func encodeGRPCDeleteRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.DeleteRequest)
	return &apiv1.DeleteACMEAccountRequest{
		Name: req.Name,
	}, nil
}

func decodeGRPCGetResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.ACMEAccount)
	acmeAccount, err := decodeGRPCACMEAccount(reply)
	if err != nil {
		return nil, err
	}
	return endpoint.GetResponse{
		ACMEAccount: acmeAccount,
	}, nil
}

func encodeGRPCGetRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.GetRequest)
	return &apiv1.GetACMEAccountRequest{
		Name: req.Name,
	}, nil
}

func decodeGRPCListResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.ListACMEAccountsResponse)
	acmeAccounts, err := decodeGRPCACMEAccounts(reply.GetAcmeAccounts())
	if err != nil {
		return nil, err
	}
	return endpoint.ListResponse{
		ACMEAccounts:  acmeAccounts,
		NextPageToken: reply.GetNextPageToken(),
	}, nil
}

func encodeGRPCListRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.ListRequest)
	return &apiv1.ListACMEAccountsRequest{
		Parent:    req.Parent,
		PageSize:  int32(req.PageSize),
		PageToken: req.PageToken,
	}, nil
}

func decodeGRPCUpdateResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.ACMEAccount)
	acmeAccount, err := decodeGRPCACMEAccount(reply)
	if err != nil {
		return nil, err
	}
	return endpoint.UpdateResponse{
		ACMEAccount: acmeAccount,
	}, nil
}

func encodeGRPCUpdateRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.UpdateRequest)
	acmeAccount, err := encodeGRPCACMEAccount(req.ACMEAccount)
	if err != nil {
		return nil, err
	}
	var messageType *apiv1.ACMEAccount
	updateMask, err := fieldmaskpb.New(messageType, req.UpdateMask...)
	if err != nil {
		return nil, err
	}
	return &apiv1.UpdateACMEAccountRequest{
		Name:        req.Name,
		UpdateMask:  updateMask,
		AcmeAccount: acmeAccount,
	}, nil
}
