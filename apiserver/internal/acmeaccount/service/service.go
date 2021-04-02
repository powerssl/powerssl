package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-kit/kit/log"
	"github.com/pkg/errors"
	temporalclient "go.temporal.io/sdk/client"

	"powerssl.dev/backend/temporal"
	"powerssl.dev/sdk/apiserver/acmeaccount"
	"powerssl.dev/sdk/apiserver/api"
	"powerssl.dev/workflow"

	"powerssl.dev/apiserver/internal/model"
	"powerssl.dev/apiserver/internal/repository"
)

func New(repositories *repository.Repositories, logger log.Logger, temporalClient temporalclient.Client) acmeaccount.Service {
	var svc acmeaccount.Service
	{
		svc = NewBasicService(repositories, logger, temporalClient)
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
	if err = s.Transaction(ctx, func(ctx context.Context) error {
		if acmeAccount.ACMEServer, err = s.ACMEServers.FindOneByName(ctx, parent); err != nil {
			return err
		}
		return s.ACMEAccounts.Insert(ctx, acmeAccount)
	}); err != nil {
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
	return s.Transaction(ctx, func(ctx context.Context) error {
		if acmeAccount, err = s.ACMEAccounts.FindOneByName(ctx, name); err != nil {
			return err
		}
		return s.ACMEAccounts.Delete(ctx, acmeAccount)
	})
}

func (s basicService) Get(ctx context.Context, name string) (_ *api.ACMEAccount, err error) {
	var acmeAccount *model.ACMEAccount
	if acmeAccount, err = s.ACMEAccounts.FindOneByName(ctx, name); err != nil {
		return nil, err
	}
	return acmeAccount.ToAPI(), nil
}

func (s basicService) List(ctx context.Context, parent string, pageSize int, pageToken string) (_ []*api.ACMEAccount, _ string, err error) {
	var acmeAccounts model.ACMEAccounts
	if acmeAccounts, err = s.ACMEAccounts.FindAllByParent(ctx, parent); err != nil {
		return nil, "", errors.Wrap(err, "getting all acme accounts")
	}
	// TODO: paging
	_, nextPageToken := pageSize, pageToken
	return acmeAccounts.ToAPI(), nextPageToken, nil
}

func (s basicService) Update(ctx context.Context, name string, updateMask []string, apiACMEAccount *api.ACMEAccount) (_ *api.ACMEAccount, err error) {
	var acmeAccount *model.ACMEAccount
	if err = s.Transaction(ctx, func(ctx context.Context) error {
		if acmeAccount, err = s.ACMEAccounts.FindOneByName(ctx, name); err != nil {
			return err
		}
		var clauses map[string]interface{}
		if clauses, err = acmeAccount.UpdateWithMask(ctx, updateMask, apiACMEAccount); err != nil {
			return err
		}
		return s.ACMEAccounts.Update(ctx, acmeAccount, clauses)
	}); err != nil {
		return nil, err
	}
	return acmeAccount.ToAPI(), nil
}
