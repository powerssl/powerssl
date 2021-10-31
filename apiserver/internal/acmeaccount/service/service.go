package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	temporalclient "go.temporal.io/sdk/client"

	"powerssl.dev/backend/temporal"
	"powerssl.dev/common/log"
	"powerssl.dev/sdk/apiserver/acmeaccount"
	"powerssl.dev/sdk/apiserver/api"
	"powerssl.dev/workflow"

	"powerssl.dev/apiserver/internal/repository"
)

func New(db *pgx.Conn, logger log.Logger, temporalClient temporalclient.Client) acmeaccount.Service {
	var svc acmeaccount.Service
	{
		svc = NewBasicService(db, logger, temporalClient)
		svc = LoggingMiddleware(logger)(svc)
	}
	return svc
}

type basicService struct {
	Queries  *repository.Queries
	db       *pgx.Conn
	logger   log.Logger
	temporal temporalclient.Client
}

func NewBasicService(db *pgx.Conn, logger log.Logger, temporalClient temporalclient.Client) acmeaccount.Service {
	return basicService{
		Queries:  repository.New(db),
		db:       db,
		logger:   logger,
		temporal: temporalClient,
	}
}

func (s basicService) Create(ctx context.Context, parent string, apiACMEAccount *api.ACMEAccount) (_ *api.ACMEAccount, err error) {
	acmeAccount, err := s.Queries.CreateACMEAccountFromAPI(ctx, parent, apiACMEAccount)
	if err != nil {
		return nil, err
	}
	acmeServer, err := s.Queries.GetACMEServer(ctx, acmeAccount.AcmeServerID)
	_, err = s.temporal.ExecuteWorkflow(ctx, temporalclient.StartWorkflowOptions{
		ID:        fmt.Sprintf("%s/create-account", acmeAccount.Name()),
		TaskQueue: temporal.WorkerTaskQueue,
	}, workflow.CreateAccount, workflow.CreateAccountParams{
		Account:              acmeAccount.Name(),
		DirectoryURL:         acmeServer.DirectoryUrl,
		TermsOfServiceAgreed: acmeAccount.TermsOfServiceAgreed,
		Contacts:             strings.Split(acmeAccount.Contacts, ","),
	})
	return acmeAccount.ToAPI(), nil
}

func (s basicService) Delete(ctx context.Context, name string) (err error) {
	queries, rollback, err := s.Queries.NewTx(ctx)
	defer rollback(&err)
	n := strings.Split(name, "/")
	id, err := uuid.Parse(n[3])
	if err != nil {
		return err
	}
	acmeAccount, err := queries.GetACMEAccount(ctx, id)
	if err != nil {
		return err
	}
	if err = queries.DeleteACMEAccount(ctx, acmeAccount.ID); err != nil {
		return err
	}
	return queries.Tx().Commit(ctx)
}

func (s basicService) Get(ctx context.Context, name string) (_ *api.ACMEAccount, err error) {
	n := strings.Split(name, "/")
	id, err := uuid.Parse(n[3])
	if err != nil {
		return nil, err
	}
	acmeAccount, err := s.Queries.GetACMEAccount(ctx, id)
	if err != nil {
		return nil, err
	}
	return acmeAccount.ToAPI(), nil
}

func (s basicService) List(ctx context.Context, parent string, pageSize int, pageToken string) (_ []*api.ACMEAccount, _ string, err error) {
	var nextPageToken string
	acmeAccounts, err := s.Queries.ListACMEAccounts(ctx, repository.ListACMEAccountsParams{
		SqlOrder:  "created_at",
		SqlOffset: 0,
		SqlLimit:  int32(pageSize),
	})
	if err != nil {
		return nil, "", err
	}
	return repository.AcmeAccounts(acmeAccounts).ToAPI(), nextPageToken, nil
}

func (s basicService) Update(ctx context.Context, name string, updateMask []string, apiACMEAccount *api.ACMEAccount) (_ *api.ACMEAccount, err error) {
	queries, rollback, err := s.Queries.NewTx(ctx)
	defer rollback(&err)
	n := strings.Split(name, "/")
	id, err := uuid.Parse(n[1])
	if err != nil {
		return nil, err
	}
	acmeAccount, err := queries.UpdateACMEAccountWithMask(ctx, id, updateMask, apiACMEAccount)
	if err != nil {
		return nil, err
	}
	if err = queries.Tx().Commit(ctx); err != nil {
		return nil, err
	}
	return acmeAccount.ToAPI(), nil
}
