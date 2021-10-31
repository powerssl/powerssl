package service

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"

	"powerssl.dev/common/log"
	"powerssl.dev/sdk/apiserver/acmeserver"
	"powerssl.dev/sdk/apiserver/api"

	"powerssl.dev/apiserver/internal/repository"
)

func New(db *pgx.Conn, logger log.Logger) acmeserver.Service {
	var svc acmeserver.Service
	{
		svc = NewBasicService(db, logger)
		svc = LoggingMiddleware(logger)(svc)
	}
	return svc
}

type basicService struct {
	Queries *repository.Queries
	db      *pgx.Conn
	logger  log.Logger
}

func NewBasicService(db *pgx.Conn, logger log.Logger) acmeserver.Service {
	return basicService{
		Queries: repository.New(db),
		db:      db,
		logger:  logger,
	}
}

func (s basicService) Create(ctx context.Context, apiACMEServer *api.ACMEServer) (_ *api.ACMEServer, err error) {
	acmeServer, err := s.Queries.CreateACMEServerFromAPI(ctx, apiACMEServer)
	if err != nil {
		return nil, err
	}
	return acmeServer.ToAPI(), nil
}

func (s basicService) Delete(ctx context.Context, name string) error {
	queries, rollback, err := s.Queries.NewTx(ctx)
	defer rollback(&err)
	n := strings.Split(name, "/")
	id, err := uuid.Parse(n[1])
	if err != nil {
		return err
	}
	acmeServer, err := queries.GetACMEServer(ctx, id)
	if err != nil {
		return err
	}
	if err = queries.DeleteACMEServer(ctx, acmeServer.ID); err != nil {
		return err
	}
	return queries.Tx().Commit(ctx)
}

func (s basicService) Get(ctx context.Context, name string) (_ *api.ACMEServer, err error) {
	n := strings.Split(name, "/")
	id, err := uuid.Parse(n[1])
	if err != nil {
		return nil, err
	}
	acmeServer, err := s.Queries.GetACMEServer(ctx, id)
	if err != nil {
		return nil, err
	}
	return acmeServer.ToAPI(), nil
}

func (s basicService) List(ctx context.Context, pageSize int, pageToken string) (_ []*api.ACMEServer, _ string, err error) {
	var nextPageToken string
	acmeServers, err := s.Queries.ListACMEServers(ctx, repository.ListACMEServersParams{
		SqlOrder:  "created_at",
		SqlOffset: 0,
		SqlLimit:  int32(pageSize),
	})
	if err != nil {
		return nil, "", err
	}
	return repository.AcmeServers(acmeServers).ToAPI(), nextPageToken, nil
}

func (s basicService) Update(ctx context.Context, name string, updateMask []string, apiACMEServer *api.ACMEServer) (_ *api.ACMEServer, err error) {
	queries, rollback, err := s.Queries.NewTx(ctx)
	defer rollback(&err)
	n := strings.Split(name, "/")
	id, err := uuid.Parse(n[1])
	if err != nil {
		return nil, err
	}
	acmeServer, err := queries.UpdateACMEServerWithMask(ctx, id, updateMask, apiACMEServer)
	if err != nil {
		return nil, err
	}
	if err = queries.Tx().Commit(ctx); err != nil {
		return nil, err
	}
	return acmeServer.ToAPI(), nil
}
