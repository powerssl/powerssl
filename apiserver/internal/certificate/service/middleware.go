package service

import (
	"context"

	"github.com/go-kit/kit/log"

	"powerssl.dev/sdk/apiserver/api"
	"powerssl.dev/sdk/apiserver/certificate"
)

type Middleware func(certificate.Service) certificate.Service

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next certificate.Service) certificate.Service {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   certificate.Service
}

func (mw loggingMiddleware) Create(ctx context.Context, certificate *api.Certificate) (*api.Certificate, error) {
	defer func() {
		mw.logger.Log("method", "Create", "certificate", true)
	}()
	return mw.next.Create(ctx, certificate)
}

func (mw loggingMiddleware) Delete(ctx context.Context, name string) error {
	defer func() {
		mw.logger.Log("method", "Delete", "name", name)
	}()
	return mw.next.Delete(ctx, name)
}

func (mw loggingMiddleware) Get(ctx context.Context, name string) (*api.Certificate, error) {
	defer func() {
		mw.logger.Log("method", "Get", "name", name)
	}()
	return mw.next.Get(ctx, name)
}

func (mw loggingMiddleware) List(ctx context.Context, pageSize int, pageToken string) ([]*api.Certificate, string, error) {
	defer func() {
		mw.logger.Log("method", "List", "pageSize", pageSize, "pageToken", pageToken)
	}()
	return mw.next.List(ctx, pageSize, pageToken)
}

func (mw loggingMiddleware) Update(ctx context.Context, name string, certificate *api.Certificate) (*api.Certificate, error) {
	defer func() {
		mw.logger.Log("method", "Update", "name", name, "certificate", true)
	}()
	return mw.next.Update(ctx, name, certificate)
}
