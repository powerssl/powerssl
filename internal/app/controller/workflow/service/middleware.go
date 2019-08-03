package service

import (
	"context"

	"github.com/go-kit/kit/log"

	"powerssl.dev/powerssl/pkg/controller/api"
	"powerssl.dev/powerssl/pkg/controller/workflow"
)

type Middleware func(workflow.Service) workflow.Service

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next workflow.Service) workflow.Service {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   workflow.Service
}

func (mw loggingMiddleware) Create(ctx context.Context, workflow *api.Workflow) (*api.Workflow, error) {
	defer func() {
		mw.logger.Log("method", "Create")
	}()
	return mw.next.Create(ctx, workflow)
}
