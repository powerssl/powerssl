package transport // import "powerssl.dev/sdk/apiserver/certificate/transport"

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	stdopentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/protobuf/types/known/emptypb"

	apiv1 "powerssl.dev/api/apiserver/v1"
	"powerssl.dev/common/log"

	"powerssl.dev/sdk/apiserver/certificate/endpoint"
)

type grpcServer struct {
	create grpctransport.Handler
	delete grpctransport.Handler
	get    grpctransport.Handler
	list   grpctransport.Handler
	update grpctransport.Handler
	apiv1.UnimplementedCertificateServiceServer
}

func NewGRPCServer(endpoints endpoint.Endpoints, logger log.Logger, tracer stdopentracing.Tracer) apiv1.CertificateServiceServer {
	kitLoger := log.KitLogger(logger)

	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(kitLoger),
		grpctransport.ServerBefore(jwt.GRPCToContext()),
	}

	return &grpcServer{
		create: grpctransport.NewServer(
			endpoints.CreateEndpoint,
			decodeGRPCCreateRequest,
			encodeGRPCCreateResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, fmt.Sprintf("/%s/Create", serviceName), kitLoger)))...,
		),
		delete: grpctransport.NewServer(
			endpoints.DeleteEndpoint,
			decodeGRPCDeleteRequest,
			encodeGRPCDeleteResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, fmt.Sprintf("/%s/Delete", serviceName), kitLoger)))...,
		),
		get: grpctransport.NewServer(
			endpoints.GetEndpoint,
			decodeGRPCGetRequest,
			encodeGRPCGetResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, fmt.Sprintf("/%s/Get", serviceName), kitLoger)))...,
		),
		list: grpctransport.NewServer(
			endpoints.ListEndpoint,
			decodeGRPCListRequest,
			encodeGRPCListResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, fmt.Sprintf("/%s/List", serviceName), kitLoger)))...,
		),
		update: grpctransport.NewServer(
			endpoints.UpdateEndpoint,
			decodeGRPCUpdateRequest,
			encodeGRPCUpdateResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, fmt.Sprintf("/%s/Update", serviceName), kitLoger)))...,
		),
	}
}

func (s *grpcServer) Create(ctx context.Context, req *apiv1.CreateCertificateRequest) (*apiv1.Certificate, error) {
	_, rep, err := s.create.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.Certificate), nil
}

func (s *grpcServer) Delete(ctx context.Context, req *apiv1.DeleteCertificateRequest) (*emptypb.Empty, error) {
	_, rep, err := s.delete.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func (s *grpcServer) Get(ctx context.Context, req *apiv1.GetCertificateRequest) (*apiv1.Certificate, error) {
	_, rep, err := s.get.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.Certificate), nil
}

func (s *grpcServer) List(ctx context.Context, req *apiv1.ListCertificatesRequest) (*apiv1.ListCertificatesResponse, error) {
	_, rep, err := s.list.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.ListCertificatesResponse), nil
}

func (s *grpcServer) Update(ctx context.Context, req *apiv1.UpdateCertificateRequest) (*apiv1.Certificate, error) {
	_, rep, err := s.update.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.Certificate), nil
}
