package certificate_authority

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"

	apiv1 "powerssl.io/pkg/api/v1"
	"powerssl.io/pkg/resources"
	"powerssl.io/pkg/resources/certificate_authority/endpoints"
	"powerssl.io/pkg/resources/certificate_authority/service"
	"powerssl.io/pkg/resources/certificate_authority/transport"
)

type certificateAuthority struct {
	db        *gorm.DB
	endpoints endpoints.Endpoints
	logger    log.Logger
}

func New(db *gorm.DB, logger log.Logger, duration metrics.Histogram) resources.Resource {
	svc := service.New()
	ep := endpoints.New(svc, logger, duration)

	return &certificateAuthority{
		db:        db,
		endpoints: ep,
		logger:    logger,
	}
}

func (ca *certificateAuthority) RegisterGRPCServer(baseServer *grpc.Server) {
	grpcServer := transport.NewGRPCServer(ca.endpoints, ca.logger)
	apiv1.RegisterCertificateAuthorityServiceServer(baseServer, grpcServer)
}
