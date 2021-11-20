package acmeserver

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"

	apiv1 "powerssl.dev/api/apiserver/v1"
	"powerssl.dev/apiserver/internal/repository"
)

var ServiceDesc = &apiv1.ACMEServerService_ServiceDesc

type Service struct {
	apiv1.UnimplementedACMEServerServiceServer
	logger  *zap.SugaredLogger
	queries *repository.Queries
}

func New(logger *zap.SugaredLogger, queries *repository.Queries) *Service {
	return &Service{
		logger:  logger,
		queries: queries,
	}
}

func (s Service) Create(ctx context.Context, request *apiv1.CreateACMEServerRequest) (*apiv1.ACMEServer, error) {
	acmeServer, err := s.queries.CreateACMEServerFromAPI(ctx, request.GetAcmeServer())
	if err != nil {
		return nil, err
	}
	return acmeServer.ToAPI(), nil
}

func (s Service) Delete(ctx context.Context, request *apiv1.DeleteACMEServerRequest) (*emptypb.Empty, error) {
	queries, rollback, err := s.queries.NewTx(ctx)
	defer rollback(&err)
	n := strings.Split(request.GetName(), "/")
	id, err := uuid.Parse(n[1])
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
	n := strings.Split(request.GetName(), "/")
	id, err := uuid.Parse(n[1])
	if err != nil {
		return nil, err
	}
	acmeServer, err := s.queries.GetACMEServer(ctx, id)
	if err != nil {
		return nil, err
	}
	return acmeServer.ToAPI(), nil
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
		AcmeServers:   repository.AcmeServers(acmeServers).ToAPI(),
		NextPageToken: "",
	}, nil
}

func (s Service) Update(ctx context.Context, request *apiv1.UpdateACMEServerRequest) (*apiv1.ACMEServer, error) {
	queries, rollback, err := s.queries.NewTx(ctx)
	defer rollback(&err)
	n := strings.Split(request.GetName(), "/")
	id, err := uuid.Parse(n[1])
	if err != nil {
		return nil, err
	}
	request.GetUpdateMask()
	acmeServer, err := queries.UpdateACMEServerWithMask(ctx, id, request.GetUpdateMask(), request.GetAcmeServer())
	if err != nil {
		return nil, err
	}
	if err = queries.Tx().Commit(ctx); err != nil {
		return nil, err
	}
	return acmeServer.ToAPI(), nil
}
