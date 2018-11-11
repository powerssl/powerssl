package acmeserver

import (
	"context"

	"github.com/go-kit/kit/log"

	"powerssl.io/pkg/apiserver/api"
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

func (mw loggingMiddleware) Create(ctx context.Context, acmeServer *api.ACMEServer) (*api.ACMEServer, error) {
	defer func() {
		mw.logger.Log("method", "Create", "acmeServer", true)
	}()
	return mw.next.Create(ctx, acmeServer)
}

func (mw loggingMiddleware) Delete(ctx context.Context, name string) error {
	defer func() {
		mw.logger.Log("method", "Delete", "name", name)
	}()
	return mw.next.Delete(ctx, name)
}

func (mw loggingMiddleware) Get(ctx context.Context, name string) (*api.ACMEServer, error) {
	defer func() {
		mw.logger.Log("method", "Get", "name", name)
	}()
	return mw.next.Get(ctx, name)
}

func (mw loggingMiddleware) List(ctx context.Context, pageSize int, pageToken string) ([]*api.ACMEServer, string, error) {
	defer func() {
		mw.logger.Log("method", "List", "pageSize", pageSize, "pageToken", pageToken)
	}()
	return mw.next.List(ctx, pageSize, pageToken)
}

func (mw loggingMiddleware) Update(ctx context.Context, name string, acmeServer *api.ACMEServer) (*api.ACMEServer, error) {
	defer func() {
		mw.logger.Log("method", "Update", "name", name, "acmeServer", true)
	}()
	return mw.next.Update(ctx, name, acmeServer)
}
