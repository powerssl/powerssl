package certificate

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	apiv1 "powerssl.dev/api/apiserver/v1"
)

var ServiceDesc = &apiv1.CertificateService_ServiceDesc

type Service struct {
	apiv1.UnimplementedCertificateServiceServer
}

func New() *Service {
	return &Service{}
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
