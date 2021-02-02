package service

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/google/uuid"

	"powerssl.dev/powerssl/internal/app/apiserver/model"
	"powerssl.dev/powerssl/internal/app/apiserver/repository"
	"powerssl.dev/powerssl/internal/app/apiserver/unitofwork"
	"powerssl.dev/powerssl/pkg/apiserver/acmeserver"
	"powerssl.dev/powerssl/pkg/apiserver/api"
)

func New(repositories *repository.Repositories, logger log.Logger) acmeserver.Service {
	var svc acmeserver.Service
	{
		svc = NewBasicService(repositories, logger)
		svc = UnitOfWorkMiddleware(repositories, logger)(svc)
		svc = LoggingMiddleware(logger)(svc)
	}
	return svc
}

type basicService struct {
	*repository.Repositories
	logger log.Logger
}

func NewBasicService(repositories *repository.Repositories, logger log.Logger) acmeserver.Service {
	return basicService{
		logger:       logger,
		Repositories: repositories,
	}
}

func (s basicService) Create(ctx context.Context, apiACMEServer *api.ACMEServer) (_ *api.ACMEServer, err error) {
	acmeServer := model.NewACMEServerFromAPI(apiACMEServer, uuid.New().String())
	if err = s.ACMEServers.Add(ctx, acmeServer); err != nil {
		return nil, err
	}
	return acmeServer.ToAPI(), nil
}

func (s basicService) Delete(ctx context.Context, name string) (err error) {
	var acmeServer *model.ACMEServer
	if acmeServer, err = s.ACMEServers.FindByName(ctx, name); err != nil {
		return err
	}
	if err = s.ACMEServers.Remove(ctx, acmeServer); err != nil {
		return err
	}
	return nil
}

func (s basicService) Get(ctx context.Context, name string) (_ *api.ACMEServer, err error) {
	var acmeServer *model.ACMEServer
	if acmeServer, err = s.ACMEServers.FindByName(ctx, name); err != nil {
		return nil, err
	}
	return acmeServer.ToAPI(), nil
}

func (s basicService) List(ctx context.Context, pageSize int, pageToken string) (_ []*api.ACMEServer, _ string, err error) {
	var acmeServers *model.ACMEServers
	var nextPageToken string
	if acmeServers, nextPageToken, err = s.ACMEServers.GetRange(ctx, pageSize, pageToken); err != nil {
		return nil, "", err
	}
	return acmeServers.ToAPI(), nextPageToken, nil
}

func (s basicService) Update(ctx context.Context, name string, apiACMEServer *api.ACMEServer) (_ *api.ACMEServer, err error) {
	var acmeServer *model.ACMEServer
	if acmeServer, err = s.ACMEServers.FindByName(ctx, name); err != nil {
		return nil, err
	}
	updatedACMEServer := model.NewACMEServerFromAPI(apiACMEServer, acmeServer.ID)
	if err = unitofwork.GetUnit(ctx).Alter(updatedACMEServer); err != nil {
		return nil, err
	}
	return acmeServer.ToAPI(), nil
}
