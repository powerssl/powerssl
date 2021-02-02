package certificate

import (
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/jmoiron/sqlx"
	stdopentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"

	"powerssl.dev/powerssl/internal/app/apiserver/certificate/endpoint"
	"powerssl.dev/powerssl/internal/app/apiserver/certificate/service"
	"powerssl.dev/powerssl/internal/app/apiserver/certificate/transport"
	apiv1 "powerssl.dev/powerssl/internal/pkg/apiserver/api/v1"
)

type Certificate struct {
	endpoints endpoint.Endpoints
	logger    log.Logger
	tracer    stdopentracing.Tracer
}

func New(db *sqlx.DB, logger log.Logger, tracer stdopentracing.Tracer, duration metrics.Histogram, auth kitendpoint.Middleware) *Certificate {
	svc := service.New(db, logger)
	endpoints := endpoint.NewEndpoints(svc, logger, tracer, duration, auth)

	return &Certificate{
		endpoints: endpoints,
		logger:    logger,
		tracer:    tracer,
	}
}

func (certificate *Certificate) RegisterGRPCServer(baseServer *grpc.Server) {
	grpcServer := transport.NewGRPCServer(certificate.endpoints, certificate.logger, certificate.tracer)
	apiv1.RegisterCertificateServiceServer(baseServer, grpcServer)
}

func (*Certificate) ServiceName() string {
	return "powerssl.apiserver.v1.CertificateService"
}
