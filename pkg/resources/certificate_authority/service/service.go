package service

import (
	"context"

	"powerssl.io/pkg/api"
)

type Service interface {
	Create(ctx context.Context, ca api.CertificateAuthority) (api.CertificateAuthority, error)
	Delete(ctx context.Context, name string) error
	Get(ctx context.Context, name string) (api.CertificateAuthority, error)
	List(ctx context.Context) ([]api.CertificateAuthority, error)
	Update(ctx context.Context, ca api.CertificateAuthority) (api.CertificateAuthority, error)
}

func New() Service {
	var svc Service
	{
		svc = newBasicService()
	}
	return svc
}

type basicService struct{}

func newBasicService() Service {
	return basicService{}
}

func (bs basicService) Create(_ context.Context, ca api.CertificateAuthority) (api.CertificateAuthority, error) {
	return ca, nil
}

func (bs basicService) Delete(_ context.Context, name string) error {
	return nil
}

func (bs basicService) Get(_ context.Context, name string) (api.CertificateAuthority, error) {
	return api.CertificateAuthority{}, nil
}

func (bs basicService) List(_ context.Context) ([]api.CertificateAuthority, error) {
	return []api.CertificateAuthority{}, nil
}

func (bs basicService) Update(_ context.Context, ca api.CertificateAuthority) (api.CertificateAuthority, error) {
	return ca, nil
}
