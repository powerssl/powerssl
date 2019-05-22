package transport

import (
	"context"

	"github.com/gogo/protobuf/types"

	"powerssl.io/powerssl/internal/app/apiserver/certificate/endpoint"
	apiv1 "powerssl.io/powerssl/internal/pkg/apiserver/api/v1"
	"powerssl.io/powerssl/pkg/apiserver/api"
)

func decodeGRPCCertificate(certificate *apiv1.Certificate) (*api.Certificate, error) {
	createTime, err := types.TimestampFromProto(certificate.GetCreateTime())
	if err != nil {
		return nil, err
	}
	updateTime, err := types.TimestampFromProto(certificate.GetUpdateTime())
	if err != nil {
		return nil, err
	}
	return &api.Certificate{
		Name:            certificate.GetName(),
		CreateTime:      createTime,
		UpdateTime:      updateTime,
		DisplayName:     certificate.GetDisplayName(),
		Title:           certificate.GetTitle(),
		Description:     certificate.GetDescription(),
		Labels:          certificate.GetLabels(),
		Dnsnames:        certificate.GetDnsnames(),
		KeyAlgorithm:    certificate.GetKeyAlgorithm().String(),
		KeySize:         certificate.GetKeySize(),
		DigestAlgorithm: certificate.GetDigestAlgorithm().String(),
		AutoRenew:       certificate.GetAutoRenew(),
	}, nil
}

func encodeGRPCCertificate(certificate *api.Certificate) (*apiv1.Certificate, error) {
	createTime, err := types.TimestampProto(certificate.CreateTime)
	if err != nil {
		return nil, err
	}
	updateTime, err := types.TimestampProto(certificate.UpdateTime)
	if err != nil {
		return nil, err
	}
	return &apiv1.Certificate{
		Name:            certificate.Name,
		CreateTime:      createTime,
		UpdateTime:      updateTime,
		DisplayName:     certificate.DisplayName,
		Title:           certificate.Title,
		Description:     certificate.Description,
		Labels:          certificate.Labels,
		Dnsnames:        certificate.Dnsnames,
		KeyAlgorithm:    apiv1.KeyAlgorithm(apiv1.KeyAlgorithm_value[certificate.KeyAlgorithm]),
		KeySize:         certificate.KeySize,
		DigestAlgorithm: apiv1.DigestAlgorithm(apiv1.DigestAlgorithm_value[certificate.DigestAlgorithm]),
		AutoRenew:       certificate.AutoRenew,
	}, nil
}

func decodeGRPCCertificates(grpcCertificates []*apiv1.Certificate) ([]*api.Certificate, error) {
	certificates := make([]*api.Certificate, len(grpcCertificates))
	for i, grpcCertificate := range grpcCertificates {
		certificate, err := decodeGRPCCertificate(grpcCertificate)
		if err != nil {
			return nil, err
		}
		certificates[i] = certificate
	}
	return certificates, nil
}

func encodeGRPCCertificates(certificates []*api.Certificate) ([]*apiv1.Certificate, error) {
	grpcCertificates := make([]*apiv1.Certificate, len(certificates))
	for i, certificate := range certificates {
		grpcCertificate, err := encodeGRPCCertificate(certificate)
		if err != nil {
			return nil, err
		}
		grpcCertificates[i] = grpcCertificate
	}
	return grpcCertificates, nil
}

func decodeGRPCCreateRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.CreateCertificateRequest)
	certificate, err := decodeGRPCCertificate(req.GetCertificate())
	if err != nil {
		return nil, err
	}
	return endpoint.CreateRequest{
		Certificate: certificate,
	}, nil
}

func decodeGRPCCreateResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.Certificate)
	certificate, err := decodeGRPCCertificate(reply)
	if err != nil {
		return nil, err
	}
	return endpoint.CreateResponse{
		Certificate: certificate,
	}, nil
}

func encodeGRPCCreateResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.CreateResponse)
	return encodeGRPCCertificate(resp.Certificate)
}

func encodeGRPCCreateRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.CreateRequest)
	certificate, err := encodeGRPCCertificate(req.Certificate)
	if err != nil {
		return nil, err
	}
	return &apiv1.CreateCertificateRequest{
		Certificate: certificate,
	}, nil
}

func decodeGRPCDeleteRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.DeleteCertificateRequest)
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
	return &apiv1.DeleteCertificateRequest{
		Name: req.Name,
	}, nil
}

func decodeGRPCGetRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.GetCertificateRequest)
	return endpoint.GetRequest{
		Name: req.GetName(),
	}, nil
}

func decodeGRPCGetResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.Certificate)
	certificate, err := decodeGRPCCertificate(reply)
	if err != nil {
		return nil, err
	}
	return endpoint.GetResponse{
		Certificate: certificate,
	}, nil
}

func encodeGRPCGetResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.GetResponse)
	return encodeGRPCCertificate(resp.Certificate)
}

func encodeGRPCGetRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.GetRequest)
	return &apiv1.GetCertificateRequest{
		Name: req.Name,
	}, nil
}

func decodeGRPCListRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.ListCertificatesRequest)
	return endpoint.ListRequest{
		PageSize:  int(req.GetPageSize()),
		PageToken: req.GetPageToken(),
	}, nil
}

func decodeGRPCListResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.ListCertificatesResponse)
	certificates, err := decodeGRPCCertificates(reply.GetCertificates())
	if err != nil {
		return nil, err
	}
	return endpoint.ListResponse{
		Certificates:  certificates,
		NextPageToken: reply.GetNextPageToken(),
	}, nil
}

func encodeGRPCListResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.ListResponse)
	certificates, err := encodeGRPCCertificates(resp.Certificates)
	if err != nil {
		return nil, err
	}
	return &apiv1.ListCertificatesResponse{
		Certificates:  certificates,
		NextPageToken: resp.NextPageToken,
	}, nil
}

func encodeGRPCListRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.ListRequest)
	return &apiv1.ListCertificatesRequest{
		PageSize:  int32(req.PageSize),
		PageToken: req.PageToken,
	}, nil
}

func decodeGRPCUpdateRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.UpdateCertificateRequest)
	certificate, err := decodeGRPCCertificate(req.GetCertificate())
	if err != nil {
		return nil, err
	}
	return endpoint.UpdateRequest{
		Name:        req.GetName(),
		Certificate: certificate,
	}, nil
}

func decodeGRPCUpdateResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.Certificate)
	certificate, err := decodeGRPCCertificate(reply)
	if err != nil {
		return nil, err
	}
	return endpoint.UpdateResponse{
		Certificate: certificate,
	}, nil
}

func encodeGRPCUpdateResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.UpdateResponse)
	return encodeGRPCCertificate(resp.Certificate)
}

func encodeGRPCUpdateRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.UpdateRequest)
	certificate, err := encodeGRPCCertificate(req.Certificate)
	if err != nil {
		return nil, err
	}
	return &apiv1.UpdateCertificateRequest{
		Name:        req.Name,
		Certificate: certificate,
	}, nil
}
