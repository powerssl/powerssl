package service // import "powerssl.io/pkg/controller/workflow/service"

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"

	"powerssl.io/pkg/controller/api"
	"powerssl.io/pkg/controller/integration"
	"powerssl.io/pkg/controller/workflow/engine/activity"
	"powerssl.io/pkg/controller/workflow/engine/workflow"
)

type Service interface {
	Create(ctx context.Context, kind string) (*api.Workflow, error)
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

func (bs basicService) Create(_ context.Context, kind string) (*api.Workflow, error) {
	// TODO: Decide which workflow
	workflow := workflow.New(kind)

	a := activity.New(api.Activity_ACME_CREATE_ACCOUNT)
	a.GetRequest = func(activity *api.Activity) (*api.Activity, string, bool, []string, error) {
		return activity, "example.com", true, []string{"foo"}, nil
	}
	a.SetResponse = func(account *api.Account, erro *api.Error) error {
		fmt.Printf("Activity: %#v\n", a)
		fmt.Printf("Account: %#v\n", account)
		fmt.Println("Status: ", account.Status)
		return nil
	}
	workflow.AddActivity(a)
	integ, err := integration.Integrations.GetByKind(a.IntegrationKind())
	if err != nil {
		panic(err) // TODO
	}
	a.Execute(integ)
	fmt.Println(activity.Activities)
	b, err := activity.Activities.Get(a.UUID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", b)

	return workflow.API(), nil
}
