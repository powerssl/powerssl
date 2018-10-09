package certificateissue

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"

	"powerssl.io/pkg/api"
)

type Service interface {
	Create(ctx context.Context, certificateIssue *api.CertificateIssue) (*api.CertificateIssue, error)
	Delete(ctx context.Context, name string) error
	Get(ctx context.Context, name string) (*api.CertificateIssue, error)
	List(ctx context.Context) ([]*api.CertificateIssue, error)
	Update(ctx context.Context, certificateIssue *api.CertificateIssue) (*api.CertificateIssue, error)
}

func New(db *gorm.DB, logger log.Logger) Service {
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

func (bs basicService) Create(_ context.Context, certificateIssue *api.CertificateIssue) (*api.CertificateIssue, error) {
	return nil, nil
}

func (bs basicService) Delete(_ context.Context, name string) error {
	return nil
}

func (bs basicService) Get(_ context.Context, name string) (*api.CertificateIssue, error) {
	return nil, nil
}

func (bs basicService) List(_ context.Context) ([]*api.CertificateIssue, error) {
	return nil, nil
}

func (bs basicService) Update(_ context.Context, certificateIssue *api.CertificateIssue) (*api.CertificateIssue, error) {
	return nil, nil
}
