package service

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"

	"powerssl.dev/sdk/apiserver/api"
	"powerssl.dev/sdk/apiserver/certificate"

	"powerssl.dev/apiserver/internal/repository"
)

var ErrUnimplemented = status.Error(codes.Unimplemented, "Coming soon")

func New(repositories *repository.Repositories, logger log.Logger) certificate.Service {
	var svc certificate.Service
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

func NewBasicService(repositories *repository.Repositories, logger log.Logger) certificate.Service {
	return basicService{
		Repositories: repositories,
		logger:       logger,
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
