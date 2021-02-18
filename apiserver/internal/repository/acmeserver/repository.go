package acmeserver

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"

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

var _ _interface.ACMEServerRepository = &Repository{}

func NewRepository(db *sqlx.DB, logger *zap.Logger) *Repository {
	return &Repository{
		db:     db,
		logger: logger,
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

func (r Repository) GetRange(ctx context.Context, pageSize int, pageToken string) (_ *model.ACMEServers, _ string, err error) {
	if pageSize < 1 {
		pageSize = 10
	} else if pageSize > 20 {
		pageSize = 20
	}
	offset := pageSize + 1
	var acmeServers model.ACMEServers
	if pageToken != "" {
		var createdAt time.Time
		if createdAt, _, err = decodeCursor(pageToken); err != nil {
			return nil, "", err
		}
		if err = r.db.SelectContext(ctx, &acmeServers, `
select id, display_name, directory_url, integration_name, created_at, updated_at, deleted_at
from acme_servers
where deleted_at is null and created_at = $1
order by created_at desc
limit $2
`, createdAt, offset); err != nil {
			return nil, "", err
		}
	} else {
		if err = r.db.SelectContext(ctx, &acmeServers, `
select id, display_name, directory_url, integration_name, created_at, updated_at, deleted_at
from acme_servers
where deleted_at is null
order by created_at desc
limit $1
`, offset); err != nil {
			return nil, "", err
		}
	}
	for _, acmeServer := range acmeServers {
		if err = unitofwork.GetUnit(ctx).Register(acmeServer); err != nil {
			return nil, "", err
		}
	}
	var nextPageToken string
	if len(acmeServers) > pageSize {
		acmeServer := acmeServers[len(acmeServers)-1]
		nextPageToken = encodeCursor(acmeServer.CreatedAt, acmeServer.ID)
		acmeServers = acmeServers[:len(acmeServers)-1]
	}
	return &acmeServers, nextPageToken, err
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
