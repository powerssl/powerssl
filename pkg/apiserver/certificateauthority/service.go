package certificateauthority

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"

	"powerssl.io/pkg/apiserver/api"
	controllerclient "powerssl.io/pkg/controller/client"
)

type Service interface {
	Create(ctx context.Context, certificateAuthority *api.CertificateAuthority) (*api.CertificateAuthority, error)
	Delete(ctx context.Context, name string) error
	Get(ctx context.Context, name string) (*api.CertificateAuthority, error)
	List(ctx context.Context, pageSize int, pageToken string) ([]*api.CertificateAuthority, string, error)
	Update(ctx context.Context, name string, certificateAuthority *api.CertificateAuthority) (*api.CertificateAuthority, error)
}

func New(db *gorm.DB, logger log.Logger, _ *controllerclient.GRPCClient) Service {
	var svc Service
	{
		svc = NewBasicService(db)
		svc = LoggingMiddleware(logger)(svc)
	}
	return svc
}

type basicService struct {
	db *gorm.DB
}

func NewBasicService(db *gorm.DB) Service {
	return basicService{
		db: db,
	}
}

func (bs basicService) Create(_ context.Context, certificateAuthority *api.CertificateAuthority) (*api.CertificateAuthority, error) {
	return nil, nil
}

func (bs basicService) Delete(_ context.Context, name string) error {
	return nil
}

func (bs basicService) Get(_ context.Context, name string) (*api.CertificateAuthority, error) {
	return nil, nil
}

func (bs basicService) List(_ context.Context, pageSize int, pageToken string) ([]*api.CertificateAuthority, string, error) {
	return nil, "", nil
}

func (bs basicService) Update(_ context.Context, name string, certificateAuthority *api.CertificateAuthority) (*api.CertificateAuthority, error) {
	return nil, nil
}
