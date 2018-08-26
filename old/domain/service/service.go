package domainservice

import (
	"context"
	"strings"

	"github.com/go-kit/kit/log"

	"github.com/jinzhu/gorm"

	"powerssl.io/pkg/domain/model"
)

type Service interface {
	Create(ctx context.Context, domain domainmodel.Domain) (domainmodel.Domain, error)
	Delete(ctx context.Context, name string) error
	Get(ctx context.Context, name string) (domainmodel.Domain, error)
	List(ctx context.Context, pageSize int, pageToken string) ([]domainmodel.Domain, error)
	Update(ctx context.Context, domain domainmodel.Domain, updateMask string) (domainmodel.Domain, error) // TODO: Update Mask type
}

func New(db *gorm.DB, logger log.Logger) Service {
	var svc Service
	{
		svc = NewBasicService(db)
		svc = LoggingMiddleware(logger)(svc)
	}
	return svc
}

func NewBasicService(db *gorm.DB) Service {
	return basicService{db}
}

type basicService struct {
	db *gorm.DB
}

func (s basicService) Create(_ context.Context, domain domainmodel.Domain) (domainmodel.Domain, error) {
	s.db.Create(&domain)

	return domain, nil
}

func (s basicService) Delete(_ context.Context, name string) error {
	var domain domainmodel.Domain
	dnsName := strings.TrimPrefix(name, "domains/")
	s.db.Where("dns_name = ?", dnsName).First(&domain)
	s.db.Delete(&domain)

	return nil
}

func (s basicService) Get(_ context.Context, name string) (domainmodel.Domain, error) {
	var domain domainmodel.Domain
	dnsName := strings.TrimPrefix(name, "domains/")
	s.db.Where("dns_name = ?", dnsName).First(&domain)
	return domain, nil
}

func (s basicService) List(_ context.Context, pageSize int, pageToken string) ([]domainmodel.Domain, error) {
	var domains []domainmodel.Domain
	if err := s.db.Find(&domains).Error; err != nil {
		return []domainmodel.Domain{}, err
	}

	return domains, nil
}

func (s basicService) Update(_ context.Context, domain domainmodel.Domain, updateMask string) (domainmodel.Domain, error) { // TODO: Update Mask type
	var d domainmodel.Domain
	dnsName := strings.TrimPrefix(domain.Name, "domains/")
	s.db.Where("dns_name = ?", dnsName).First(&d)
	d = domain
	s.db.Save(&d)
	return d, nil
}
