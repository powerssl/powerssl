package transport // import "powerssl.io/pkg/controller/workflow/transport"

import (
	"context"

	"powerssl.io/pkg/controller/api"
	apiv1 "powerssl.io/pkg/controller/api/v1"
	"powerssl.io/pkg/controller/workflow/endpoint"
)

func decodeGRPCCreateRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.CreateWorkflowRequest)
	return endpoint.CreateRequest{
		Kind: req.Kind,
	}, nil
}

func decodeGRPCCreateResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.Workflow)
	return endpoint.CreateResponse{
		Workflow: &api.Workflow{
			Name: reply.GetName(),
			Kind: reply.GetKind(),
		},
	}, nil
}

func encodeGRPCCreateResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.CreateResponse)
	return &apiv1.Workflow{
		Name: resp.Workflow.Name,
		Kind: resp.Workflow.Kind,
	}, nil
}

func encodeGRPCCreateRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.CreateRequest)
	return &apiv1.CreateWorkflowRequest{
		Kind: req.Kind,
	}, nil
}
