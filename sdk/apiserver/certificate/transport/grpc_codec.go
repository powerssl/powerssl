package transport // import "powerssl.dev/sdk/apiserver/certificate/transport"

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	apiv1 "powerssl.dev/api/apiserver/v1"

	"powerssl.dev/sdk/apiserver/api"
	"powerssl.dev/sdk/apiserver/certificate/endpoint"
)

func decodeGRPCCertificate(certificate *apiv1.Certificate) (*api.Certificate, error) {
	return &api.Certificate{
		Name:            certificate.GetName(),
		CreateTime:      certificate.GetCreateTime().AsTime(),
		UpdateTime:      certificate.GetUpdateTime().AsTime(),
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
	return &apiv1.Certificate{
		Name:            certificate.Name,
		CreateTime:      timestamppb.New(certificate.CreateTime),
		UpdateTime:      timestamppb.New(certificate.UpdateTime),
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

func decodeGRPCDeleteResponse(_ context.Context, _ interface{}) (interface{}, error) {
	return endpoint.DeleteResponse{}, nil
}

func encodeGRPCDeleteRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.DeleteRequest)
	return &apiv1.DeleteCertificateRequest{
		Name: req.Name,
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

func encodeGRPCGetRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.GetRequest)
	return &apiv1.GetCertificateRequest{
		Name: req.Name,
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

func encodeGRPCListRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.ListRequest)
	return &apiv1.ListCertificatesRequest{
		PageSize:  int32(req.PageSize),
		PageToken: req.PageToken,
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
