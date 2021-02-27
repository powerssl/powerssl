package service

import (
	"context"

	"github.com/go-kit/kit/log"

	"powerssl.dev/sdk/apiserver/api"
	"powerssl.dev/sdk/apiserver/user"
)

type Middleware func(user.Service) user.Service

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next user.Service) user.Service {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   user.Service
}

func (mw loggingMiddleware) Create(ctx context.Context, user *api.User) (*api.User, error) {
	defer func() {
		mw.logger.Log("method", "Create", "user", true)
	}()
	return mw.next.Create(ctx, user)
}

func (mw loggingMiddleware) Delete(ctx context.Context, name string) error {
	defer func() {
		mw.logger.Log("method", "Delete", "name", name)
	}()
	return mw.next.Delete(ctx, name)
}

func (mw loggingMiddleware) Get(ctx context.Context, name string) (*api.User, error) {
	defer func() {
		mw.logger.Log("method", "Get", "name", name)
	}()
	return mw.next.Get(ctx, name)
}

func (mw loggingMiddleware) List(ctx context.Context, pageSize int, pageToken string) ([]*api.User, string, error) {
	defer func() {
		mw.logger.Log("method", "List", "pageSize", pageSize, "pageToken", pageToken)
	}()
	return mw.next.List(ctx, pageSize, pageToken)
}

func (mw loggingMiddleware) Update(ctx context.Context, name string, user *api.User) (*api.User, error) {
	defer func() {
		mw.logger.Log("method", "Update", "name", name, "user", true)
	}()
	return mw.next.Update(ctx, name, user)
}