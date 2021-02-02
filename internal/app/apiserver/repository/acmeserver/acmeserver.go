package acmeserver

import (
	"context"
	"strings"

	"github.com/gogo/status"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"

	"powerssl.dev/powerssl/internal/app/apiserver/model"
	"powerssl.dev/powerssl/internal/app/apiserver/repository"
	"powerssl.dev/powerssl/internal/app/apiserver/unitofwork"
)

type Repository struct {
	db *sqlx.DB
}

var _ repository.ACMEServerRepository = &Repository{}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r Repository) Add(ctx context.Context, acmeServer *model.ACMEServer) (err error) {
	return r.AddRange(ctx, &model.ACMEServers{acmeServer})
}

func (r Repository) AddRange(ctx context.Context, acmeServers *model.ACMEServers) (err error) {
	unitOfWork := unitofwork.GetUnit(ctx)
	for _, acmeServer := range *acmeServers {
		if err := unitOfWork.Add(acmeServer); err != nil {
			return err
		}
	}
	return nil
}

func (r Repository) Find(ctx context.Context, predicate string) (_ *model.ACMEServer, err error) {
	var acmeServer model.ACMEServer
	if err = r.db.GetContext(ctx, &acmeServer, `select * from acme_servers where `+predicate+` limit 1`); err != nil {
		return nil, err
	}
	if err := unitofwork.GetUnit(ctx).Register(&acmeServer); err != nil {
		return nil, err
	}
	return &acmeServer, err
}

func (r Repository) Get(ctx context.Context, id string) (_ *model.ACMEServer, err error) {
	var acmeServer model.ACMEServer
	if err = r.db.GetContext(ctx, &acmeServer, `select * from acme_servers where id = $1 and deleted_at is null limit 1`, id); err != nil {
		return nil, err
	}
	if err := unitofwork.GetUnit(ctx).Register(&acmeServer); err != nil {
		return nil, err
	}
	return &acmeServer, err
}

func (r Repository) GetAll(ctx context.Context) (_ *model.ACMEServers, err error) {
	var acmeServers model.ACMEServers
	if err = r.db.SelectContext(ctx, &acmeServers, `select id, display_name, directory_url, integration_name, created_at, updated_at, deleted_at from acme_servers where deleted_at is null`); err != nil {
		return nil, err
	}
	for _, acmeServer := range acmeServers {
		if err := unitofwork.GetUnit(ctx).Register(acmeServer); err != nil {
			return nil, err
		}
	}
	return &acmeServers, err
}

func (r Repository) Remove(ctx context.Context, acmeServer *model.ACMEServer) (err error) {
	return r.RemoveRange(ctx, &model.ACMEServers{acmeServer})
}

func (r Repository) RemoveRange(ctx context.Context, acmeServers *model.ACMEServers) (err error) {
	unitOfWork := unitofwork.GetUnit(ctx)
	for _, acmeServer := range *acmeServers {
		if err := unitOfWork.Remove(acmeServer); err != nil {
			return err
		}
	}
	return nil
}

func (r Repository) FindByName(ctx context.Context, name string) (acmeServer *model.ACMEServer, err error) {
	s := strings.Split(name, "/")
	if len(s) != 2 {
		return nil, status.Error(codes.InvalidArgument, "malformed name")
	}
	return r.Get(ctx, s[1])
}
