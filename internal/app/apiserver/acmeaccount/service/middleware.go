package service

import (
	"context"
	"fmt"
	"powerssl.dev/powerssl/internal/app/apiserver/unitofwork"

	"github.com/freerware/work/v4/unit"
	"github.com/go-kit/kit/log"

	"powerssl.dev/powerssl/pkg/apiserver/acmeaccount"
	"powerssl.dev/powerssl/pkg/apiserver/api"
)

type Middleware func(acmeaccount.Service) acmeaccount.Service

func UnitOfWorkMiddleware(uniter unit.Uniter, logger log.Logger) Middleware {
	return func(next acmeaccount.Service) acmeaccount.Service {
		return unitOfWorkMiddleware{logger, uniter, next}
	}
}

type unitOfWorkMiddleware struct {
	logger log.Logger
	uniter unit.Uniter
	next   acmeaccount.Service
}

func (u unitOfWorkMiddleware) Create(ctx context.Context, parent string, acmeAccount *api.ACMEAccount) (_ *api.ACMEAccount, err error) {
	var unit unit.Unit
	if unit, err = u.uniter.Unit(); err != nil {
		return nil, err
	}
	if ctx, err = unitofwork.SetUnit(ctx, unit); err != nil {
		return nil, err
	}
	defer func() {
		if err == nil {
			err = unit.Save(ctx)
		}
	}()
	return u.next.Create(ctx, parent, acmeAccount)
}

func (u unitOfWorkMiddleware) Delete(ctx context.Context, name string) (err error) {
	var unit unit.Unit
	if unit, err = u.uniter.Unit(); err != nil {
		return err
	}
	if ctx, err = unitofwork.SetUnit(ctx, unit); err != nil {
		return err
	}
	defer func() {
		if err == nil {
			err = unit.Save(ctx)
		}
	}()
	return u.next.Delete(ctx, name)
}

func (u unitOfWorkMiddleware) Get(ctx context.Context, name string) (_ *api.ACMEAccount, err error) {
	var unit unit.Unit
	if unit, err = u.uniter.Unit(); err != nil {
		return nil, err
	}
	if ctx, err = unitofwork.SetUnit(ctx, unit); err != nil {
		return nil, err
	}
	defer func() {
		if err == nil {
			err = unit.Save(ctx)
		}
	}()
	return u.next.Get(ctx, name)
}

func (u unitOfWorkMiddleware) List(ctx context.Context, parent string, pageSize int, pageToken string) (_ []*api.ACMEAccount, _ string, err error) {
	var unit unit.Unit
	if unit, err = u.uniter.Unit(); err != nil {
		return nil, "", err
	}
	if ctx, err = unitofwork.SetUnit(ctx, unit); err != nil {
		return nil, "", err
	}
	defer func() {
		if err == nil {
			err = unit.Save(ctx)
		}
	}()
	return u.next.List(ctx, parent, pageSize, pageToken)
}

func (u unitOfWorkMiddleware) Update(ctx context.Context, name string, updateMask []string, acmeAccount *api.ACMEAccount) (_ *api.ACMEAccount, err error) {
	var unit unit.Unit
	if unit, err = u.uniter.Unit(); err != nil {
		return nil, err
	}
	if ctx, err = unitofwork.SetUnit(ctx, unit); err != nil {
		return nil, err
	}
	defer func() {
		if err == nil {
			err = unit.Save(ctx)
		}
	}()
	return u.next.Update(ctx, name, updateMask, acmeAccount)
}

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
