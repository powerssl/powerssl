package service

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"

	"powerssl.dev/powerssl/pkg/apiserver/acmeaccount"
	"powerssl.dev/powerssl/pkg/apiserver/api"
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

func (mw loggingMiddleware) Create(ctx context.Context, parent string, acmeAccount *api.ACMEAccount) (*api.ACMEAccount, error) {
	defer func() {
		mw.logger.Log("method", "Create", "parent", parent, "acmeAccount", true)
	}()
	return mw.next.Create(ctx, parent, acmeAccount)
}

func (mw loggingMiddleware) Delete(ctx context.Context, name string) error {
	defer func() {
		mw.logger.Log("method", "Delete", "name", name)
	}()
	return mw.next.Delete(ctx, name)
}

func (mw loggingMiddleware) Get(ctx context.Context, name string) (*api.ACMEAccount, error) {
	defer func() {
		mw.logger.Log("method", "Get", "name", name)
	}()
	return mw.next.Get(ctx, name)
}

func (mw loggingMiddleware) List(ctx context.Context, parent string, pageSize int, pageToken string) ([]*api.ACMEAccount, string, error) {
	defer func() {
		mw.logger.Log("method", "List", "parent", parent, "pageSize", pageSize, "pageToken", pageToken)
	}()
	return mw.next.List(ctx, parent, pageSize, pageToken)
}

func (mw loggingMiddleware) Update(ctx context.Context, name string, updateMask []string, acmeAccount *api.ACMEAccount) (*api.ACMEAccount, error) {
	defer func() {
		mw.logger.Log("method", "Update", "name", name, "updateMask", fmt.Sprintf("%+v", updateMask), "acmeAccount", true)
	}()
	return mw.next.Update(ctx, name, updateMask, acmeAccount)
}
