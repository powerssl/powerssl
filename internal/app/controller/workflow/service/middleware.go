package service

import (
	"context"

	"github.com/go-kit/kit/log"

	"powerssl.io/internal/app/controller/workflow/meta"
	"powerssl.io/pkg/controller/api"
)

type Middleware func(meta.Service) meta.Service

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next meta.Service) meta.Service {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   meta.Service
}

func (mw loggingMiddleware) Create(ctx context.Context, workflow *api.Workflow) (*api.Workflow, error) {
	defer func() {
		mw.logger.Log("method", "Create")
	}()
	return mw.next.Create(ctx, workflow)
}
