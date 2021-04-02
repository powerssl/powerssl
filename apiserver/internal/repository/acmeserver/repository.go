package acmeserver

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"powerssl.dev/apiserver/internal/model"
	"powerssl.dev/apiserver/internal/repository/interface"
)

const tableName = "acme_servers"

var _ _interface.ACMEServerRepository = &Repository{}

type Repository struct {
	_interface.SQLX
	logger *zap.Logger
}

func NewRepository(interfacer _interface.SQLX, logger *zap.Logger) *Repository {
	return &Repository{
		SQLX:   interfacer,
		logger: logger,
	}
}

func (r Repository) Delete(ctx context.Context, acmeServers ...*model.ACMEServer) (err error) {
	ids := make([]string, len(acmeServers))
	for i, acmeServer := range acmeServers {
		ids[i] = acmeServer.ID
	}
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	builder := psql.
		Update(tableName).
		Set("deleted_at", sq.Expr("now()")).
		Where(sq.Eq{"id": ids}).
		Where(sq.NotEq{"deleted_at": nil})
	var sql string
	var args []interface{}
	if sql, args, err = builder.ToSql(); err != nil {
		return err
	}
	_, err = r.ExecerContext(ctx).ExecContext(ctx, sql, args...)
	return err
}

func (r Repository) FindAll(ctx context.Context, pageSize int, pageToken string) (_ model.ACMEServers, nextPageToken string, err error) {
	if pageSize < 1 {
		pageSize = 10
	} else if pageSize > 20 {
		pageSize = 20
	}
	offset := pageSize + 1
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	builder := psql.
		Select("*").
		From(tableName).
		Where(sq.Eq{"deleted_at": nil}).
		OrderBy("created_at desc").
		Limit(uint64(offset))
	if pageToken != "" {
		var createdAt time.Time
		if createdAt, _, err = decodeCursor(pageToken); err != nil {
			return nil, "", err
		}
		builder = builder.
			Where(sq.Eq{"created_at": createdAt})
	}
	var sql string
	var args []interface{}
	if sql, args, err = builder.ToSql(); err != nil {
		return nil, "", err
	}
	var acmeServers []*model.ACMEServer
	if err = sqlx.SelectContext(ctx, r.QueryerContext(ctx), &acmeServers, sql, args...); err != nil {
		return nil, "", err
	}
	if len(acmeServers) > pageSize {
		acmeServer := acmeServers[len(acmeServers)-1]
		nextPageToken = encodeCursor(acmeServer.CreatedAt, acmeServer.ID)
		acmeServers = acmeServers[:len(acmeServers)-1]
	}
	return acmeServers, "", nil
}

func (r Repository) FindOneByName(ctx context.Context, name string) (_ *model.ACMEServer, err error) {
	s := strings.Split(name, "/")
	if len(s) != 2 {
		return nil, status.Error(codes.InvalidArgument, "malformed name")
	}
	acmeServerID := s[1]
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	builder := psql.
		Select("*").
		From(tableName).
		Where(sq.Eq{
			"id":         acmeServerID,
			"deleted_at": nil,
		}).
		Limit(1)
	var sql string
	var args []interface{}
	if sql, args, err = builder.ToSql(); err != nil {
		return nil, err
	}
	var acmeServer model.ACMEServer
	if err = sqlx.GetContext(ctx, r.QueryerContext(ctx), &acmeServer, sql, args...); err != nil {
		return nil, err
	}
	return &acmeServer, nil
}

func (r Repository) Insert(ctx context.Context, acmeServers ...*model.ACMEServer) (err error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	builder := psql.
		Insert(tableName).
		Columns("id",
			"display_name",
			"directory_url",
			"integration_name")
	for _, acmeServer := range acmeServers {
		builder = builder.Values(acmeServer.ID,
			acmeServer.DisplayName,
			acmeServer.DirectoryURL,
			acmeServer.IntegrationName)
	}
	var sql string
	var args []interface{}
	if sql, args, err = builder.ToSql(); err != nil {
		return err
	}
	_, err = r.ExecerContext(ctx).ExecContext(ctx, sql, args...)
	return err
}

func (r Repository) Update(ctx context.Context, acmeServer *model.ACMEServer, clauses map[string]interface{}) (err error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	builder := psql.
		Update(tableName).
		SetMap(clauses).
		Set("updated_at", sq.Expr("now()")).
		Where(sq.Eq{
			"id":         acmeServer.ID,
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

func decodeCursor(encodedCursor string) (_ time.Time, _ string, err error) {
	var byt []byte
	if byt, err = base64.StdEncoding.DecodeString(encodedCursor); err != nil {
		return
	}

	arrStr := strings.Split(string(byt), ",")
	if len(arrStr) != 2 {
		err = errors.New("cursor is invalid")
		return
	}

	var res time.Time
	if res, err = time.Parse(time.RFC3339Nano, arrStr[0]); err != nil {
		return
	}

	return res, arrStr[1], err
}

func encodeCursor(t time.Time, uuid string) string {
	key := fmt.Sprintf("%s,%s", t.Format(time.RFC3339Nano), uuid)
	return base64.StdEncoding.EncodeToString([]byte(key))
}
