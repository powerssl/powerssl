package service

import (
	"context"

	"powerssl.dev/common/log"
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

func (mw loggingMiddleware) Create(ctx context.Context, user *api.User) (_ *api.User, err error) {
	defer func() {
		mw.logger.Infow("Called method", "method", "Create", "user", true, "err", err)
	}()
	return mw.next.Create(ctx, user)
}

func (mw loggingMiddleware) Delete(ctx context.Context, name string) (err error) {
	defer func() {
		mw.logger.Infow("Called method", "method", "Delete", "name", name, "err", err)
	}()
	return mw.next.Delete(ctx, name)
}

func (mw loggingMiddleware) Get(ctx context.Context, name string) (_ *api.User, err error) {
	defer func() {
		mw.logger.Infow("Called method", "method", "Get", "name", name, "err", err)
	}()
	return mw.next.Get(ctx, name)
}

func (mw loggingMiddleware) List(ctx context.Context, pageSize int, pageToken string) (_ []*api.User, _ string, err error) {
	defer func() {
		mw.logger.Infow("Called method", "method", "List", "pageSize", pageSize, "pageToken", pageToken, "err", err)
	}()
	return mw.next.List(ctx, pageSize, pageToken)
}

func (mw loggingMiddleware) Update(ctx context.Context, name string, user *api.User) (_ *api.User, err error) {
	defer func() {
		mw.logger.Infow("Called method", "method", "Update", "name", name, "user", true, "err", err)
	}()
	return mw.next.Update(ctx, name, user)
}
