package acmeaccount

import (
	"context"
	"strings"

	"github.com/gogo/status"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"

	"powerssl.dev/apiserver/internal/model"
	"powerssl.dev/apiserver/internal/repository/interface"
	"powerssl.dev/apiserver/internal/unitofwork"
)

type Repository struct {
	db     *sqlx.DB
	logger *zap.Logger
}

var _ _interface.ACMEAccountRepository = &Repository{}

func NewRepository(db *sqlx.DB, logger *zap.Logger) *Repository {
	return &Repository{
		db:     db,
		logger: logger,
	}
}

func (r Repository) Add(ctx context.Context, acmeAccount *model.ACMEAccount) (err error) {
	return r.AddRange(ctx, &model.ACMEAccounts{acmeAccount})
}

func (r Repository) AddRange(ctx context.Context, acmeAccounts *model.ACMEAccounts) (err error) {
	unitOfWork := unitofwork.GetUnit(ctx)
	for _, acmeAccount := range *acmeAccounts {
		if err := unitOfWork.Add(acmeAccount); err != nil {
			return err
		}
	}
	return nil
}

func (r Repository) Find(ctx context.Context, predicate string) (_ *model.ACMEAccount, err error) {
	var acmeAccount model.ACMEAccount
	if err = r.db.GetContext(ctx, &acmeAccount, `select * from acme_accounts where `+predicate+` limit 1`); err != nil {
		return nil, err
	}
	if err := unitofwork.GetUnit(ctx).Register(&acmeAccount); err != nil {
		return nil, err
	}
	return &acmeAccount, err
}

func (r Repository) Get(ctx context.Context, id string) (_ *model.ACMEAccount, err error) {
	var acmeAccount model.ACMEAccount
	if err = r.db.GetContext(ctx, &acmeAccount, `select * from acme_accounts where id = $1 and deleted_at is null limit 1`, id); err != nil {
		return nil, err
	}
	if err := unitofwork.GetUnit(ctx).Register(&acmeAccount); err != nil {
		return nil, err
	}
	return &acmeAccount, err
}

func (r Repository) GetAll(ctx context.Context) (_ *model.ACMEAccounts, err error) {
	var acmeAccounts model.ACMEAccounts
	if err = r.db.SelectContext(ctx, &acmeAccounts, `select id, display_name, title, description, terms_of_service_agreed, contacts, account_url, created_at, updated_at, deleted_at from acme_accounts where deleted_at is null`); err != nil {
		return nil, err
	}
	for _, acmeAccount := range acmeAccounts {
		if err := unitofwork.GetUnit(ctx).Register(acmeAccount); err != nil {
			return nil, err
		}
	}
	return &acmeAccounts, err
}

func (r Repository) Remove(ctx context.Context, acmeAccount *model.ACMEAccount) (err error) {
	return r.RemoveRange(ctx, &model.ACMEAccounts{acmeAccount})
}

func (r Repository) RemoveRange(ctx context.Context, acmeAccounts *model.ACMEAccounts) (err error) {
	unitOfWork := unitofwork.GetUnit(ctx)
	for _, acmeAccount := range *acmeAccounts {
		if err := unitOfWork.Remove(acmeAccount); err != nil {
			return err
		}
	}
	return nil
}

func (r Repository) FindByName(ctx context.Context, name string) (_ *model.ACMEAccount, err error) {
	s := strings.Split(name, "/")
	if len(s) != 4 || s[0] != "acmeServers" || s[2] != "acmeAccounts" {
		return nil, status.Error(codes.InvalidArgument, "malformed name")
	}
	acmeServerID, acmeAccountID := s[1], s[3]

	var acmeAccount model.ACMEAccount
	if acmeServerID == "-" {
		if err = r.db.GetContext(ctx, &acmeAccount,
			`select acme_accounts.* from acme_accounts where acme_accounts.id = $1`,
			acmeAccountID); err != nil {
			return nil, err
		}
	} else {
		if err = r.db.GetContext(ctx, &acmeAccount,
			`select acme_accounts.* from acme_accounts inner join acme_servers ON acme_accounts.acme_server_id = acme_servers.id where acme_accounts.id = $1 AND acme_servers.id = $2`,
			acmeAccountID, acmeServerID); err != nil {
			return nil, err
		}

	}
	if err := unitofwork.GetUnit(ctx).Register(&acmeAccount); err != nil {
		return nil, err
	}
	return &acmeAccount, err
}
