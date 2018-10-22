package service // import "powerssl.io/pkg/controller/workflow/service"

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/google/uuid"

	"powerssl.io/pkg/controller/api"
	workflowengine "powerssl.io/pkg/controller/workflow/engine"
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
	bs.workflowengine.Create(kind)
	return &api.Workflow{Name: fmt.Sprintf("workflows/%s", uuid.New()), Kind: kind}, nil
}
