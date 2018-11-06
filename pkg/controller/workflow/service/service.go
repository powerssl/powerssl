package service // import "powerssl.io/pkg/controller/workflow/service"

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"

	"powerssl.io/pkg/controller/api"
	"powerssl.io/pkg/controller/workflow/engine/workflow"
	"powerssl.io/pkg/util/tracing"
)

type Service interface {
	Create(ctx context.Context, workflow *api.Workflow) (*api.Workflow, error)
}

func New(logger log.Logger) Service {
	var svc Service
	{
		svc = NewBasicService(logger)
		svc = LoggingMiddleware(logger)(svc)
	}
	return svc
}

type basicService struct {
	logger log.Logger
}

func NewBasicService(logger log.Logger) Service {
	return basicService{
		logger: logger,
	}
}

func (bs basicService) Create(ctx context.Context, apiWorkflow *api.Workflow) (*api.Workflow, error) {
	workflow, err := newWorkflowFromAPI(apiWorkflow)
	if err != nil {
		return nil, err
	}

	// Create new context that won't be canceled when the RPC returns.
	// Inherit the span from previous context though.
	workflow.Execute(tracing.ContextWithSpanFromContext(context.Background(), ctx))

	return workflow.ToAPI(), nil
}

func newWorkflowFromAPI(apiWorkflow *api.Workflow) (*workflow.Workflow, error) {
	var definition workflow.Definition
	switch apiWorkflow.Kind {
	case api.WorkflowKindCreateACMEAccount:
		input, ok := apiWorkflow.Input.(*api.CreateACMEAccountInput)
		if !ok {
			return nil, fmt.Errorf("wrong input for workflow")
		}
		definition = workflow.CreateAccount{
			DirectoryURL:         input.DirectoryURL,
			TermsOfServiceAgreed: input.TermsOfServiceAgreed,
			Contacts:             input.Contacts,
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
