package service

import (
	"context"
	"fmt"

	"powerssl.dev/common/log"
	"powerssl.dev/sdk/apiserver/acmeaccount"
	"powerssl.dev/sdk/apiserver/api"
)

type Middleware func(acmeaccount.Service) acmeaccount.Service

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next acmeaccount.Service) acmeaccount.Service {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   acmeaccount.Service
}

func (mw loggingMiddleware) Create(ctx context.Context, parent string, acmeAccount *api.ACMEAccount) (_ *api.ACMEAccount, err error) {
	defer func() {
		mw.logger.Infow("Called method", "method", "Create", "parent", parent, "acmeAccount", true, "err", err)
	}()
	return mw.next.Create(ctx, parent, acmeAccount)
}

func (mw loggingMiddleware) Delete(ctx context.Context, name string) (err error) {
	defer func() {
		mw.logger.Infow("Called method", "method", "Delete", "name", name, "err", err)
	}()
	return mw.next.Delete(ctx, name)
}

func (mw loggingMiddleware) Get(ctx context.Context, name string) (_ *api.ACMEAccount, err error) {
	defer func() {
		mw.logger.Infow("Called method", "method", "Get", "name", name, "err", err)
	}()
	return mw.next.Get(ctx, name)
}

func (mw loggingMiddleware) List(ctx context.Context, parent string, pageSize int, pageToken string) (_ []*api.ACMEAccount, _ string, err error) {
	defer func() {
		mw.logger.Infow("Called method", "method", "List", "parent", parent, "pageSize", pageSize, "pageToken", pageToken, "err", err)
	}()
	return mw.next.List(ctx, parent, pageSize, pageToken)
}

func (mw loggingMiddleware) Update(ctx context.Context, name string, updateMask []string, acmeAccount *api.ACMEAccount) (_ *api.ACMEAccount, err error) {
	defer func() {
		mw.logger.Infow("Called method", "method", "Update", "name", name, "updateMask", fmt.Sprintf("%+v", updateMask), "acmeAccount", true, "err", err)
	}()
	return mw.next.Update(ctx, name, updateMask, acmeAccount)
}
