package transport

import (
	"context"

	"powerssl.dev/powerssl/internal/app/controller/workflow/endpoint"
	apiv1 "powerssl.dev/powerssl/internal/pkg/controller/api/v1"
	"powerssl.dev/powerssl/pkg/controller/api"
)

func decodeGRPCWorkflowIntegrationFilter(workflowIntegrationFilter *apiv1.Workflow_IntegrationFilter) (*api.WorkflowIntegrationFilter, error) {
	return &api.WorkflowIntegrationFilter{
		Kind: api.IntegrationKind(workflowIntegrationFilter.GetKind()),
		Name: workflowIntegrationFilter.GetName(),
	}, nil
}

func encodeGRPCWorkflowIntegrationFilter(workflowIntegrationFilter *api.WorkflowIntegrationFilter) (*apiv1.Workflow_IntegrationFilter, error) {
	return &apiv1.Workflow_IntegrationFilter{
		Kind: apiv1.IntegrationKind(workflowIntegrationFilter.Kind),
		Name: workflowIntegrationFilter.Name,
	}, nil
}

func decodeGRPCWorkflowIntegrationFilters(grpcWorkflowIntegrationFilters []*apiv1.Workflow_IntegrationFilter) ([]*api.WorkflowIntegrationFilter, error) {
	workflowIntegrationFilters := make([]*api.WorkflowIntegrationFilter, len(grpcWorkflowIntegrationFilters))
	for i, grpcWorkflowIntegrationFilter := range grpcWorkflowIntegrationFilters {
		var err error
		workflowIntegrationFilters[i], err = decodeGRPCWorkflowIntegrationFilter(grpcWorkflowIntegrationFilter)
		if err != nil {
			return nil, err
		}
	}
	return workflowIntegrationFilters, nil
}

func encodeGRPCWorkflowIntegrationFilters(workflowIntegrationFilters []*api.WorkflowIntegrationFilter) ([]*apiv1.Workflow_IntegrationFilter, error) {
	grpcWorkflowIntegrationFilters := make([]*apiv1.Workflow_IntegrationFilter, len(workflowIntegrationFilters))
	for i, workflowIntegrationFilter := range workflowIntegrationFilters {
		var err error
		grpcWorkflowIntegrationFilters[i], err = encodeGRPCWorkflowIntegrationFilter(workflowIntegrationFilter)
		if err != nil {
			return nil, err
		}
	}
	return grpcWorkflowIntegrationFilters, nil
}

func decodeGRPCWorkflow(grpcWorkflow *apiv1.Workflow) (*api.Workflow, error) {
	integrationFilters, err := decodeGRPCWorkflowIntegrationFilters(grpcWorkflow.GetIntegrationFilters())
	if err != nil {
		return nil, err
	}
	workflow := &api.Workflow{
		Name:               grpcWorkflow.GetName(),
		Kind:               api.WorkflowKind(grpcWorkflow.GetKind()),
		IntegrationFilters: integrationFilters,
	}
	switch grpcWorkflow.GetInput().(type) {
	case *apiv1.Workflow_CreateAcmeAccountInput:
		input := grpcWorkflow.GetCreateAcmeAccountInput()
		workflow.Input = &api.CreateACMEAccountInput{
			Account:              input.GetAccount(),
			DirectoryURL:         input.GetDirectoryUrl(),
			TermsOfServiceAgreed: input.GetTermsOfServiceAgreed(),
			Contacts:             input.GetContacts(),
		}
	case *apiv1.Workflow_RequestAcmeCertificateInput:
		input := grpcWorkflow.GetRequestAcmeCertificateInput()
		workflow.Input = &api.RequestACMECertificateInput{
			DirectoryURL: input.GetDirectoryUrl(),
			AccountURL:   input.GetAccountUrl(),
			Dnsnames:     input.GetDnsnames(),
			NotBefore:    input.GetNotBefore(),
			NotAfter:     input.GetNotAfter(),
		}
	}
	return workflow, nil
}

func encodeGRPCWorkflow(workflow *api.Workflow) (*apiv1.Workflow, error) {
	if workflow == nil {
		workflow = &api.Workflow{}
	}
	integrationFilters, err := encodeGRPCWorkflowIntegrationFilters(workflow.IntegrationFilters)
	if err != nil {
		return nil, err
	}
	grpcWorkflow := &apiv1.Workflow{
		Name:               workflow.Name,
		Kind:               apiv1.WorkflowKind(workflow.Kind),
		IntegrationFilters: integrationFilters,
	}
	switch workflow.Input.(type) {
	case *api.CreateACMEAccountInput:
		input := workflow.Input.(*api.CreateACMEAccountInput)
		grpcWorkflow.Input = &apiv1.Workflow_CreateAcmeAccountInput{
			CreateAcmeAccountInput: &apiv1.CreateACMEAccountInput{
				Account:              input.Account,
				DirectoryUrl:         input.DirectoryURL,
				TermsOfServiceAgreed: input.TermsOfServiceAgreed,
				Contacts:             input.Contacts,
			},
		}
	case *api.RequestACMECertificateInput:
		input := workflow.Input.(*api.RequestACMECertificateInput)
		grpcWorkflow.Input = &apiv1.Workflow_RequestAcmeCertificateInput{
			RequestAcmeCertificateInput: &apiv1.RequestACMECertificateInput{
				DirectoryUrl: input.DirectoryURL,
				AccountUrl:   input.AccountURL,
				Dnsnames:     input.Dnsnames,
				NotBefore:    input.NotBefore,
				NotAfter:     input.NotAfter,
			},
		}
	}
	return grpcWorkflow, nil
}

func decodeGRPCCreateRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.CreateWorkflowRequest)
	workflow, err := decodeGRPCWorkflow(req.GetWorkflow())
	if err != nil {
		return nil, err
	}
	return endpoint.CreateRequest{
		Workflow: workflow,
	}, nil
}

func decodeGRPCCreateResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.Workflow)
	workflow, err := decodeGRPCWorkflow(reply)
	if err != nil {
		return nil, err
	}
	return endpoint.CreateResponse{
		Workflow: workflow,
	}, nil
}

func encodeGRPCCreateResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.CreateResponse)
	return encodeGRPCWorkflow(resp.Workflow)
}

func encodeGRPCCreateRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.CreateRequest)
	workflow, err := encodeGRPCWorkflow(req.Workflow)
	if err != nil {
		return nil, err
	}
	return &apiv1.CreateWorkflowRequest{
		Workflow: workflow,
	}, nil
}
