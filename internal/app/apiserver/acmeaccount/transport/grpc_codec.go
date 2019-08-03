package transport

import (
	"context"

	"github.com/gogo/protobuf/types"

	"powerssl.dev/powerssl/internal/app/apiserver/acmeaccount/endpoint"
	apiv1 "powerssl.dev/powerssl/internal/pkg/apiserver/api/v1"
	"powerssl.dev/powerssl/pkg/apiserver/api"
)

func decodeGRPCACMEAccount(acmeAccount *apiv1.ACMEAccount) (*api.ACMEAccount, error) {
	createTime, err := types.TimestampFromProto(acmeAccount.GetCreateTime())
	if err != nil {
		return nil, err
	}
	updateTime, err := types.TimestampFromProto(acmeAccount.GetUpdateTime())
	if err != nil {
		return nil, err
	}
	return &api.ACMEAccount{
		Name:                 acmeAccount.GetName(),
		CreateTime:           createTime,
		UpdateTime:           updateTime,
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
	createTime, err := types.TimestampProto(acmeAccount.CreateTime)
	if err != nil {
		return nil, err
	}
	updateTime, err := types.TimestampProto(acmeAccount.UpdateTime)
	if err != nil {
		return nil, err
	}
	return &apiv1.ACMEAccount{
		Name:                 acmeAccount.Name,
		CreateTime:           createTime,
		UpdateTime:           updateTime,
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

func encodeGRPCACMEAccounts(acmeAccounts []*api.ACMEAccount) ([]*apiv1.ACMEAccount, error) {
	grpcACMEAccounts := make([]*apiv1.ACMEAccount, len(acmeAccounts))
	for i, acmeAccount := range acmeAccounts {
		grpcACMEAccount, err := encodeGRPCACMEAccount(acmeAccount)
		if err != nil {
			return nil, err
		}
		grpcACMEAccounts[i] = grpcACMEAccount
	}
	return grpcACMEAccounts, nil
}

func decodeGRPCCreateRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.CreateACMEAccountRequest)
	acmeAccount, err := decodeGRPCACMEAccount(req.GetAcmeAccount())
	if err != nil {
		return nil, err
	}
	return endpoint.CreateRequest{
		Parent:      req.GetParent(),
		ACMEAccount: acmeAccount,
	}, nil
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

func encodeGRPCCreateResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.CreateResponse)
	return encodeGRPCACMEAccount(resp.ACMEAccount)
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

func decodeGRPCDeleteRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.DeleteACMEAccountRequest)
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
	return &apiv1.DeleteACMEAccountRequest{
		Name: req.Name,
	}, nil
}

func decodeGRPCGetRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.GetACMEAccountRequest)
	return endpoint.GetRequest{
		Name: req.GetName(),
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

func encodeGRPCGetResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.GetResponse)
	return encodeGRPCACMEAccount(resp.ACMEAccount)
}

func encodeGRPCGetRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.GetRequest)
	return &apiv1.GetACMEAccountRequest{
		Name: req.Name,
	}, nil
}

func decodeGRPCListRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.ListACMEAccountsRequest)
	return endpoint.ListRequest{
		Parent:    req.GetParent(),
		PageSize:  int(req.GetPageSize()),
		PageToken: req.GetPageToken(),
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

func encodeGRPCListResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.ListResponse)
	acmeAccounts, err := encodeGRPCACMEAccounts(resp.ACMEAccounts)
	if err != nil {
		return nil, err
	}
	return &apiv1.ListACMEAccountsResponse{
		AcmeAccounts:  acmeAccounts,
		NextPageToken: resp.NextPageToken,
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

func decodeGRPCUpdateRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.UpdateACMEAccountRequest)
	acmeAccount, err := decodeGRPCACMEAccount(req.GetAcmeAccount())
	if err != nil {
		return nil, err
	}
	return endpoint.UpdateRequest{
		Name:        req.GetName(),
		UpdateMask:  req.GetUpdateMask().GetPaths(),
		ACMEAccount: acmeAccount,
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

func encodeGRPCUpdateResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.UpdateResponse)
	return encodeGRPCACMEAccount(resp.ACMEAccount)
}

func encodeGRPCUpdateRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.UpdateRequest)
	acmeAccount, err := encodeGRPCACMEAccount(req.ACMEAccount)
	if err != nil {
		return nil, err
	}
	return &apiv1.UpdateACMEAccountRequest{
		Name:        req.Name,
		UpdateMask:  &types.FieldMask{Paths: req.UpdateMask},
		AcmeAccount: acmeAccount,
	}, nil
}
