package service

import (
	"context"
	"strconv"

	"github.com/go-kit/kit/log"
	"github.com/gogo/status"
	"github.com/jinzhu/gorm"
	otgorm "github.com/smacker/opentracing-gorm"
	"google.golang.org/grpc/codes"

	"powerssl.io/powerssl/internal/app/apiserver/certificate/model"
	"powerssl.io/powerssl/pkg/apiserver/api"
	"powerssl.io/powerssl/pkg/apiserver/certificate"
	controllerapi "powerssl.io/powerssl/pkg/controller/api"
	controllerclient "powerssl.io/powerssl/pkg/controller/client"
)

func New(db *gorm.DB, logger log.Logger, client *controllerclient.GRPCClient) certificate.Service {
	db.AutoMigrate(&model.Certificate{})
	var svc certificate.Service
	{
		svc = NewBasicService(db, logger, client)
		svc = LoggingMiddleware(logger)(svc)
	}
	return svc
}

type basicService struct {
	controllerclient *controllerclient.GRPCClient
	db               *gorm.DB
	logger           log.Logger
}

func NewBasicService(db *gorm.DB, logger log.Logger, client *controllerclient.GRPCClient) certificate.Service {
	return basicService{
		controllerclient: client,
		db:               db,
		logger:           logger,
	}
}

func (bs basicService) Create(ctx context.Context, certificate *api.Certificate) (*api.Certificate, error) {
	db := otgorm.SetSpanToGorm(ctx, bs.db)

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return nil, tx.Error
	}

	cert := model.NewCertificateFromAPI(certificate)
	if err := tx.Create(cert).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	workflow, err := bs.controllerclient.Workflow.Create(ctx, &controllerapi.Workflow{
		Kind: controllerapi.WorkflowKindRequestACMECertificate,
		IntegrationFilters: []*controllerapi.WorkflowIntegrationFilter{
			&controllerapi.WorkflowIntegrationFilter{},
		},
		Input: &controllerapi.RequestACMECertificateInput{
			DirectoryURL: "https://example.com/directory",
			AccountURL:   "https://example.com/acct/123",
			Dnsnames:     []string{"example.com", "example.net"},
			NotBefore:    "",
			NotAfter:     "",
		},
	})
	if err != nil {
		st := status.Convert(err)
		bs.logger.Log("err", "rpc", "code", st.Code(), "message", st.Message())
		return nil, status.Error(st.Code(), st.Message())
	}
	bs.logger.Log("workflow", workflow.Name)

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return cert.ToAPI(), nil
}

func (bs basicService) Delete(ctx context.Context, name string) error {
	db := otgorm.SetSpanToGorm(ctx, bs.db)

	certificate, err := model.FindCertificateByName(name, db)
	if err != nil {
		return err
	}
	return db.Delete(certificate).Error
}

func (bs basicService) Get(ctx context.Context, name string) (*api.Certificate, error) {
	db := otgorm.SetSpanToGorm(ctx, bs.db)

	certificate, err := model.FindCertificateByName(name, db)
	if err != nil {
		return nil, err
	}
	return certificate.ToAPI(), nil
}

func (bs basicService) List(ctx context.Context, pageSize int, pageToken string) ([]*api.Certificate, string, error) {
	db := otgorm.SetSpanToGorm(ctx, bs.db)

	var (
		certificates  model.Certificates
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
			return nil, "", status.Error(codes.InvalidArgument, "malformed page token")
		}
	}
	if err := db.Limit(pageSize + 1).Offset(offset).Find(&certificates).Error; err != nil {
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

func (bs basicService) Update(ctx context.Context, name string, certificate *api.Certificate) (*api.Certificate, error) {
	db := otgorm.SetSpanToGorm(ctx, bs.db)

	cert, err := model.FindCertificateByName(name, db)
	if err != nil {
		return nil, err
	}
	if err := db.Model(cert).Updates(model.NewCertificateFromAPI(certificate)).Error; err != nil {
		return nil, err
	}
	cert, err = model.FindCertificateByName(name, db)
	if err != nil {
		return nil, err
	}
	return cert.ToAPI(), nil
}
