package transport // import "powerssl.io/pkg/controller/ca/transport"

import (
	"context"

	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/gogo/protobuf/types"

	apiv1 "powerssl.io/pkg/controller/api/v1"
	"powerssl.io/pkg/controller/ca/endpoint"
)

type grpcServer struct {
	getAuthorizeDomainRequest  grpctransport.Handler
	setAuthorizeDomainResponse grpctransport.Handler

	getRequestCertificateRequest  grpctransport.Handler
	setRequestCertificateResponse grpctransport.Handler

	getRevokeCertificateRequest  grpctransport.Handler
	setRevokeCertificateResponse grpctransport.Handler

	getVerifyDomainRequest  grpctransport.Handler
	setVerifyDomainResponse grpctransport.Handler
}

func NewGRPCServer(endpoints endpoint.Endpoints, logger log.Logger) apiv1.CAServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}

	return &grpcServer{
		getAuthorizeDomainRequest: grpctransport.NewServer(
			endpoints.GetAuthorizeDomainRequestEndpoint,
			decodeGRPCGetAuthorizeDomainRequestRequest,
			encodeGRPCGetAuthorizeDomainRequestResponse,
			options...,
		),
		setAuthorizeDomainResponse: grpctransport.NewServer(
			endpoints.SetAuthorizeDomainResponseEndpoint,
			decodeGRPCSetAuthorizeDomainResponseRequest,
			encodeGRPCSetAuthorizeDomainResponseResponse,
			options...,
		),
		getRequestCertificateRequest: grpctransport.NewServer(
			endpoints.GetRequestCertificateRequestEndpoint,
			decodeGRPCGetRequestCertificateRequestRequest,
			encodeGRPCGetRequestCertificateRequestResponse,
			options...,
		),
		setRequestCertificateResponse: grpctransport.NewServer(
			endpoints.SetRequestCertificateResponseEndpoint,
			decodeGRPCSetRequestCertificateResponseRequest,
			encodeGRPCSetRequestCertificateResponseResponse,
			options...,
		),
		getRevokeCertificateRequest: grpctransport.NewServer(
			endpoints.GetRevokeCertificateRequestEndpoint,
			decodeGRPCGetRevokeCertificateRequestRequest,
			encodeGRPCGetRevokeCertificateRequestResponse,
			options...,
		),
		setRevokeCertificateResponse: grpctransport.NewServer(
			endpoints.SetRevokeCertificateResponseEndpoint,
			decodeGRPCSetRevokeCertificateResponseRequest,
			encodeGRPCSetRevokeCertificateResponseResponse,
			options...,
		),
		getVerifyDomainRequest: grpctransport.NewServer(
			endpoints.GetVerifyDomainRequestEndpoint,
			decodeGRPCGetVerifyDomainRequestRequest,
			encodeGRPCGetVerifyDomainRequestResponse,
			options...,
		),
		setVerifyDomainResponse: grpctransport.NewServer(
			endpoints.SetVerifyDomainResponseEndpoint,
			decodeGRPCSetVerifyDomainResponseRequest,
			encodeGRPCSetVerifyDomainResponseResponse,
			options...,
		),
	}
}

func (s *grpcServer) GetAuthorizeDomainRequest(ctx context.Context, req *apiv1.Activity) (*apiv1.GetAuthorizeDomainRequestResponse, error) {
	_, rep, err := s.getAuthorizeDomainRequest.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.GetAuthorizeDomainRequestResponse), nil
}

func (s *grpcServer) SetAuthorizeDomainResponse(ctx context.Context, req *apiv1.SetAuthorizeDomainResponseRequest) (*types.Empty, error) {
	_, rep, err := s.setAuthorizeDomainResponse.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*types.Empty), nil
}

func (s *grpcServer) GetRequestCertificateRequest(ctx context.Context, req *apiv1.Activity) (*apiv1.GetRequestCertificateRequestResponse, error) {
	_, rep, err := s.getRequestCertificateRequest.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.GetRequestCertificateRequestResponse), nil
}

func (s *grpcServer) SetRequestCertificateResponse(ctx context.Context, req *apiv1.SetRequestCertificateResponseRequest) (*types.Empty, error) {
	_, rep, err := s.setRequestCertificateResponse.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*types.Empty), nil
}

func (s *grpcServer) GetRevokeCertificateRequest(ctx context.Context, req *apiv1.Activity) (*apiv1.GetRevokeCertificateRequestResponse, error) {
	_, rep, err := s.getRevokeCertificateRequest.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.GetRevokeCertificateRequestResponse), nil
}

func (s *grpcServer) SetRevokeCertificateResponse(ctx context.Context, req *apiv1.SetRevokeCertificateResponseRequest) (*types.Empty, error) {
	_, rep, err := s.setRevokeCertificateResponse.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*types.Empty), nil
}

func (s *grpcServer) GetVerifyDomainRequest(ctx context.Context, req *apiv1.Activity) (*apiv1.GetVerifyDomainRequestResponse, error) {
	_, rep, err := s.getVerifyDomainRequest.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.GetVerifyDomainRequestResponse), nil
}

func (s *grpcServer) SetVerifyDomainResponse(ctx context.Context, req *apiv1.SetVerifyDomainResponseRequest) (*types.Empty, error) {
	_, rep, err := s.setVerifyDomainResponse.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*types.Empty), nil
}
