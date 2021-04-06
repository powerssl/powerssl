package service

import (
	"context"

	"powerssl.dev/common/log"
	"powerssl.dev/sdk/apiserver/acmeserver"
	"powerssl.dev/sdk/apiserver/api"

	"powerssl.dev/apiserver/internal/model"
	"powerssl.dev/apiserver/internal/repository"
)

func New(repositories *repository.Repositories, logger log.Logger) acmeserver.Service {
	var svc acmeserver.Service
	{
		svc = NewBasicService(repositories, logger)
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
	acmeServer := model.NewACMEServerFromAPI(apiACMEServer, "")
	if err = s.ACMEServers.Insert(ctx, acmeServer); err != nil {
		return nil, err
	}
	return acmeServer.ToAPI(), nil
}

func (s basicService) Delete(ctx context.Context, name string) (err error) {
	var acmeServer *model.ACMEServer
	return s.Transaction(ctx, func(ctx context.Context) error {
		if acmeServer, err = s.ACMEServers.FindOneByName(ctx, name); err != nil {
			return err
		}
		return s.ACMEServers.Delete(ctx, acmeServer)
	})
}

func (s basicService) Get(ctx context.Context, name string) (_ *api.ACMEServer, err error) {
	var acmeServer *model.ACMEServer
	if acmeServer, err = s.ACMEServers.FindOneByName(ctx, name); err != nil {
		return nil, err
	}
	return acmeServer.ToAPI(), nil
}

func (s basicService) List(ctx context.Context, pageSize int, pageToken string) (_ []*api.ACMEServer, _ string, err error) {
	var acmeServers model.ACMEServers
	var nextPageToken string
	if acmeServers, nextPageToken, err = s.ACMEServers.FindAll(ctx, pageSize, pageToken); err != nil {
		return nil, "", err
	}
	return acmeServers.ToAPI(), nextPageToken, nil
}

func (s basicService) Update(ctx context.Context, name string, updateMask []string, apiACMEServer *api.ACMEServer) (_ *api.ACMEServer, err error) {
	var acmeServer *model.ACMEServer
	if err = s.Transaction(ctx, func(ctx context.Context) error {
		if acmeServer, err = s.ACMEServers.FindOneByName(ctx, name); err != nil {
			return err
		}
		var clauses map[string]interface{}
		if clauses, err = acmeServer.UpdateWithMask(ctx, updateMask, apiACMEServer); err != nil {
			return err
		}
		return s.ACMEServers.Update(ctx, acmeServer, clauses)
	}); err != nil {
		return nil, err
	}
	return acmeServer.ToAPI(), nil
}
