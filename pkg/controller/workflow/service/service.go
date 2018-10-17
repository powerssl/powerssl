package service // import "powerssl.io/pkg/controller/workflow/service"

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/google/uuid"

	"powerssl.io/pkg/controller/api"
)

type Service interface {
	Create(ctx context.Context, kind string) (*api.Workflow, error)
}

func New(logger log.Logger) Service {
	var svc Service
	{
		svc = NewBasicService()
		svc = LoggingMiddleware(logger)(svc)
	}
	return svc
}

type basicService struct {
}

func NewBasicService() Service {
	return basicService{}
}

func (bs basicService) Create(_ context.Context, kind string) (*api.Workflow, error) {
	return &api.Workflow{Name: fmt.Sprintf("workflows/%s", uuid.New()), Kind: kind}, nil
}
