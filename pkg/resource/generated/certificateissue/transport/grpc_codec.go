// Code generated by protoc-gen-gotemplate. DO NOT EDIT.

package transport // import "powerssl.io/pkg/resource/generated/certificateissue/transport"

import (
	"context"
	"errors"

	"github.com/gogo/protobuf/types"

	"powerssl.io/pkg/api"
	apiv1 "powerssl.io/pkg/api/v1"
	"powerssl.io/pkg/resource/generated/certificateissue/endpoint"
)

// Avoid import errors
var _ = types.Timestamp{}

var UnknownError = errors.New("Unknown Error")

func decodeGRPCCertificateIssue(certificateIssue *apiv1.CertificateIssue) (*api.CertificateIssue, error) {
	createTime, err := types.TimestampFromProto(certificateIssue.GetCreateTime())
	if err != nil {
		return nil, err
	}
	updateTime, err := types.TimestampFromProto(certificateIssue.GetUpdateTime())
	if err != nil {
		return nil, err
	}
	return &api.CertificateIssue{
		Name:            certificateIssue.GetName(),
		CreateTime:      createTime,
		UpdateTime:      updateTime,
		DisplayName:     certificateIssue.GetDisplayName(),
		Title:           certificateIssue.GetTitle(),
		Description:     certificateIssue.GetDescription(),
		Labels:          certificateIssue.GetLabels(),
		Dnsnames:        certificateIssue.GetDnsnames(),
		KeyAlgorithm:    certificateIssue.GetKeyAlgorithm(),
		KeySize:         certificateIssue.GetKeySize(),
		DigestAlgorithm: certificateIssue.GetDigestAlgorithm(),
		AutoRenew:       certificateIssue.GetAutoRenew(),
	}, nil
}

func encodeGRPCCertificateIssue(certificateIssue *api.CertificateIssue) (*apiv1.CertificateIssue, error) {
	if certificateIssue == nil {
		return nil, UnknownError
	}
	createTime, err := types.TimestampProto(certificateIssue.CreateTime)
	if err != nil {
		return nil, err
	}
	updateTime, err := types.TimestampProto(certificateIssue.UpdateTime)
	if err != nil {
		return nil, err
	}
	return &apiv1.CertificateIssue{
		Name:            certificateIssue.Name,
		CreateTime:      createTime,
		UpdateTime:      updateTime,
		DisplayName:     certificateIssue.DisplayName,
		Title:           certificateIssue.Title,
		Description:     certificateIssue.Description,
		Labels:          certificateIssue.Labels,
		Dnsnames:        certificateIssue.Dnsnames,
		KeyAlgorithm:    certificateIssue.KeyAlgorithm,
		KeySize:         certificateIssue.KeySize,
		DigestAlgorithm: certificateIssue.DigestAlgorithm,
		AutoRenew:       certificateIssue.AutoRenew,
	}, nil
}

func decodeGRPCCertificateIssues(grpcCertificateIssues []*apiv1.CertificateIssue) ([]*api.CertificateIssue, error) {
	certificateIssues := make([]*api.CertificateIssue, len(grpcCertificateIssues))
	for i, grpcCertificateIssue := range grpcCertificateIssues {
		certificateIssue, err := decodeGRPCCertificateIssue(grpcCertificateIssue)
		if err != nil {
			return nil, err
		}
		certificateIssues[i] = certificateIssue
	}
	return certificateIssues, nil
}

func encodeGRPCCertificateIssues(certificateIssues []*api.CertificateIssue) ([]*apiv1.CertificateIssue, error) {
	grpcCertificateIssues := make([]*apiv1.CertificateIssue, len(certificateIssues))
	for i, certificateIssue := range certificateIssues {
		grpcCertificateIssue, err := encodeGRPCCertificateIssue(certificateIssue)
		if err != nil {
			return nil, err
		}
		grpcCertificateIssues[i] = grpcCertificateIssue
	}
	return grpcCertificateIssues, nil
}

func decodeGRPCCreateRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.CreateCertificateIssueRequest)
	certificateIssue, err := decodeGRPCCertificateIssue(req.GetCertificateIssue())
	if err != nil {
		return nil, err
	}
	return endpoint.CreateRequest{
		Parent:           req.GetParent(),
		CertificateIssue: certificateIssue,
	}, nil
}

func decodeGRPCCreateResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.CertificateIssue)
	certificateIssue, err := decodeGRPCCertificateIssue(reply)
	if err != nil {
		return nil, err
	}
	return endpoint.CreateResponse{
		CertificateIssue: certificateIssue,
	}, nil
}

func encodeGRPCCreateResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.CreateResponse)
	return encodeGRPCCertificateIssue(resp.CertificateIssue)
}

func encodeGRPCCreateRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.CreateRequest)
	certificateIssue, err := encodeGRPCCertificateIssue(req.CertificateIssue)
	if err != nil {
		return nil, err
	}
	return &apiv1.CreateCertificateIssueRequest{
		Parent:           req.Parent,
		CertificateIssue: certificateIssue,
	}, nil
}

func decodeGRPCDeleteRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.DeleteCertificateIssueRequest)
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
	return &apiv1.DeleteCertificateIssueRequest{
		Name: req.Name,
	}, nil
}

func decodeGRPCGetRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.GetCertificateIssueRequest)
	return endpoint.GetRequest{
		Name: req.GetName(),
	}, nil
}

func decodeGRPCGetResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.CertificateIssue)
	certificateIssue, err := decodeGRPCCertificateIssue(reply)
	if err != nil {
		return nil, err
	}
	return endpoint.GetResponse{
		CertificateIssue: certificateIssue,
	}, nil
}

func encodeGRPCGetResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.GetResponse)
	return encodeGRPCCertificateIssue(resp.CertificateIssue)
}

func encodeGRPCGetRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.GetRequest)
	return &apiv1.GetCertificateIssueRequest{
		Name: req.Name,
	}, nil
}

func decodeGRPCListRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.ListCertificateIssuesRequest)
	return endpoint.ListRequest{
		Parent:    req.GetParent(),
		PageSize:  int(req.GetPageSize()),
		PageToken: req.GetPageToken(),
	}, nil
}

func decodeGRPCListResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.ListCertificateIssuesResponse)
	certificateIssues, err := decodeGRPCCertificateIssues(reply.GetCertificateIssues())
	if err != nil {
		return nil, err
	}
	return endpoint.ListResponse{
		CertificateIssues: certificateIssues,
		NextPageToken:     reply.GetNextPageToken(),
	}, nil
}

func encodeGRPCListResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.ListResponse)
	certificateIssues, err := encodeGRPCCertificateIssues(resp.CertificateIssues)
	if err != nil {
		return nil, err
	}
	return &apiv1.ListCertificateIssuesResponse{
		CertificateIssues: certificateIssues,
		NextPageToken:     resp.NextPageToken,
	}, nil
}

func encodeGRPCListRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.ListRequest)
	return &apiv1.ListCertificateIssuesRequest{
		Parent:    req.Parent,
		PageSize:  int32(req.PageSize),
		PageToken: req.PageToken,
	}, nil
}

func decodeGRPCUpdateRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.UpdateCertificateIssueRequest)
	certificateIssue, err := decodeGRPCCertificateIssue(req.GetCertificateIssue())
	if err != nil {
		return nil, err
	}
	return endpoint.UpdateRequest{
		Name:             req.GetName(),
		CertificateIssue: certificateIssue,
	}, nil
}

func decodeGRPCUpdateResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.CertificateIssue)
	certificateIssue, err := decodeGRPCCertificateIssue(reply)
	if err != nil {
		return nil, err
	}
	return endpoint.UpdateResponse{
		CertificateIssue: certificateIssue,
	}, nil
}

func encodeGRPCUpdateResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.UpdateResponse)
	return encodeGRPCCertificateIssue(resp.CertificateIssue)
}

func encodeGRPCUpdateRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.UpdateRequest)
	certificateIssue, err := encodeGRPCCertificateIssue(req.CertificateIssue)
	if err != nil {
		return nil, err
	}
	return &apiv1.UpdateCertificateIssueRequest{
		Name:             req.Name,
		CertificateIssue: certificateIssue,
	}, nil
}
