package acmeserver

import (
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/go-pg/pg/v10"
	stdopentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"

	"powerssl.dev/powerssl/internal/app/apiserver/acmeserver/endpoint"
	"powerssl.dev/powerssl/internal/app/apiserver/acmeserver/service"
	"powerssl.dev/powerssl/internal/app/apiserver/acmeserver/transport"
	apiv1 "powerssl.dev/powerssl/internal/pkg/apiserver/api/v1"
)

type ACMEServer struct {
	endpoints endpoint.Endpoints
	logger    log.Logger
	tracer    stdopentracing.Tracer
}

func New(db *pg.DB, logger log.Logger, tracer stdopentracing.Tracer, duration metrics.Histogram, auth kitendpoint.Middleware) *ACMEServer {
	svc := service.New(db, logger)
	endpoints := endpoint.NewEndpoints(svc, logger, tracer, duration, auth)

	return &ACMEServer{
		endpoints: endpoints,
		logger:    logger,
		tracer:    tracer,
	}
}

func (acmeServer *ACMEServer) RegisterGRPCServer(baseServer *grpc.Server) {
	grpcServer := transport.NewGRPCServer(acmeServer.endpoints, acmeServer.logger, acmeServer.tracer)
	apiv1.RegisterACMEServerServiceServer(baseServer, grpcServer)
}

func (*ACMEServer) ServiceName() string {
	return "powerssl.apiserver.v1.ACMEServerService"
}
