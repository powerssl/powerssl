package acmeaccount

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"go.temporal.io/sdk/client"
	"google.golang.org/protobuf/types/known/emptypb"

	apiv1 "powerssl.dev/api/apiserver/v1"
	"powerssl.dev/apiserver/internal/repository"
	"powerssl.dev/backend/temporal"
	"powerssl.dev/common/log"
	"powerssl.dev/workflow"
)

var ServiceDesc = &apiv1.ACMEAccountService_ServiceDesc

type Service struct {
	apiv1.UnimplementedACMEAccountServiceServer
	logger   log.Logger
	temporal client.Client
	queries  *repository.Queries
}

func New(logger log.Logger, temporal client.Client, queries *repository.Queries) *Service {
	return &Service{
		logger:   logger,
		temporal: temporal,
		queries:  queries,
	}
}

func (s Service) Create(ctx context.Context, request *apiv1.CreateACMEAccountRequest) (*apiv1.ACMEAccount, error) {
	acmeAccount, err := s.queries.CreateACMEAccountFromAPI(ctx, request.GetParent(), request.GetAcmeAccount())
	if err != nil {
		return nil, err
	}
	acmeServer, err := s.queries.GetACMEServer(ctx, acmeAccount.AcmeServerID)
	_, err = s.temporal.ExecuteWorkflow(ctx, client.StartWorkflowOptions{
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

func (s Service) Delete(ctx context.Context, request *apiv1.DeleteACMEAccountRequest) (*emptypb.Empty, error) {
	queries, rollback, err := s.queries.NewTx(ctx)
	defer rollback(&err)
	n := strings.Split(request.GetName(), "/")
	id, err := uuid.Parse(n[3])
	if err != nil {
		return nil, err
	}
	acmeAccount, err := queries.GetACMEAccount(ctx, id)
	if err != nil {
		return nil, err
	}
	if err = queries.DeleteACMEAccount(ctx, acmeAccount.ID); err != nil {
		return nil, err
	}
	if err = queries.Tx().Commit(ctx); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s Service) Get(ctx context.Context, request *apiv1.GetACMEAccountRequest) (*apiv1.ACMEAccount, error) {
	n := strings.Split(request.GetName(), "/")
	id, err := uuid.Parse(n[3])
	if err != nil {
		return nil, err
	}
	acmeAccount, err := s.queries.GetACMEAccount(ctx, id)
	if err != nil {
		return nil, err
	}
	return acmeAccount.ToAPI(), nil
}

func (s Service) List(ctx context.Context, request *apiv1.ListACMEAccountsRequest) (*apiv1.ListACMEAccountsResponse, error) {
	acmeAccounts, err := s.queries.ListACMEAccounts(ctx, repository.ListACMEAccountsParams{
		SqlOrder:  "created_at",
		SqlOffset: 0,
		SqlLimit:  request.GetPageSize(),
	})
	if err != nil {
		return nil, err
	}
	return &apiv1.ListACMEAccountsResponse{
		AcmeAccounts: repository.AcmeAccounts(acmeAccounts).ToAPI(),
	}, nil
}

func (s Service) Update(ctx context.Context, request *apiv1.UpdateACMEAccountRequest) (*apiv1.ACMEAccount, error) {
	queries, rollback, err := s.queries.NewTx(ctx)
	defer rollback(&err)
	n := strings.Split(request.GetName(), "/")
	id, err := uuid.Parse(n[1])
	if err != nil {
		return nil, err
	}
	acmeAccount, err := queries.UpdateACMEAccountWithMask(ctx, id, nil, request.GetAcmeAccount())
	if err != nil {
		return nil, err
	}
	if err = queries.Tx().Commit(ctx); err != nil {
		return nil, err
	}
	return acmeAccount.ToAPI(), nil
}
