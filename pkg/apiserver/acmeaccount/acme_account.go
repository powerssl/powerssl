package acmeaccount // import "powerssl.io/pkg/apiserver/acmeaccount"

import (
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/jinzhu/gorm"
	stdopentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"

	"powerssl.io/internal/pkg/util/vault"
	"powerssl.io/pkg/apiserver/acmeaccount/endpoint"
	"powerssl.io/pkg/apiserver/acmeaccount/service"
	"powerssl.io/pkg/apiserver/acmeaccount/transport"
	apiv1 "powerssl.io/pkg/apiserver/api/v1"
	controllerclient "powerssl.io/pkg/controller/client"
)

type ACMEAccount struct {
	endpoints endpoint.Endpoints
	logger    log.Logger
	tracer    stdopentracing.Tracer
}

func New(db *gorm.DB, logger log.Logger, tracer stdopentracing.Tracer, duration metrics.Histogram, client *controllerclient.GRPCClient, vaultClient *vault.Client, auth kitendpoint.Middleware) *ACMEAccount {
	svc := service.New(db, logger, client, vaultClient)
	endpoints := endpoint.NewEndpoints(svc, logger, tracer, duration, auth)

	return &ACMEAccount{
		endpoints: endpoints,
		logger:    logger,
		tracer:    tracer,
	}
}

func (acmeAccount *ACMEAccount) RegisterGRPCServer(baseServer *grpc.Server) {
	grpcServer := transport.NewGRPCServer(acmeAccount.endpoints, acmeAccount.logger, acmeAccount.tracer)
	apiv1.RegisterACMEAccountServiceServer(baseServer, grpcServer)
}

func (*ACMEAccount) ServiceName() string {
	return "powerssl.apiserver.v1.ACMEAccountService"
}
