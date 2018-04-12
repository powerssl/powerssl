package service

import (
	"context"

	"github.com/go-kit/kit/log"

	domainmodel "powerssl.io/pkg/domain"
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

func (mw loggingMiddleware) Create(ctx context.Context, domain domainmodel.Domain) (_ domainmodel.Domain, err error) {
	defer func() {
		mw.logger.Log("method", "Create", "err", err)
	}()
	return mw.next.Create(ctx, domain)
}

func (mw loggingMiddleware) Delete(ctx context.Context, name string) (err error) {
	defer func() {
		mw.logger.Log("method", "Delete", "name", name, "err", err)
	}()
	return mw.next.Delete(ctx, name)
}

func (mw loggingMiddleware) Get(ctx context.Context, name string) (_ domainmodel.Domain, err error) {
	defer func() {
		mw.logger.Log("method", "Get", "name", name, "err", err)
	}()
	return mw.next.Get(ctx, name)
}

func (mw loggingMiddleware) List(ctx context.Context, pageSize int, pageToken string) (_ []domainmodel.Domain, err error) {
	defer func() {
		mw.logger.Log("method", "List", "err", err)
	}()
	return mw.next.List(ctx, pageSize, pageToken)
}

func (mw loggingMiddleware) Update(ctx context.Context, domain domainmodel.Domain, updateMask string) (_ domainmodel.Domain, err error) {
	defer func() {
		mw.logger.Log("method", "Update", "err", err)
	}()
	return mw.next.Update(ctx, domain, updateMask)
}
