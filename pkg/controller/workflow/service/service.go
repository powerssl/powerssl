package service // import "powerssl.io/pkg/controller/workflow/service"

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"

	"powerssl.io/pkg/controller/api"
	workflowengine "powerssl.io/pkg/controller/workflow/engine"
	"powerssl.io/pkg/controller/workflow/engine/activity"
	"powerssl.io/pkg/controller/workflow/engine/workflow"
)

type Service interface {
	Create(ctx context.Context, kind string) (*api.Workflow, error)
}

func New(logger log.Logger, workflowengine *workflowengine.Engine) Service {
	var svc Service
	{
		svc = NewBasicService(logger, workflowengine)
		svc = LoggingMiddleware(logger)(svc)
	}
	return svc
}

type basicService struct {
	logger         log.Logger
	workflowengine *workflowengine.Engine
}

func NewBasicService(logger log.Logger, workflowengine *workflowengine.Engine) Service {
	return basicService{
		logger:         logger,
		workflowengine: workflowengine,
	}
}

func (bs basicService) Create(_ context.Context, kind string) (*api.Workflow, error) {
	// TODO: Decide which workflow
	workflow := workflow.New(kind)
	bs.workflowengine.AddWorkflow(workflow)

	activity := activity.New(api.Activity_ACME_CREATE_ACCOUNT)
	activity.GetRequest = func(activity *api.Activity) (*api.Activity, string, bool, []string, error) {
		return activity, "example.com", true, []string{"foo"}, nil
	}
	activity.SetResponse = func(account *api.Account, erro *api.Error) {
		fmt.Printf("Activity: %#v\n", activity)
		fmt.Printf("Account: %#v\n", account)
		fmt.Println("Status: ", account.Status)
	}
	workflow.AddActivity(activity)
	bs.workflowengine.AddActivity(activity)

	return workflow.API(), nil
}
