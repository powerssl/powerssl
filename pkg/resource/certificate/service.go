package certificate

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"

	"powerssl.io/pkg/api"
)

type Service interface {
	Create(ctx context.Context, certificate *api.Certificate) (*api.Certificate, error)
	Delete(ctx context.Context, name string) error
	Get(ctx context.Context, name string) (*api.Certificate, error)
	List(ctx context.Context, pageSize int, pageToken string) ([]*api.Certificate, string, error)
	Update(ctx context.Context, name string, certificate *api.Certificate) (*api.Certificate, error)
}

func New(db *gorm.DB, logger log.Logger) Service {
	db.AutoMigrate(&Certificate{})

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

func (bs basicService) Create(_ context.Context, certificate *api.Certificate) (*api.Certificate, error) {
	cert := NewCertificateFromAPI(certificate)
	if err := bs.db.Create(cert).Error; err != nil {
		return nil, err
	}
	return cert.ToAPI(), nil
}

func (bs basicService) Delete(_ context.Context, name string) error {
	certificate, err := FindCertificateByName(name, bs.db)
	if err != nil {
		return err
	}
	return bs.db.Delete(certificate).Error
}

func (bs basicService) Get(_ context.Context, name string) (*api.Certificate, error) {
	certificate, err := FindCertificateByName(name, bs.db)
	if err != nil {
		return nil, err
	}
	return certificate.ToAPI(), nil
}

func (bs basicService) List(_ context.Context, pageSize int, pageToken string) ([]*api.Certificate, string, error) {
	certificates, err := FindCertificates(bs.db)
	if err != nil {
		return nil, "", err
	}
	return certificates.ToAPI(), "", nil
}

func (bs basicService) Update(_ context.Context, name string, certificate *api.Certificate) (*api.Certificate, error) {
	cert, err := FindCertificateByName(name, bs.db)
	if err != nil {
		return nil, err
	}
	if err := bs.db.Model(cert).Updates(NewCertificateFromAPI(certificate)).Error; err != nil {
		return nil, err
	}
	cert, err = FindCertificateByName(name, bs.db)
	if err != nil {
		return nil, err
	}
	return cert.ToAPI(), nil
}
