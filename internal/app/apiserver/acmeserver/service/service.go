package service

import (
	"context"
	"strconv"

	"github.com/go-kit/kit/log"
	"github.com/go-pg/pg/v10"
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"

	"powerssl.dev/powerssl/internal/app/apiserver/acmeserver/model"
	"powerssl.dev/powerssl/pkg/apiserver/acmeserver"
	"powerssl.dev/powerssl/pkg/apiserver/api"
)

var ErrUnimplemented = status.Error(codes.Unimplemented, "Coming soon")

func New(db *pg.DB, logger log.Logger) acmeserver.Service {
	var svc acmeserver.Service
	{
		svc = NewBasicService(db, logger)
		svc = LoggingMiddleware(logger)(svc)
	}
	return svc
}

type basicService struct {
	db     *pg.DB
	logger log.Logger
}

func NewBasicService(db *pg.DB, logger log.Logger) acmeserver.Service {
	return basicService{
		db:     db,
		logger: logger,
	}
}

func (bs basicService) Create(ctx context.Context, apiacmeserver *api.ACMEServer) (*api.ACMEServer, error) {
	acmeServer := model.NewACMEServerFromAPI(apiacmeserver)
	if err := bs.db.RunInTransaction(ctx, func(tx *pg.Tx) error {
		_, err := tx.Model(acmeServer).Insert()
		return err
	}); err != nil {
		return nil, err
	}
	return acmeServer.ToAPI(), nil
}

func (bs basicService) Delete(ctx context.Context, name string) error {
	acmeServer, err := model.FindACMEServerByName(name, bs.db)
	if err != nil {
		return err
	}
	_, err = bs.db.Model(acmeServer).Delete()
	return err
}

func (bs basicService) Get(ctx context.Context, name string) (*api.ACMEServer, error) {
	acmeServer, err := model.FindACMEServerByName(name, bs.db)
	if err != nil {
		return nil, err
	}
	return acmeServer.ToAPI(), nil
}

func (bs basicService) List(ctx context.Context, pageSize int, pageToken string) ([]*api.ACMEServer, string, error) {
	if pageSize < 1 {
		pageSize = 10
	} else if pageSize > 20 {
		pageSize = 20
	}
	offset := -1
	if pageToken != "" {
		var err error
		offset, err = strconv.Atoi(pageToken)
		if err != nil {
			return nil, "", status.Error(codes.InvalidArgument, "malformed page token")
		}
	}
	var acmeServers model.ACMEServers
	if err := bs.db.Model(&acmeServers).Limit(pageSize + 1).Offset(offset).Select(); err != nil {
		return nil, "", err
	}
	var nextPageToken string
	if len(acmeServers) > pageSize {
		acmeServers = acmeServers[:len(acmeServers)-1]
		if offset == -1 {
			nextPageToken = strconv.Itoa(pageSize)
		} else {
			nextPageToken = strconv.Itoa(offset + pageSize)
		}
	}
	return acmeServers.ToAPI(), nextPageToken, nil
}

func (bs basicService) Update(ctx context.Context, name string, acmeServer *api.ACMEServer) (*api.ACMEServer, error) {
	return nil, ErrUnimplemented
}
