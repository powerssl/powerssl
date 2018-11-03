package certificateauthority

import (
	"context"
	"fmt"

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

func (mw loggingMiddleware) Create(ctx context.Context, certificateAuthority *api.CertificateAuthority) (*api.CertificateAuthority, error) {
	defer func() {
		mw.logger.Log("method", "Create", "certificateAuthority", fmt.Sprintf("%+v", certificateAuthority))
	}()
	return mw.next.Create(ctx, certificateAuthority)
}

func (mw loggingMiddleware) Delete(ctx context.Context, name string) error {
	defer func() {
		mw.logger.Log("method", "Delete", "name", name)
	}()
	return mw.next.Delete(ctx, name)
}

func (mw loggingMiddleware) Get(ctx context.Context, name string) (*api.CertificateAuthority, error) {
	defer func() {
		mw.logger.Log("method", "Get", "name", name)
	}()
	return mw.next.Get(ctx, name)
}

func (mw loggingMiddleware) List(ctx context.Context, pageSize int, pageToken string) ([]*api.CertificateAuthority, string, error) {
	defer func() {
		mw.logger.Log("method", "List")
	}()
	return mw.next.List(ctx, pageSize, pageToken)
}

func (mw loggingMiddleware) Update(ctx context.Context, name string, certificateAuthority *api.CertificateAuthority) (*api.CertificateAuthority, error) {
	defer func() {
		mw.logger.Log("method", "Update", "certificateAuthority", fmt.Sprintf("%+v", certificateAuthority))
	}()
	return mw.next.Update(ctx, name, certificateAuthority)
}
