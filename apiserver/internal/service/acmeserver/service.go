package acmeserver

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	apiv1 "powerssl.dev/api/apiserver/v1"
	"powerssl.dev/apiserver/internal/model"
	"powerssl.dev/common/log"
	"powerssl.dev/common/telemetry"

	"powerssl.dev/apiserver/internal/repository"
)

var ServiceDesc = &apiv1.ACMEServerService_ServiceDesc

type Service struct {
	apiv1.UnimplementedACMEServerServiceServer
	logger    log.Logger
	queries   *repository.Queries
	telemeter *telemetry.Telemeter
}

func New(logger log.Logger, queries *repository.Queries, telemeter *telemetry.Telemeter) *Service {
	return &Service{
		logger:    logger,
		queries:   queries,
		telemeter: telemeter,
	}
}

func (s Service) Create(ctx context.Context, request *apiv1.CreateACMEServerRequest) (*apiv1.ACMEServer, error) {
	acmeServer, err := s.queries.CreateACMEServer(ctx, repository.CreateACMEServerParams{
		DisplayName:     request.GetAcmeServer().GetDisplayName(),
		DirectoryUrl:    request.GetAcmeServer().GetDirectoryUrl(),
		IntegrationName: request.GetAcmeServer().GetIntegrationName(),
	})
	if err != nil {
		return nil, err
	}
	return model.NewAcmeServer(acmeServer).Encode(), nil
}

func (s Service) Delete(ctx context.Context, request *apiv1.DeleteACMEServerRequest) (*emptypb.Empty, error) {
	queries, rollback, err := s.queries.NewTx(ctx)
	if err != nil {
		return nil, err
	}
	defer rollback(&err)
	id, err := model.ParseAcmeServerID(request.GetName())
	if err != nil {
		return nil, err
	}
	acmeServer, err := queries.GetACMEServer(ctx, id)
	if err != nil {
		return nil, err
	}
	if err = queries.DeleteACMEServer(ctx, acmeServer.ID); err != nil {
		return nil, err
	}
	if err = queries.Tx().Commit(ctx); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s Service) Get(ctx context.Context, request *apiv1.GetACMEServerRequest) (*apiv1.ACMEServer, error) {
	id, err := model.ParseAcmeServerID(request.GetName())
	if err != nil {
		return nil, err
	}
	acmeServer, err := s.queries.GetACMEServer(ctx, id)
	if err != nil {
		return nil, err
	}
	return model.NewAcmeServer(acmeServer).Encode(), nil
}

func (s Service) List(ctx context.Context, request *apiv1.ListACMEServersRequest) (*apiv1.ListACMEServersResponse, error) {
	var limit int32
	if request.GetPageSize() < 1 {
		limit = 10
	}
	acmeServers, err := s.queries.ListACMEServers(ctx, repository.ListACMEServersParams{
		SqlOrder:  "created_at",
		SqlOffset: 0,
		SqlLimit:  limit,
	})
	if err != nil {
		return nil, err
	}
	return &apiv1.ListACMEServersResponse{
		AcmeServers:   model.NewAcmeServers(acmeServers).Encode(),
		NextPageToken: "",
	}, nil
}

func (s Service) Update(ctx context.Context, request *apiv1.UpdateACMEServerRequest) (*apiv1.ACMEServer, error) {
	queries, rollback, err := s.queries.NewTx(ctx)
	if err != nil {
		return nil, err
	}
	defer rollback(&err)
	id, err := model.ParseAcmeServerID(request.GetName())
	if err != nil {
		return nil, err
	}
	updateACMEServerParams, err := model.AcmeServerUpdateParams(ctx, id, request.GetUpdateMask(), request.GetAcmeServer())
	if err != nil {
		return nil, err
	}
	acmeServer, err := queries.UpdateACMEServer(ctx, updateACMEServerParams)
	if err != nil {
		return nil, err
	}
	if err = queries.Tx().Commit(ctx); err != nil {
		return nil, err
	}
	return model.NewAcmeServer(acmeServer).Encode(), nil
}
