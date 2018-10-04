package certificate_authority

import (
	"google.golang.org/grpc"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"

	"github.com/jinzhu/gorm"

	pb "powerssl.io/pkg/api/v1"
	model "powerssl.io/pkg/db"

	"powerssl.io/pkg/resources"
	endpoints "powerssl.io/pkg/resources/certificate_authority/endpoints"
	"powerssl.io/pkg/resources/certificate_authority/service"
	"powerssl.io/pkg/resources/certificate_authority/transport"
)

type CertificateAuthority struct {
	db     *gorm.DB
	logger log.Logger

	Model *model.CertificateAuthority

	endpoints endpoints.Endpoints
	service   *service.Service
}

func New(db *gorm.DB, logger log.Logger, duration metrics.Histogram) resources.Resource {
	svc := service.New()
	ep := endpoints.New(svc, logger, duration)

	return &CertificateAuthority{
		db:     db,
		logger: logger,

		Model: &model.CertificateAuthority{},

		endpoints: ep,
		service:   &svc,
	}
}

func (ca CertificateAuthority) RegisterGRPCServer(baseServer *grpc.Server) {
	pb.RegisterCertificateAuthorityServiceServer(baseServer, ca.grpcServer())
}

func (ca CertificateAuthority) grpcServer() pb.CertificateAuthorityServiceServer {
	return transport.NewGRPCServer(ca.endpoints, ca.logger)
}
