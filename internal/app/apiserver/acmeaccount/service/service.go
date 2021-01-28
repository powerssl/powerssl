package service

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-kit/kit/log"
	"github.com/go-pg/pg/v10"
	"github.com/gogo/status"
	temporalclient "go.temporal.io/sdk/client"
	"google.golang.org/grpc/codes"

	"powerssl.dev/powerssl/internal/app/apiserver/acmeaccount/model"
	acmeservermodel "powerssl.dev/powerssl/internal/app/apiserver/acmeserver/model"
	"powerssl.dev/powerssl/internal/pkg/temporal"
	"powerssl.dev/powerssl/internal/pkg/vault"
	"powerssl.dev/powerssl/pkg/apiserver/acmeaccount"
	"powerssl.dev/powerssl/pkg/apiserver/api"
	"powerssl.dev/powerssl/internal/pkg/temporal/workflow"
)

func New(db *pg.DB, logger log.Logger, temporalClient temporalclient.Client, vaultClient *vault.Client) acmeaccount.Service {
	var svc acmeaccount.Service
	{
		svc = NewBasicService(db, logger, temporalClient, vaultClient)
		svc = LoggingMiddleware(logger)(svc)
	}
	return svc
}

type basicService struct {
	temporalClient temporalclient.Client
	db             *pg.DB
	logger         log.Logger
	vaultClient    *vault.Client
}

func NewBasicService(db *pg.DB, logger log.Logger, temporalClient temporalclient.Client, vaultClient *vault.Client) acmeaccount.Service {
	return basicService{
		temporalClient: temporalClient,
		vaultClient:    vaultClient,
		db:             db,
		logger:         logger,
	}
}

func (bs basicService) Create(ctx context.Context, parent string, apiacmeaccount *api.ACMEAccount) (*api.ACMEAccount, error) {
	acmeServer, err := acmeservermodel.FindACMEServerByName(parent, bs.db)
	if err != nil {
		return nil, err
	}
	acmeAccount, err := model.NewACMEAccountFromAPI(parent, apiacmeaccount)
	if err != nil {
		return nil, err
	}
	if err := bs.db.RunInTransaction(ctx, func(tx *pg.Tx) error {
		_, err := tx.Model(acmeAccount).Insert()
		if err != nil {
			return err
		}
		if err := bs.vaultClient.CreateTransitKey(ctx, acmeAccount.ID); err != nil {
			return err
		}
		_, err = bs.temporalClient.ExecuteWorkflow(ctx, temporalclient.StartWorkflowOptions{
			ID:        fmt.Sprintf("%s/create-account", acmeAccount.Name()),
			TaskQueue: temporal.TaskQueue,
		}, workflow.CreateAccount, workflow.CreateAccountParams{
			Account:              acmeAccount.Name(),
			DirectoryURL:         acmeServer.DirectoryURL,
			TermsOfServiceAgreed: acmeAccount.TermsOfServiceAgreed,
			Contacts:             strings.Split(acmeAccount.Contacts, ","),
		})
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return acmeAccount.ToAPI(), nil
}

func (bs basicService) Delete(ctx context.Context, name string) error {
	acmeAccount, err := model.FindACMEAccountByName(name, bs.db)
	if err != nil {
		return err
	}
	_, err = bs.db.Model(acmeAccount).Delete()
	return err
}

func (bs basicService) Get(ctx context.Context, name string) (*api.ACMEAccount, error) {
	acmeAccount, err := model.FindACMEAccountByName(name, bs.db)
	if err != nil {
		return nil, err
	}
	return acmeAccount.ToAPI(), nil
}

func (bs basicService) List(ctx context.Context, parent string, pageSize int, pageToken string) ([]*api.ACMEAccount, string, error) {
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
	if err := bs.db.Model(&acmeAccounts).Limit(pageSize + 1).Offset(offset).Select(); err != nil {
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
	dbACMEAccount, err := model.FindACMEAccountByName(name, bs.db)
	if err != nil {
		return nil, err
	}
	var updates []string
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
	if _, err := bs.db.Model(updateACMEAccount).Column(updates...).WherePK().Update(); err != nil {
		return nil, err
	}
	return dbACMEAccount.ToAPI(), nil
}
