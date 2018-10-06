package transport

import (
	"context"

	"powerssl.io/pkg/api"
	apiv1 "powerssl.io/pkg/api/v1"
	"powerssl.io/pkg/resources/certificate_authority/endpoints"
	"powerssl.io/pkg/resources/transport"
)

func decodeGRPCCertificateAuthorities(certificateAuthorities []*apiv1.CertificateAuthority) []*api.CertificateAuthority {
	items := make([]*api.CertificateAuthority, len(certificateAuthorities))
	for i, certificateAuthority := range certificateAuthorities {
		items[i] = decodeGRPCCertificateAuthority(certificateAuthority)
	}
	return items
}

func decodeGRPCCertificateAuthority(certificateAuthority *apiv1.CertificateAuthority) *api.CertificateAuthority {
	return &api.CertificateAuthority{
		TypeMeta:   transport.DecodeGRPCTypeMeta(certificateAuthority.GetTypeMeta()),
		ObjectMeta: transport.DecodeGRPCObjectMeta(certificateAuthority.GetObjectMeta()),
		Spec: api.CertificateAuthoritySpec{
			Vendor: certificateAuthority.GetSpec().GetVendor(),
		},
	}
}

func encodeGRPCCertificateAuthorities(certificateAuthorities []*api.CertificateAuthority) []*apiv1.CertificateAuthority {
	items := make([]*apiv1.CertificateAuthority, len(certificateAuthorities))
	for i, certificateAuthority := range certificateAuthorities {
		items[i] = encodeGRPCCertificateAuthority(certificateAuthority)
	}
	return items
}

func encodeGRPCCertificateAuthority(certificateAuthority *api.CertificateAuthority) *apiv1.CertificateAuthority {
	return &apiv1.CertificateAuthority{
		TypeMeta:   transport.EncodeGRPCTypeMeta(certificateAuthority.TypeMeta),
		ObjectMeta: transport.EncodeGRPCObjectMeta(certificateAuthority.ObjectMeta),
		Spec: &apiv1.CertificateAuthoritySpec{
			Vendor: certificateAuthority.Spec.Vendor,
		},
	}
}

func decodeGRPCCreateRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.CreateCertificateAuthorityRequest)
	return endpoints.CreateRequest{
		CertificateAuthority: decodeGRPCCertificateAuthority(req.GetCertificateAuthority()),
	}, nil
}

func decodeGRPCCreateResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.CertificateAuthority)
	return endpoints.CreateResponse{
		CertificateAuthority: decodeGRPCCertificateAuthority(reply),
	}, nil
}

func decodeGRPCDeleteRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.DeleteCertificateAuthorityRequest)
	return endpoints.DeleteRequest{
		Name: req.GetName(),
	}, nil
}

func decodeGRPCDeleteResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	return endpoints.DeleteResponse{}, nil
}

func decodeGRPCGetRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.GetCertificateAuthorityRequest)
	return endpoints.GetRequest{
		Name: req.GetName(),
	}, nil
}

func decodeGRPCGetResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.CertificateAuthority)
	return endpoints.GetResponse{
		CertificateAuthority: decodeGRPCCertificateAuthority(reply),
	}, nil
}

func decodeGRPCListRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return endpoints.ListRequest{}, nil
}

func decodeGRPCListResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.ListCertificateAuthoritiesResponse)
	return endpoints.ListResponse{
		CertificateAuthorities: decodeGRPCCertificateAuthorities(reply.GetItems()),
	}, nil
}

func decodeGRPCUpdateRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.UpdateCertificateAuthorityRequest)
	return endpoints.UpdateRequest{
		CertificateAuthority: decodeGRPCCertificateAuthority(req.GetCertificateAuthority()),
	}, nil
}

func decodeGRPCUpdateResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.CertificateAuthority)
	return endpoints.UpdateResponse{
		CertificateAuthority: decodeGRPCCertificateAuthority(reply),
	}, nil
}

func encodeGRPCCreateResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.CreateResponse)
	return encodeGRPCCertificateAuthority(resp.CertificateAuthority), nil
}

func encodeGRPCCreateRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoints.CreateRequest)
	return &apiv1.CreateCertificateAuthorityRequest{
		CertificateAuthority: encodeGRPCCertificateAuthority(req.CertificateAuthority),
	}, nil
}

func encodeGRPCDeleteResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func encodeGRPCDeleteRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoints.DeleteRequest)
	return &apiv1.DeleteCertificateAuthorityRequest{
		Name: req.Name,
	}, nil
}

func encodeGRPCGetResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.GetResponse)
	return encodeGRPCCertificateAuthority(resp.CertificateAuthority), nil
}

func encodeGRPCGetRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoints.GetRequest)
	return &apiv1.GetCertificateAuthorityRequest{
		Name: req.Name,
	}, nil
}

func encodeGRPCListResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.ListResponse)
	return encodeGRPCCertificateAuthorities(resp.CertificateAuthorities), nil
}

func encodeGRPCListRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &apiv1.ListCertificateAuthoritiesRequest{}, nil
}

func encodeGRPCUpdateResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.UpdateResponse)
	return encodeGRPCCertificateAuthority(resp.CertificateAuthority), nil
}

func encodeGRPCUpdateRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoints.UpdateRequest)
	return &apiv1.UpdateCertificateAuthorityRequest{
		CertificateAuthority: encodeGRPCCertificateAuthority(req.CertificateAuthority),
	}, nil
}
