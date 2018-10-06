package service

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"

	"powerssl.io/pkg/api"
)

type Service interface {
	Create(ctx context.Context, ca *api.CertificateAuthority) (*api.CertificateAuthority, error)
	Delete(ctx context.Context, name string) error
	Get(ctx context.Context, name string) (*api.CertificateAuthority, error)
	List(ctx context.Context) ([]*api.CertificateAuthority, error)
	Update(ctx context.Context, ca *api.CertificateAuthority) (*api.CertificateAuthority, error)
}

func New(logger log.Logger) Service {
	var svc Service
	{
		svc = NewBasicService()
		svc = LoggingMiddleware(logger)(svc)
	}
	return svc
}

type basicService struct{}

func NewBasicService() Service {
	return basicService{}
}

func (bs basicService) Create(_ context.Context, ca *api.CertificateAuthority) (*api.CertificateAuthority, error) {
	time.Sleep(1 * time.Second)
	return ca, nil
}

func (bs basicService) Delete(_ context.Context, name string) error {
	return nil
}

func (bs basicService) Get(_ context.Context, name string) (*api.CertificateAuthority, error) {
	return nil, nil
}

func (bs basicService) List(_ context.Context) ([]*api.CertificateAuthority, error) {
	return nil, nil
}

func (bs basicService) Update(_ context.Context, ca *api.CertificateAuthority) (*api.CertificateAuthority, error) {
	return nil, nil
}
