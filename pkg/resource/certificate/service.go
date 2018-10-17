package certificate

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"

	"powerssl.io/pkg/api"
	controllerclient "powerssl.io/pkg/controller/client"
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
		svc = NewBasicService(db, logger)
		svc = LoggingMiddleware(logger)(svc)
	}
	return svc
}

type basicService struct {
	controllerclient *controllerclient.GRPCClient
	db               *gorm.DB
	logger           log.Logger
}

func NewBasicService(db *gorm.DB, logger log.Logger) Service {
	client, err := controllerclient.NewGRPCClient("localhost:8081", "", "", true, true, logger)
	if err != nil {
		panic(err)
	}

	return basicService{
		controllerclient: client,
		db:               db,
		logger:           logger,
	}
}

func (bs basicService) Create(_ context.Context, certificate *api.Certificate) (*api.Certificate, error) {
	tx := bs.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return nil, tx.Error
	}

	cert := NewCertificateFromAPI(certificate)
	if err := tx.Create(cert).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	workflow, err := bs.controllerclient.Workflow.Create(context.Background(), fmt.Sprint("certificates/", cert.ID))
	if err != nil {
		return nil, err
	}
	bs.logger.Log("workflow", workflow.Name)

	if err := tx.Commit().Error; err != nil {
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
	var (
		certificates  Certificates
		nextPageToken string
	)
	offset := -1
	if pageSize < 1 {
		pageSize = 10
	} else if pageSize > 20 {
		pageSize = 20
	}
	if pageToken != "" {
		var err error
		offset, err = strconv.Atoi(pageToken)
		if err != nil {
			return nil, "", fmt.Errorf("Invalid page token")
		}
	}
	if err := bs.db.Limit(pageSize + 1).Offset(offset).Find(&certificates).Error; err != nil {
		return nil, "", err
	}
	if len(certificates) > pageSize {
		certificates = certificates[:len(certificates)-1]
		if offset == -1 {
			nextPageToken = strconv.Itoa(pageSize)
		} else {
			nextPageToken = strconv.Itoa(offset + pageSize)
		}
	}
	return certificates.ToAPI(), nextPageToken, nil
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
