package service

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-kit/kit/log"
	"github.com/gogo/status"
	"github.com/jinzhu/gorm"
	otgorm "github.com/smacker/opentracing-gorm"
	temporalclient "go.temporal.io/sdk/client"
	"google.golang.org/grpc/codes"

	"powerssl.dev/powerssl/internal/app/apiserver/acmeaccount/model"
	"powerssl.dev/powerssl/internal/pkg/temporal"
	"powerssl.dev/powerssl/internal/pkg/vault"
	"powerssl.dev/powerssl/pkg/apiserver/acmeaccount"
	"powerssl.dev/powerssl/pkg/apiserver/api"
	controllerapi "powerssl.dev/powerssl/pkg/controller/api"
	controllerclient "powerssl.dev/powerssl/pkg/controller/client"
)

func New(db *gorm.DB, logger log.Logger, client *controllerclient.GRPCClient, temporalClient temporalclient.Client, vaultClient *vault.Client) acmeaccount.Service {
	db.AutoMigrate(&model.ACMEAccount{})
	var svc acmeaccount.Service
	{
		svc = NewBasicService(db, logger, client, temporalClient, vaultClient)
		svc = LoggingMiddleware(logger)(svc)
	}
	return svc
}

type basicService struct {
	controllerClient *controllerclient.GRPCClient
	temporalClient   temporalclient.Client
	db               *gorm.DB
	logger           log.Logger
	vaultClient      *vault.Client
}

func NewBasicService(db *gorm.DB, logger log.Logger, client *controllerclient.GRPCClient, temporalClient temporalclient.Client, vaultClient *vault.Client) acmeaccount.Service {
	return basicService{
		controllerClient: client,
		temporalClient:   temporalClient,
		vaultClient:      vaultClient,
		db:               db,
		logger:           logger,
	}
}

func (bs basicService) Create(ctx context.Context, parent string, acmeAccount *api.ACMEAccount) (*api.ACMEAccount, error) {
	db := otgorm.SetSpanToGorm(ctx, bs.db)

	account, err := model.NewACMEAccountFromAPI(parent, acmeAccount)
	if err != nil {
		return nil, err
	}
	if err := db.Create(account).Error; err != nil {
		return nil, err
	}

	acmeServer, err := account.ACMEServer(db, "directory_url, integration_name")
	if err != nil { // TODO
		return nil, status.Error(codes.NotFound, "not found")
	}

	if err := bs.vaultClient.CreateTransitKey(ctx, account.ID); err != nil {
		return nil, err
	}


	workflow, err := bs.controllerClient.Workflow.Create(ctx, &controllerapi.Workflow{
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

func (bs basicService) Update(ctx context.Context, name string, updateMask []string, acmeAccount *api.ACMEAccount) (*api.ACMEAccount, error) {
	db := otgorm.SetSpanToGorm(ctx, bs.db)

	dbACMEAccount, err := model.FindACMEAccountByName(name, db)
	if err != nil {
		return nil, err
	}
	updates := []interface{}{nil}
	for _, path := range updateMask {
		switch path {
		case "account_url", "contacts":
			updates = append(updates, path)
		default:
			// TODO: error
		}
	}
	updateACMEAccount, err := model.NewACMEAccountFromAPI(fmt.Sprintf("acmeServers/%s", dbACMEAccount.ACMEServerID), acmeAccount)
	if err != nil {
		return nil, err
	}
	// TODO: Compare name with acmeAccount.Name (empty or match)
	if err := db.Model(dbACMEAccount).Select(nil, updates...).Updates(updateACMEAccount).Error; err != nil {
		return nil, err
	}
	return dbACMEAccount.ToAPI(), nil
}
