package service // import "powerssl.io/pkg/controller/workflow/service"

import (
	"context"

	"github.com/go-kit/kit/log"

	"powerssl.io/pkg/controller/api"
)

type Middleware func(Service) Service

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next Service) Service {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   Service
}

func (mw loggingMiddleware) Create(ctx context.Context, workflow *api.Workflow) (*api.Workflow, error) {
	defer func() {
		mw.logger.Log("method", "Create")
	}()
	return mw.next.Create(ctx, workflow)
}
