package service

import (
	"context"

	"github.com/gogo/status"
	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc/codes"

	"powerssl.dev/common/log"
	"powerssl.dev/sdk/apiserver/api"
	"powerssl.dev/sdk/apiserver/certificate"
)

var ErrUnimplemented = status.Error(codes.Unimplemented, "Coming soon")

func New(db *pgx.Conn, logger log.Logger) certificate.Service {
	var svc certificate.Service
	{
		svc = NewBasicService(db, logger)
		svc = LoggingMiddleware(logger)(svc)
	}
	return svc
}

type basicService struct {
	db     *pgx.Conn
	logger log.Logger
}

func NewBasicService(db *pgx.Conn, logger log.Logger) certificate.Service {
	return basicService{
		db:     db,
		logger: logger,
	}
}

func (bs basicService) Create(ctx context.Context, apicertificate *api.Certificate) (*api.Certificate, error) {
	return nil, ErrUnimplemented
}

func (bs basicService) Delete(ctx context.Context, name string) error {
	return ErrUnimplemented
}

func (bs basicService) Get(ctx context.Context, name string) (*api.Certificate, error) {
	return nil, ErrUnimplemented
}

func (bs basicService) List(ctx context.Context, pageSize int, pageToken string) ([]*api.Certificate, string, error) {
	return nil, "", ErrUnimplemented
}

func (bs basicService) Update(ctx context.Context, name string, certificate *api.Certificate) (*api.Certificate, error) {
	return nil, ErrUnimplemented
}
