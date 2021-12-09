package acmeaccount

import (
	"context"
	"fmt"
	"strings"

	"go.temporal.io/sdk/client"
	"google.golang.org/protobuf/types/known/emptypb"

	apiv1 "powerssl.dev/api/apiserver/v1"
	"powerssl.dev/apiserver/internal/model"
	"powerssl.dev/backend/temporal"
	"powerssl.dev/common/log"
	"powerssl.dev/common/telemetry"
	"powerssl.dev/workflow"

	"powerssl.dev/apiserver/internal/repository"
)

var ServiceDesc = &apiv1.ACMEAccountService_ServiceDesc

type Service struct {
	apiv1.UnimplementedACMEAccountServiceServer
	logger    log.Logger
	queries   *repository.Queries
	telemeter *telemetry.Telemeter
	temporal  client.Client
}

func New(logger log.Logger, queries *repository.Queries, telemeter *telemetry.Telemeter, temporal client.Client) *Service {
	return &Service{
		logger:    logger,
		queries:   queries,
		telemeter: telemeter,
		temporal:  temporal,
	}
}

func (s Service) Create(ctx context.Context, request *apiv1.CreateACMEAccountRequest) (*apiv1.ACMEAccount, error) {
	acmeServerID, err := model.ParseAcmeServerID(request.GetParent())
	if err != nil {
		return nil, err
	}
	acmeAccount, err := s.queries.CreateACMEAccount(ctx, repository.CreateACMEAccountParams{
		AcmeServerID:         acmeServerID,
		DisplayName:          request.GetAcmeAccount().GetDisplayName(),
		Title:                request.GetAcmeAccount().GetTitle(),
		Description:          request.GetAcmeAccount().GetDescription(),
		TermsOfServiceAgreed: request.GetAcmeAccount().GetTermsOfServiceAgreed(),
		Contacts:             strings.Join(request.GetAcmeAccount().GetContacts(), ","),
	})
	if err != nil {
		return nil, err
	}
	acmeAccountModel := model.NewAcmeAccount(acmeAccount)
	acmeServer, err := s.queries.GetACMEServer(ctx, acmeAccountModel.AcmeServerID)
	if err != nil {
		return nil, err
	}
	if _, err = s.temporal.ExecuteWorkflow(ctx, client.StartWorkflowOptions{
		ID:        fmt.Sprintf("%s/create-account", acmeAccountModel.Name()),
		TaskQueue: temporal.WorkerTaskQueue,
	}, workflow.CreateAccount, workflow.CreateAccountParams{
		Account:              acmeAccountModel.Name(),
		DirectoryURL:         acmeServer.DirectoryUrl,
		TermsOfServiceAgreed: acmeAccountModel.TermsOfServiceAgreed,
		Contacts:             strings.Split(acmeAccountModel.Contacts, ","),
	}); err != nil {
		return nil, err
	}
	return acmeAccountModel.Encode(), nil
}

func (s Service) Delete(ctx context.Context, request *apiv1.DeleteACMEAccountRequest) (*emptypb.Empty, error) {
	queries, rollback, err := s.queries.NewTx(ctx)
	if err != nil {
		return nil, err
	}
	defer rollback(&err)
	id, err := model.ParseAcmeAccountID(request.GetName())
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
	id, err := model.ParseAcmeAccountID(request.GetName())
	if err != nil {
		return nil, err
	}
	acmeAccount, err := s.queries.GetACMEAccount(ctx, id)
	if err != nil {
		return nil, err
	}
	return model.NewAcmeAccount(acmeAccount).Encode(), nil
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
		AcmeAccounts:  model.NewAcmeAccounts(acmeAccounts).Encode(),
		NextPageToken: "",
	}, nil
}

func (s Service) Update(ctx context.Context, request *apiv1.UpdateACMEAccountRequest) (*apiv1.ACMEAccount, error) {
	queries, rollback, err := s.queries.NewTx(ctx)
	if err != nil {
		return nil, err
	}
	defer rollback(&err)
	id, err := model.ParseAcmeAccountID(request.GetName())
	if err != nil {
		return nil, err
	}
	updateACMEAccountParams, err := model.AcmeAccountUpdateParams(ctx, id, request.GetUpdateMask(), request.GetAcmeAccount())
	if err != nil {
		return nil, err
	}
	acmeAccount, err := queries.UpdateACMEAccount(ctx, updateACMEAccountParams)
	if err != nil {
		return nil, err
	}
	if err = queries.Tx().Commit(ctx); err != nil {
		return nil, err
	}
	return model.NewAcmeAccount(acmeAccount).Encode(), nil
}
