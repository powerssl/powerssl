package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-kit/kit/log"
	"github.com/pkg/errors"
	temporalclient "go.temporal.io/sdk/client"

	"powerssl.dev/apiserver/internal/model"
	"powerssl.dev/apiserver/internal/repository"
	"powerssl.dev/backend/temporal"
	"powerssl.dev/sdk/apiserver/acmeaccount"
	"powerssl.dev/sdk/apiserver/api"
	"powerssl.dev/workflow"
)

func New(repositories *repository.Repositories, logger log.Logger, temporalClient temporalclient.Client) acmeaccount.Service {
	var svc acmeaccount.Service
	{
		svc = NewBasicService(repositories, logger, temporalClient)
		svc = UnitOfWorkMiddleware(repositories, logger)(svc)
		svc = LoggingMiddleware(logger)(svc)
	}
	return svc
}

type basicService struct {
	*repository.Repositories
	logger   log.Logger
	temporal temporalclient.Client
}

func NewBasicService(repositories *repository.Repositories, logger log.Logger, temporalClient temporalclient.Client) acmeaccount.Service {
	return basicService{
		Repositories: repositories,
		logger:       logger,
		temporal:     temporalClient,
	}
}

func (s basicService) Create(ctx context.Context, parent string, apiACMEAccount *api.ACMEAccount) (_ *api.ACMEAccount, err error) {
	var acmeAccount *model.ACMEAccount
	if acmeAccount, err = model.NewACMEAccountFromAPI(parent, apiACMEAccount, ""); err != nil {
		return nil, err
	}
	if acmeAccount.ACMEServer, err = s.ACMEServers.FindByName(ctx, parent); err != nil {
		return nil, err
	}
	if err = s.ACMEAccounts.Add(ctx, acmeAccount); err != nil {
		return nil, err
	}
	_, err = s.temporal.ExecuteWorkflow(ctx, temporalclient.StartWorkflowOptions{
		ID:        fmt.Sprintf("%s/create-account", acmeAccount.Name()),
		TaskQueue: temporal.WorkerTaskQueue,
	}, workflow.CreateAccount, workflow.CreateAccountParams{
		Account:              acmeAccount.Name(),
		DirectoryURL:         acmeAccount.ACMEServer.DirectoryURL,
		TermsOfServiceAgreed: acmeAccount.TermsOfServiceAgreed,
		Contacts:             strings.Split(acmeAccount.Contacts, ","),
	})
	return acmeAccount.ToAPI(), nil
}

func (s basicService) Delete(ctx context.Context, name string) (err error) {
	var acmeAccount *model.ACMEAccount
	if acmeAccount, err = s.ACMEAccounts.FindByName(ctx, name); err != nil {
		return err
	}
	if err = s.ACMEAccounts.Remove(ctx, acmeAccount); err != nil {
		return err
	}
	return nil
}

func (s basicService) Get(ctx context.Context, name string) (_ *api.ACMEAccount, err error) {
	var acmeAccount *model.ACMEAccount
	if acmeAccount, err = s.ACMEAccounts.FindByName(ctx, name); err != nil {
		return nil, err
	}
	return acmeAccount.ToAPI(), nil
}

func (s basicService) List(ctx context.Context, parent string, pageSize int, pageToken string) (_ []*api.ACMEAccount, _ string, err error) {
	var acmeAccounts *model.ACMEAccounts
	if acmeAccounts, err = s.ACMEAccounts.GetAll(ctx); err != nil {
		return nil, "", errors.Wrap(err, "getting all acme accounts")
	}
	// TODO: paging
	_, nextPageToken := pageSize, pageToken
	return acmeAccounts.ToAPI(), nextPageToken, nil
}

func (s basicService) Update(ctx context.Context, name string, updateMask []string, apiACMEAccount *api.ACMEAccount) (_ *api.ACMEAccount, err error) {
	var acmeAccount *model.ACMEAccount
	if acmeAccount, err = s.ACMEAccounts.FindByName(ctx, name); err != nil {
		return nil, err
	}
	//var updates []string
	//for _, path := range updateMask {
	//	switch path {
	//	case "account_url", "contacts":
	//		updates = append(updates, path)
	//	default:
	//		// TODO: error
	//	}
	//}
	//var updatedACMEAccount *model.ACMEAccount
	//updatedACMEAccount, err = model.NewACMEAccountFromAPI(acmeAccount.ACMEServer.Name(), apiACMEAccount, acmeAccount.ID)
	//if err = unitofwork.GetUnit(ctx).Alter(updatedACMEAccount); err != nil {
	//	return nil, err
	//}
	return acmeAccount.ToAPI(), nil
}
