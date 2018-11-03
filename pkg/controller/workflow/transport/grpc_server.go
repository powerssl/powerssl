package transport // import "powerssl.io/pkg/controller/workflow/transport"

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	stdopentracing "github.com/opentracing/opentracing-go"

	apiv1 "powerssl.io/pkg/controller/api/v1"
	"powerssl.io/pkg/controller/workflow/endpoint"
)

type grpcServer struct {
	create grpctransport.Handler
}

func NewGRPCServer(endpoints endpoint.Endpoints, logger log.Logger, tracer stdopentracing.Tracer) apiv1.WorkflowServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}

	return &grpcServer{
		create: grpctransport.NewServer(
			endpoints.CreateEndpoint,
			decodeGRPCCreateRequest,
			encodeGRPCCreateResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, fmt.Sprintf("/%s/Create", serviceName), logger)))...,
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
