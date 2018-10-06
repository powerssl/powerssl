package service

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"

	"powerssl.io/pkg/api"
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

func (mw loggingMiddleware) Create(ctx context.Context, in *api.CertificateAuthority) (out *api.CertificateAuthority, err error) {
	defer func() {
		mw.logger.Log("method", "Create", "in", fmt.Sprintf("%+v", in), "out", fmt.Sprintf("%+v", out), "err", err)
	}()
	return mw.next.Create(ctx, in)
}

func (mw loggingMiddleware) Delete(ctx context.Context, name string) error {
	return mw.next.Delete(ctx, name)
}

func (mw loggingMiddleware) Get(ctx context.Context, name string) (*api.CertificateAuthority, error) {
	return mw.next.Get(ctx, name)
}

func (mw loggingMiddleware) List(ctx context.Context) ([]*api.CertificateAuthority, error) {
	return mw.next.List(ctx)
}

func (mw loggingMiddleware) Update(ctx context.Context, ca *api.CertificateAuthority) (*api.CertificateAuthority, error) {
	return mw.next.Update(ctx, ca)
}
