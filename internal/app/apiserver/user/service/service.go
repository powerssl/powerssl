package service

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/gogo/status"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"

	"powerssl.dev/powerssl/pkg/apiserver/api"
	"powerssl.dev/powerssl/pkg/apiserver/user"
)

var ErrUnimplemented = status.Error(codes.Unimplemented, "Coming soon")

func New(db *sqlx.DB, logger log.Logger) user.Service {
	var svc user.Service
	{
		svc = NewBasicService(db, logger)
		svc = LoggingMiddleware(logger)(svc)
	}
	return svc
}

type basicService struct {
	db     *sqlx.DB
	logger log.Logger
}

func NewBasicService(db *sqlx.DB, logger log.Logger) user.Service {
	return basicService{
		db:     db,
		logger: logger,
	}
}

func (bs basicService) Create(ctx context.Context, apiUser *api.User) (*api.User, error) {
	return nil, ErrUnimplemented
}

func (bs basicService) Delete(ctx context.Context, name string) error {
	return ErrUnimplemented
}

func (bs basicService) Get(ctx context.Context, name string) (*api.User, error) {
	return nil, ErrUnimplemented
}

func (bs basicService) List(ctx context.Context, pageSize int, pageToken string) ([]*api.User, string, error) {
	return nil, "", ErrUnimplemented
}

func (bs basicService) Update(ctx context.Context, name string, user *api.User) (*api.User, error) {
	return nil, ErrUnimplemented
}
