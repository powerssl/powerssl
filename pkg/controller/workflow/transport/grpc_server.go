package transport // import "powerssl.io/pkg/controller/workflow/transport"

import (
	"context"

	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"

	apiv1 "powerssl.io/pkg/controller/api/v1"
	"powerssl.io/pkg/controller/workflow/endpoint"
)

type grpcServer struct {
	create grpctransport.Handler
}

func NewGRPCServer(endpoints endpoint.Endpoints, logger log.Logger) apiv1.WorkflowServiceServer {
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
	}
}

func (s *grpcServer) Create(ctx context.Context, req *apiv1.CreateWorkflowRequest) (*apiv1.Workflow, error) {
	_, rep, err := s.create.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.Workflow), nil
}
