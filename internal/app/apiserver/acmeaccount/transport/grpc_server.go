package transport

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/gogo/protobuf/types"
	stdopentracing "github.com/opentracing/opentracing-go"

	"powerssl.io/internal/app/apiserver/acmeaccount/endpoint"
	apiv1 "powerssl.io/internal/pkg/apiserver/api/v1"
)

type grpcServer struct {
	create grpctransport.Handler
	delete grpctransport.Handler
	get    grpctransport.Handler
	list   grpctransport.Handler
	update grpctransport.Handler
}

func NewGRPCServer(endpoints endpoint.Endpoints, logger log.Logger, tracer stdopentracing.Tracer) apiv1.ACMEAccountServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
		grpctransport.ServerBefore(jwt.GRPCToContext()),
	}

	return &grpcServer{
		create: grpctransport.NewServer(
			endpoints.CreateEndpoint,
			decodeGRPCCreateRequest,
			encodeGRPCCreateResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, fmt.Sprintf("/%s/Create", serviceName), logger)))...,
		),
		delete: grpctransport.NewServer(
			endpoints.DeleteEndpoint,
			decodeGRPCDeleteRequest,
			encodeGRPCDeleteResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, fmt.Sprintf("/%s/Delete", serviceName), logger)))...,
		),
		get: grpctransport.NewServer(
			endpoints.GetEndpoint,
			decodeGRPCGetRequest,
			encodeGRPCGetResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, fmt.Sprintf("/%s/Get", serviceName), logger)))...,
		),
		list: grpctransport.NewServer(
			endpoints.ListEndpoint,
			decodeGRPCListRequest,
			encodeGRPCListResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, fmt.Sprintf("/%s/List", serviceName), logger)))...,
		),
		update: grpctransport.NewServer(
			endpoints.UpdateEndpoint,
			decodeGRPCUpdateRequest,
			encodeGRPCUpdateResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, fmt.Sprintf("/%s/Update", serviceName), logger)))...,
		),
	}
}

func (s *grpcServer) Create(ctx context.Context, req *apiv1.CreateACMEAccountRequest) (*apiv1.ACMEAccount, error) {
	_, rep, err := s.create.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.ACMEAccount), nil
}

func (s *grpcServer) Delete(ctx context.Context, req *apiv1.DeleteACMEAccountRequest) (*types.Empty, error) {
	_, rep, err := s.delete.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*types.Empty), nil
}

func (s *grpcServer) Get(ctx context.Context, req *apiv1.GetACMEAccountRequest) (*apiv1.ACMEAccount, error) {
	_, rep, err := s.get.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.ACMEAccount), nil
}

func (s *grpcServer) List(ctx context.Context, req *apiv1.ListACMEAccountsRequest) (*apiv1.ListACMEAccountsResponse, error) {
	_, rep, err := s.list.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.ListACMEAccountsResponse), nil
}

func (s *grpcServer) Update(ctx context.Context, req *apiv1.UpdateACMEAccountRequest) (*apiv1.ACMEAccount, error) {
	_, rep, err := s.update.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.ACMEAccount), nil
}
