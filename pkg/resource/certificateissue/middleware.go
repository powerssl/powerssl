package certificateissue

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

func (mw loggingMiddleware) Create(ctx context.Context, parent string, certificateIssue *api.CertificateIssue) (*api.CertificateIssue, error) {
	defer func() {
		mw.logger.Log("method", "Create", "parent", parent, "certificateIssue", fmt.Sprintf("%+v", certificateIssue))
	}()
	return mw.next.Create(ctx, parent, certificateIssue)
}

func (mw loggingMiddleware) Delete(ctx context.Context, name string) error {
	defer func() {
		mw.logger.Log("method", "Delete", "name", name)
	}()
	return mw.next.Delete(ctx, name)
}

func (mw loggingMiddleware) Get(ctx context.Context, name string) (*api.CertificateIssue, error) {
	defer func() {
		mw.logger.Log("method", "Get", "name", name)
	}()
	return mw.next.Get(ctx, name)
}

func (mw loggingMiddleware) List(ctx context.Context, parent string, pageSize int, pageToken string) ([]*api.CertificateIssue, string, error) {
	defer func() {
		mw.logger.Log("method", "List", "parent", parent, "pageSize", pageSize, "pageToken", pageToken)
	}()
	return mw.next.List(ctx, parent, pageSize, pageToken)
}

func (mw loggingMiddleware) Update(ctx context.Context, name string, certificateIssue *api.CertificateIssue) (*api.CertificateIssue, error) {
	defer func() {
		mw.logger.Log("method", "Update", "name", name, "certificateIssue", fmt.Sprintf("%+v", certificateIssue))
	}()
	return mw.next.Update(ctx, name, certificateIssue)
}
