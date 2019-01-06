package acmeserver // import "powerssl.io/pkg/apiserver/acmeserver"

import (
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/jinzhu/gorm"
	stdopentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"

	"powerssl.io/pkg/apiserver/acmeserver/endpoint"
	"powerssl.io/pkg/apiserver/acmeserver/service"
	"powerssl.io/pkg/apiserver/acmeserver/transport"
	apiv1 "powerssl.io/pkg/apiserver/api/v1"
	controllerclient "powerssl.io/pkg/controller/client"
)

type ACMEServer struct {
	endpoints endpoint.Endpoints
	logger    log.Logger
	tracer    stdopentracing.Tracer
}

func New(db *gorm.DB, logger log.Logger, tracer stdopentracing.Tracer, duration metrics.Histogram, client *controllerclient.GRPCClient, auth kitendpoint.Middleware) *ACMEServer {
	svc := service.New(db, logger, client)
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
