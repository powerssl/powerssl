package service

import (
	"context"

	"powerssl.dev/common/log"
	"powerssl.dev/sdk/apiserver/acmeserver"
	"powerssl.dev/sdk/apiserver/api"
)

type Middleware func(acmeserver.Service) acmeserver.Service

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next acmeserver.Service) acmeserver.Service {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   acmeserver.Service
}

func (mw loggingMiddleware) Create(ctx context.Context, acmeServer *api.ACMEServer) (_ *api.ACMEServer, err error) {
	defer func() {
		mw.logger.Infow("Called method", "method", "Create", "acmeServer", true, "err", err)
	}()
	return mw.next.Create(ctx, acmeServer)
}

func (mw loggingMiddleware) Delete(ctx context.Context, name string) (err error) {
	defer func() {
		mw.logger.Infow("Called method", "method", "Delete", "name", name, "err", err)
	}()
	return mw.next.Delete(ctx, name)
}

func (mw loggingMiddleware) Get(ctx context.Context, name string) (_ *api.ACMEServer, err error) {
	defer func() {
		mw.logger.Infow("Called method", "method", "Get", "name", name, "err", err)
	}()
	return mw.next.Get(ctx, name)
}

func (mw loggingMiddleware) List(ctx context.Context, pageSize int, pageToken string) (_ []*api.ACMEServer, _ string, err error) {
	defer func() {
		mw.logger.Infow("Called method", "method", "List", "pageSize", pageSize, "pageToken", pageToken, "err", err)
	}()
	return mw.next.List(ctx, pageSize, pageToken)
}

func (mw loggingMiddleware) Update(ctx context.Context, name string, updateMask []string, acmeServer *api.ACMEServer) (_ *api.ACMEServer, err error) {
	defer func() {
		mw.logger.Infow("Called method", "method", "Update", "name", name, "acmeServer", true, "err", err)
	}()
	return mw.next.Update(ctx, name, updateMask, acmeServer)
}
