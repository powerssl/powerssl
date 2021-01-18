package acmeaccount

import (
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/jinzhu/gorm"
	stdopentracing "github.com/opentracing/opentracing-go"
	temporalclient "go.temporal.io/sdk/client"
	"google.golang.org/grpc"

	"powerssl.dev/powerssl/internal/app/apiserver/acmeaccount/endpoint"
	"powerssl.dev/powerssl/internal/app/apiserver/acmeaccount/service"
	"powerssl.dev/powerssl/internal/app/apiserver/acmeaccount/transport"
	apiv1 "powerssl.dev/powerssl/internal/pkg/apiserver/api/v1"
	"powerssl.dev/powerssl/internal/pkg/vault"
	controllerclient "powerssl.dev/powerssl/pkg/controller/client"
)

type ACMEAccount struct {
	endpoints endpoint.Endpoints
	logger    log.Logger
	tracer    stdopentracing.Tracer
}

func New(db *gorm.DB, logger log.Logger, tracer stdopentracing.Tracer, duration metrics.Histogram, client *controllerclient.GRPCClient, temporalClient temporalclient.Client, vaultClient *vault.Client, auth kitendpoint.Middleware) *ACMEAccount {
	svc := service.New(db, logger, client, temporalClient, vaultClient)
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
