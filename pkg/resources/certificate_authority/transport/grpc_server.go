package transport

import (
	"context"

	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	types "github.com/gogo/protobuf/types"

	apiv1 "powerssl.io/pkg/api/v1"
	"powerssl.io/pkg/resources/certificate_authority/endpoints"
)

type grpcServer struct {
	create grpctransport.Handler
	delete grpctransport.Handler
	get    grpctransport.Handler
	list   grpctransport.Handler
	update grpctransport.Handler
}

func NewGRPCServer(endpoints endpoints.Endpoints, logger log.Logger) apiv1.CertificateAuthorityServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}

	return &grpcServer{
		create: grpctransport.NewServer(
			endpoints.CreateEndpoint,
			decodeGRPCCreateRequest,
			encodeGRPCCreateResponse,
			options...,
		),
		delete: grpctransport.NewServer(
			endpoints.DeleteEndpoint,
			decodeGRPCDeleteRequest,
			encodeGRPCDeleteResponse,
			options...,
		),
		get: grpctransport.NewServer(
			endpoints.GetEndpoint,
			decodeGRPCGetRequest,
			encodeGRPCGetResponse,
			options...,
		),
		list: grpctransport.NewServer(
			endpoints.ListEndpoint,
			decodeGRPCListRequest,
			encodeGRPCListResponse,
			options...,
		),
		update: grpctransport.NewServer(
			endpoints.UpdateEndpoint,
			decodeGRPCUpdateRequest,
			encodeGRPCUpdateResponse,
			options...,
		),
	}
}

func (s *grpcServer) Create(ctx context.Context, req *apiv1.CreateCertificateAuthorityRequest) (*apiv1.CertificateAuthority, error) {
	_, rep, err := s.create.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.CertificateAuthority), nil
}

func (s *grpcServer) Delete(ctx context.Context, req *apiv1.DeleteCertificateAuthorityRequest) (*types.Empty, error) {
	_, rep, err := s.delete.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*types.Empty), nil
}

func (s *grpcServer) Get(ctx context.Context, req *apiv1.GetCertificateAuthorityRequest) (*apiv1.CertificateAuthority, error) {
	_, rep, err := s.get.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.CertificateAuthority), nil
}

func (s *grpcServer) List(ctx context.Context, req *apiv1.ListCertificateAuthoritiesRequest) (*apiv1.ListCertificateAuthoritiesResponse, error) {
	_, rep, err := s.list.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.ListCertificateAuthoritiesResponse), nil
}

func (s *grpcServer) Update(ctx context.Context, req *apiv1.UpdateCertificateAuthorityRequest) (*apiv1.CertificateAuthority, error) {
	_, rep, err := s.update.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.CertificateAuthority), nil
}
