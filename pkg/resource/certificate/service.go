package certificate

import (
	"context"

	"github.com/go-kit/kit/log"

	"powerssl.io/pkg/api"
)

type Service interface {
	Create(ctx context.Context, certificate *api.Certificate) (*api.Certificate, error)
	Delete(ctx context.Context, name string) error
	Get(ctx context.Context, name string) (*api.Certificate, error)
	List(ctx context.Context) ([]*api.Certificate, error)
	Update(ctx context.Context, certificate *api.Certificate) (*api.Certificate, error)
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

func (bs basicService) Create(_ context.Context, certificate *api.Certificate) (*api.Certificate, error) {
	return nil, nil
}

func (bs basicService) Delete(_ context.Context, name string) error {
	return nil
}

func (bs basicService) Get(_ context.Context, name string) (*api.Certificate, error) {
	return nil, nil
}

func (bs basicService) List(_ context.Context) ([]*api.Certificate, error) {
	return nil, nil
}

func (bs basicService) Update(_ context.Context, certificate *api.Certificate) (*api.Certificate, error) {
	return nil, nil
}
