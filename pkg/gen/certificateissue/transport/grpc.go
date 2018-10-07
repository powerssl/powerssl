// Code generated by protoc-gen-gotemplate. DO NOT EDIT.

package transport

import (
	"context"

	"powerssl.io/pkg/api"
	apiv1 "powerssl.io/pkg/api/v1"
	"powerssl.io/pkg/gen/certificateissue/endpoint"
	"powerssl.io/pkg/resources/transport"
)

func decodeGRPCCertificateIssues(certificateIssues []*apiv1.CertificateIssue) []*api.CertificateIssue {
	items := make([]*api.CertificateIssue, len(certificateIssues))
	for i, certificateIssue := range certificateIssues {
		items[i] = decodeGRPCCertificateIssue(certificateIssue)
	}
	return items
}

func decodeGRPCCertificateIssue(certificateIssue *apiv1.CertificateIssue) *api.CertificateIssue {
	return &api.CertificateIssue{
		TypeMeta:   transport.DecodeGRPCTypeMeta(certificateIssue.GetTypeMeta()),
		ObjectMeta: transport.DecodeGRPCObjectMeta(certificateIssue.GetObjectMeta()),
		// TODO
	}
}

func encodeGRPCCertificateIssues(certificateIssues []*api.CertificateIssue) []*apiv1.CertificateIssue {
	items := make([]*apiv1.CertificateIssue, len(certificateIssues))
	for i, certificateIssue := range certificateIssues {
		items[i] = encodeGRPCCertificateIssue(certificateIssue)
	}
	return items
}

func encodeGRPCCertificateIssue(certificateIssue *api.CertificateIssue) *apiv1.CertificateIssue {
	return &apiv1.CertificateIssue{
		TypeMeta:   transport.EncodeGRPCTypeMeta(certificateIssue.TypeMeta),
		ObjectMeta: transport.EncodeGRPCObjectMeta(certificateIssue.ObjectMeta),
		// TODO
	}
}

func decodeGRPCCreateRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.CreateCertificateIssueRequest)
	return endpoint.CreateRequest{
		CertificateIssue: decodeGRPCCertificateIssue(req.GetCertificateIssue()),
	}, nil
}

func decodeGRPCCreateResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.CertificateIssue)
	return endpoint.CreateResponse{
		CertificateIssue: decodeGRPCCertificateIssue(reply),
	}, nil
}

func encodeGRPCCreateResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.CreateResponse)
	return encodeGRPCCertificateIssue(resp.CertificateIssue), nil
}

func encodeGRPCCreateRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.CreateRequest)

	return &apiv1.CreateCertificateIssueRequest{
		CertificateIssue: encodeGRPCCertificateIssue(req.CertificateIssue),
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
	return nil, nil
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
	return endpoint.GetResponse{
		CertificateIssue: decodeGRPCCertificateIssue(reply),
	}, nil
}

func encodeGRPCGetResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.GetResponse)
	return encodeGRPCCertificateIssue(resp.CertificateIssue), nil
}

func encodeGRPCGetRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.GetRequest)

	return &apiv1.GetCertificateIssueRequest{
		Name: req.Name,
	}, nil
}

func decodeGRPCListRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	// req := grpcReq.(*apiv1.ListCertificateIssuesRequest)
	return endpoint.ListRequest{}, nil
}

func decodeGRPCListResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.ListCertificateIssuesResponse)
	return endpoint.ListResponse{
		CertificateIssues: decodeGRPCCertificateIssues(reply.GetItems()),
	}, nil
}

func encodeGRPCListResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.ListResponse)
	return encodeGRPCCertificateIssues(resp.CertificateIssues), nil
}

func encodeGRPCListRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &apiv1.ListCertificateIssuesRequest{}, nil
}

func decodeGRPCUpdateRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.UpdateCertificateIssueRequest)
	return endpoint.UpdateRequest{
		CertificateIssue: decodeGRPCCertificateIssue(req.GetCertificateIssue()),
	}, nil
}

func decodeGRPCUpdateResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.CertificateIssue)
	return endpoint.UpdateResponse{
		CertificateIssue: decodeGRPCCertificateIssue(reply),
	}, nil
}

func encodeGRPCUpdateResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.UpdateResponse)
	return encodeGRPCCertificateIssue(resp.CertificateIssue), nil
}

func encodeGRPCUpdateRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.UpdateRequest)

	return &apiv1.UpdateCertificateIssueRequest{
		CertificateIssue: encodeGRPCCertificateIssue(req.CertificateIssue),
	}, nil
}
