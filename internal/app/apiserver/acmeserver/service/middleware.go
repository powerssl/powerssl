package service

import (
	"context"

	"github.com/freerware/work/v4/unit"
	"github.com/go-kit/kit/log"

	"powerssl.dev/powerssl/internal/app/apiserver/repository"
	"powerssl.dev/powerssl/internal/app/apiserver/unitofwork"
	"powerssl.dev/powerssl/pkg/apiserver/acmeserver"
	"powerssl.dev/powerssl/pkg/apiserver/api"
)

type Middleware func(acmeserver.Service) acmeserver.Service

func UnitOfWorkMiddleware(repositories *repository.Repositories, logger log.Logger) Middleware {
	return func(next acmeserver.Service) acmeserver.Service {
		return unitOfWorkMiddleware{repositories, logger, next}
	}
}

type unitOfWorkMiddleware struct {
	*repository.Repositories
	logger log.Logger
	next   acmeserver.Service
}

func (u unitOfWorkMiddleware) Create(ctx context.Context, acmeServer *api.ACMEServer) (_ *api.ACMEServer, err error) {
	var unit unit.Unit
	if unit, err = u.UnitOfWork(); err != nil {
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
	return u.next.Create(ctx, acmeServer)
}

func (u unitOfWorkMiddleware) Delete(ctx context.Context, name string) (err error) {
	var unit unit.Unit
	if unit, err = u.UnitOfWork(); err != nil {
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

func (u unitOfWorkMiddleware) Get(ctx context.Context, name string) (_ *api.ACMEServer, err error) {
	var unit unit.Unit
	if unit, err = u.UnitOfWork(); err != nil {
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

func (u unitOfWorkMiddleware) List(ctx context.Context, pageSize int, pageToken string) (_ []*api.ACMEServer, _ string, err error) {
	var unit unit.Unit
	if unit, err = u.UnitOfWork(); err != nil {
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
	return u.next.List(ctx, pageSize, pageToken)
}

func (u unitOfWorkMiddleware) Update(ctx context.Context, name string, acmeServer *api.ACMEServer) (_ *api.ACMEServer, err error) {
	var unit unit.Unit
	if unit, err = u.UnitOfWork(); err != nil {
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
	return u.next.Update(ctx, name, acmeServer)
}

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next acmeserver.Service) acmeserver.Service {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   acmeserver.Service
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
