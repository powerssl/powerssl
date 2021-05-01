package acmeaccount

import (
	"context"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"powerssl.dev/common/log"

	"powerssl.dev/apiserver/internal/model"
	"powerssl.dev/apiserver/internal/repository/interface"
)

const tableName = "acme_accounts"

var _ _interface.ACMEAccountRepository = &Repository{}

func prefixWithTableName(str string) string {
	return tableName + "." + str
}

type Repository struct {
	_interface.SQLX
	logger log.Logger
}

func NewRepository(interfacer _interface.SQLX, logger log.Logger) *Repository {
	return &Repository{
		SQLX:   interfacer,
		logger: logger,
	}
}

func (r Repository) Delete(ctx context.Context, acmeAccounts ...*model.ACMEAccount) (err error) {
	ids := make([]string, len(acmeAccounts))
	for i, acmeAccount := range acmeAccounts {
		ids[i] = acmeAccount.ID
	}
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	builder := psql.
		Update(tableName).
		Set("deleted_at", sq.Expr("now()")).
		Where(sq.Eq{"id": ids}).
		Where(sq.Eq{"deleted_at": nil})
	var sql string
	var args []interface{}
	if sql, args, err = builder.ToSql(); err != nil {
		return err
	}
	_, err = r.ExecerContext(ctx).ExecContext(ctx, sql, args...)
	return err
}

func (r Repository) FindAll(ctx context.Context) (_ model.ACMEAccounts, err error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	builder := psql.
		Select("*").
		From(tableName).
		Where(sq.Eq{"deleted_at": nil})
	var sql string
	var args []interface{}
	if sql, args, err = builder.ToSql(); err != nil {
		return nil, err
	}
	var acmeAccounts []*model.ACMEAccount
	if err = sqlx.SelectContext(ctx, r.QueryerContext(ctx), &acmeAccounts, sql, args...); err != nil {
		return nil, err
	}
	return acmeAccounts, nil
}

func (r Repository) FindAllByParent(ctx context.Context, parent string) (_ model.ACMEAccounts, err error) {
	var acmeServerID string
	if parent != "-" {
		s := strings.Split(parent, "/")
		if len(s) != 2 {
			return nil, status.Error(codes.InvalidArgument, "malformed name")
		}
		acmeServerID = s[1]
	}
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	builder := psql.
		Select("*").
		From(tableName).
		Where(sq.Eq{"deleted_at": nil})
	if acmeServerID != "" {
		builder = builder.Where(sq.Eq{"acme_server_id": acmeServerID})
	}
	var sql string
	var args []interface{}
	if sql, args, err = builder.ToSql(); err != nil {
		return nil, err
	}
	var acmeAccounts []*model.ACMEAccount
	if err = sqlx.SelectContext(ctx, r.QueryerContext(ctx), &acmeAccounts, sql, args...); err != nil {
		return nil, err
	}
	return acmeAccounts, nil
}

func (r Repository) FindOneByName(ctx context.Context, name string) (_ *model.ACMEAccount, err error) {
	s := strings.Split(name, "/")
	if len(s) != 4 || s[0] != "acmeServers" || s[2] != "acmeAccounts" {
		return nil, status.Error(codes.InvalidArgument, "malformed name")
	}
	acmeServerID, acmeAccountID := s[1], s[3]
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	builder := psql.
		Select(prefixWithTableName("*")).
		From(tableName).
		Where(sq.Eq{
			prefixWithTableName("id"):         acmeAccountID,
			prefixWithTableName("deleted_at"): nil,
		}).
		Limit(1)
	if acmeServerID != "-" {
		builder = builder.
			InnerJoin("acme_servers on acme_accounts.acme_server_id = acme_servers.id").
			Where(sq.Eq{
				"acme_servers.id":         acmeServerID,
				"acme_servers.deleted_at": nil,
			})
	}
	var sql string
	var args []interface{}
	if sql, args, err = builder.ToSql(); err != nil {
		return nil, err
	}
	var acmeAccount model.ACMEAccount
	if err = sqlx.GetContext(ctx, r.QueryerContext(ctx), &acmeAccount, sql, args...); err != nil {
		return nil, err
	}
	acmeAccount.ACMEServer = &model.ACMEServer{
		ID: acmeAccount.ACMEServerID,
	}
	return &acmeAccount, nil
}

func (r Repository) Insert(ctx context.Context, acmeAccounts ...*model.ACMEAccount) (err error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	builder := psql.
		Insert(tableName).
		Columns("id",
			"acme_server_id",
			"display_name",
			"title",
			"description",
			"terms_of_service_agreed",
			"contacts",
			"account_url")
	for _, acmeAccount := range acmeAccounts {
		builder = builder.Values(acmeAccount.ID,
			acmeAccount.ACMEServerID,
			acmeAccount.DisplayName,
			acmeAccount.Title,
			acmeAccount.Description,
			acmeAccount.TermsOfServiceAgreed,
			acmeAccount.Contacts,
			acmeAccount.AccountURL)
	}
	var sql string
	var args []interface{}
	if sql, args, err = builder.ToSql(); err != nil {
		return err
	}
	_, err = r.ExecerContext(ctx).ExecContext(ctx, sql, args...)
	return err
}

func (r Repository) Update(ctx context.Context, acmeAccount *model.ACMEAccount, clauses map[string]interface{}) (err error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	builder := psql.
		Update(tableName).
		SetMap(clauses).
		Set("updated_at", sq.Expr("now()")).
		Where(sq.Eq{
			"id":         acmeAccount.ID,
			"deleted_at": nil,
		})
	var sql string
	var args []interface{}
	if sql, args, err = builder.ToSql(); err != nil {
		return err
	}
	if _, err = r.ExecerContext(ctx).ExecContext(ctx, sql, args...); err != nil {
		return err
	}
	return nil
}
