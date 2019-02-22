package service

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"

	"powerssl.io/internal/app/controller/workflow/engine/workflow"
	"powerssl.io/internal/pkg/util/tracing"
	apiserverclient "powerssl.io/pkg/apiserver/client"
	"powerssl.io/pkg/controller/api"
	service "powerssl.io/pkg/controller/workflow"
)

func New(logger log.Logger, client *apiserverclient.GRPCClient) service.Service {
	var svc service.Service
	{
		svc = NewBasicService(logger, client)
		svc = LoggingMiddleware(logger)(svc)
	}
	return svc
}

type basicService struct {
	client *apiserverclient.GRPCClient
	logger log.Logger
}

func NewBasicService(logger log.Logger, client *apiserverclient.GRPCClient) service.Service {
	return basicService{
		client: client,
		logger: logger,
	}
}

func (bs basicService) Create(ctx context.Context, apiWorkflow *api.Workflow) (*api.Workflow, error) {
	workflow, err := bs.newWorkflowFromAPI(apiWorkflow)
	if err != nil {
		return nil, err
	}

	// Create new context that won't be canceled when the RPC returns.
	// Inherit the span from previous context though.
	workflow.Execute(tracing.ContextWithSpanFromContext(context.Background(), ctx))

	return workflow.ToAPI(), nil
}

func (bs basicService) newWorkflowFromAPI(apiWorkflow *api.Workflow) (*workflow.Workflow, error) {
	var definition workflow.Definition
	switch apiWorkflow.Kind {
	case api.WorkflowKindCreateACMEAccount:
		input, ok := apiWorkflow.Input.(*api.CreateACMEAccountInput)
		if !ok {
			return nil, fmt.Errorf("wrong input for workflow")
		}
		definition = workflow.CreateAccount{
			Account:              input.Account,
			DirectoryURL:         input.DirectoryURL,
			TermsOfServiceAgreed: input.TermsOfServiceAgreed,
			Contacts:             input.Contacts,

			Client: bs.client,
		}
	case api.WorkflowKindRequestACMECertificate:
		input, ok := apiWorkflow.Input.(*api.RequestACMECertificateInput)
		if !ok {
			return nil, fmt.Errorf("wrong input for workflow")
		}
		definition = workflow.RequestCertificate{
			DirectoryURL: input.DirectoryURL,
			AccountURL:   input.AccountURL,
			Dnsnames:     input.Dnsnames,
			NotBefore:    input.NotBefore,
			NotAfter:     input.NotAfter,
		}
	default:
		return nil, fmt.Errorf("workflow kind not found")
	}
	return workflow.New(definition), nil
}
