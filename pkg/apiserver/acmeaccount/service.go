package acmeaccount

import (
	"context"
	"strings"

	"github.com/go-kit/kit/log"
	"github.com/gogo/status"
	"github.com/jinzhu/gorm"
	otgorm "github.com/smacker/opentracing-gorm"
	"google.golang.org/grpc/codes"

	"powerssl.io/pkg/apiserver/api"
	controllerapi "powerssl.io/pkg/controller/api"
	controllerclient "powerssl.io/pkg/controller/client"
)

var ErrUnimplemented = status.Error(codes.Unimplemented, "Coming soon")

type Service interface {
	Create(ctx context.Context, parent string, acmeAccount *api.ACMEAccount) (*api.ACMEAccount, error)
	Delete(ctx context.Context, name string) error
	Get(ctx context.Context, name string) (*api.ACMEAccount, error)
	List(ctx context.Context, parent string, pageSize int, pageToken string) ([]*api.ACMEAccount, string, error)
	Update(ctx context.Context, name string, acmeAccount *api.ACMEAccount) (*api.ACMEAccount, error)
}

func New(db *gorm.DB, logger log.Logger, client *controllerclient.GRPCClient) Service {
	db.AutoMigrate(&ACMEAccount{})
	var svc Service
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

func NewBasicService(db *gorm.DB, logger log.Logger, client *controllerclient.GRPCClient) Service {
	return basicService{
		controllerclient: client,
		db:               db,
		logger:           logger,
	}
}

func (bs basicService) Create(ctx context.Context, parent string, acmeAccount *api.ACMEAccount) (*api.ACMEAccount, error) {
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

	account, err := NewACMEAccountFromAPI(parent, acmeAccount)
	if err != nil {
		return nil, err
	}
	if err := tx.Create(account).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	acmeServer, err := account.ACMEServer(db, "directory_url, integration_name")
	if err != nil {
		return nil, status.Error(codes.NotFound, "not found")
	}

	workflow, err := bs.controllerclient.Workflow.Create(ctx, &controllerapi.Workflow{
		Kind: controllerapi.WorkflowKindCreateACMEAccount,
		IntegrationFilters: []*controllerapi.WorkflowIntegrationFilter{
			{
				Kind: controllerapi.IntegrationKindACME,
				Name: acmeServer.IntegrationName,
			},
		},
		Input: &controllerapi.CreateACMEAccountInput{
			DirectoryURL:         acmeServer.DirectoryURL,
			TermsOfServiceAgreed: account.TermsOfServiceAgreed,
			Contacts:             strings.Split(account.Contacts, ","),
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

	return account.ToAPI(), nil
}

func (bs basicService) Delete(ctx context.Context, name string) error {
	db := otgorm.SetSpanToGorm(ctx, bs.db)

	acmeAccount, err := FindACMEAccountByName(name, db)
	if err != nil {
		return err
	}
	return db.Delete(acmeAccount).Error
}

func (bs basicService) Get(ctx context.Context, name string) (*api.ACMEAccount, error) {
	db := otgorm.SetSpanToGorm(ctx, bs.db)

	acmeAccount, err := FindACMEAccountByName(name, db)
	if err != nil {
		return nil, err
	}
	return acmeAccount.ToAPI(), nil
}

func (bs basicService) List(ctx context.Context, parent string, pageSize int, pageToken string) ([]*api.ACMEAccount, string, error) {
	return nil, "", ErrUnimplemented
}

func (bs basicService) Update(ctx context.Context, name string, acmeAccount *api.ACMEAccount) (*api.ACMEAccount, error) {
	return nil, ErrUnimplemented
}
