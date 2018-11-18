package service

import (
	"context"
	"strconv"
	"strings"

	"github.com/go-kit/kit/log"
	"github.com/gogo/status"
	"github.com/jinzhu/gorm"
	otgorm "github.com/smacker/opentracing-gorm"
	"google.golang.org/grpc/codes"

	"powerssl.io/pkg/apiserver/acmeaccount/meta"
	"powerssl.io/pkg/apiserver/acmeaccount/model"
	"powerssl.io/pkg/apiserver/api"
	controllerapi "powerssl.io/pkg/controller/api"
	controllerclient "powerssl.io/pkg/controller/client"
)

var ErrUnimplemented = status.Error(codes.Unimplemented, "Coming soon")

func New(db *gorm.DB, logger log.Logger, client *controllerclient.GRPCClient) meta.Service {
	db.AutoMigrate(&model.ACMEAccount{})
	var svc meta.Service
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

func NewBasicService(db *gorm.DB, logger log.Logger, client *controllerclient.GRPCClient) meta.Service {
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

	account, err := model.NewACMEAccountFromAPI(parent, acmeAccount)
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
			Account:              account.Name(),
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

	acmeAccount, err := model.FindACMEAccountByName(name, db)
	if err != nil {
		return err
	}
	return db.Delete(acmeAccount).Error
}

func (bs basicService) Get(ctx context.Context, name string) (*api.ACMEAccount, error) {
	db := otgorm.SetSpanToGorm(ctx, bs.db)

	acmeAccount, err := model.FindACMEAccountByName(name, db)
	if err != nil {
		return nil, err
	}
	return acmeAccount.ToAPI(), nil
}

func (bs basicService) List(ctx context.Context, parent string, pageSize int, pageToken string) ([]*api.ACMEAccount, string, error) {
	db := otgorm.SetSpanToGorm(ctx, bs.db)

	if pageSize < 1 {
		pageSize = 10
	} else if pageSize > 20 {
		pageSize = 20
	}
	offset := -1
	if pageToken != "" {
		var err error
		offset, err = strconv.Atoi(pageToken)
		if err != nil {
			return nil, "", status.Error(codes.InvalidArgument, "malformed page token")
		}
	}
	var acmeAccounts model.ACMEAccounts
	if err := db.Limit(pageSize + 1).Offset(offset).Find(&acmeAccounts).Error; err != nil {
		return nil, "", err
	}
	var nextPageToken string
	if len(acmeAccounts) > pageSize {
		acmeAccounts = acmeAccounts[:len(acmeAccounts)-1]
		if offset == -1 {
			nextPageToken = strconv.Itoa(pageSize)
		} else {
			nextPageToken = strconv.Itoa(offset + pageSize)
		}
	}
	return acmeAccounts.ToAPI(), nextPageToken, nil
}

func (bs basicService) Update(ctx context.Context, name string, acmeAccount *api.ACMEAccount) (*api.ACMEAccount, error) {
	return nil, ErrUnimplemented
}
