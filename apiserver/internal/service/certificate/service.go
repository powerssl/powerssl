package certificate

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	apiv1 "powerssl.dev/api/apiserver/v1"
	"powerssl.dev/common/log"
	"powerssl.dev/common/telemetry"

	"powerssl.dev/apiserver/internal/repository"
)

var ServiceDesc = &apiv1.CertificateService_ServiceDesc

type Service struct {
	apiv1.UnimplementedCertificateServiceServer
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

func (s Service) Create(ctx context.Context, request *apiv1.CreateCertificateRequest) (*apiv1.Certificate, error) {
	panic("implement me")
}

func (s Service) Delete(ctx context.Context, request *apiv1.DeleteCertificateRequest) (*emptypb.Empty, error) {
	panic("implement me")
}

func (s Service) Get(ctx context.Context, request *apiv1.GetCertificateRequest) (*apiv1.Certificate, error) {
	panic("implement me")
}

func (s Service) List(ctx context.Context, request *apiv1.ListCertificatesRequest) (*apiv1.ListCertificatesResponse, error) {
	panic("implement me")
}

func (s Service) Update(ctx context.Context, request *apiv1.UpdateCertificateRequest) (*apiv1.Certificate, error) {
	panic("implement me")
}
