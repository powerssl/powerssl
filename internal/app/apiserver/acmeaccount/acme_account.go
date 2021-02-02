package acmeaccount

import (
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	stdopentracing "github.com/opentracing/opentracing-go"
	temporalclient "go.temporal.io/sdk/client"
	"google.golang.org/grpc"

	"powerssl.dev/powerssl/internal/app/apiserver/acmeaccount/endpoint"
	"powerssl.dev/powerssl/internal/app/apiserver/acmeaccount/service"
	"powerssl.dev/powerssl/internal/app/apiserver/acmeaccount/transport"
	"powerssl.dev/powerssl/internal/app/apiserver/repository"
	apiv1 "powerssl.dev/powerssl/internal/pkg/apiserver/api/v1"
	"powerssl.dev/powerssl/internal/pkg/vault"
)

type ACMEAccount struct {
	endpoints endpoint.Endpoints
	logger    log.Logger
	tracer    stdopentracing.Tracer
}

func New(repositories *repository.Repositories, logger log.Logger, tracer stdopentracing.Tracer, duration metrics.Histogram, temporalClient temporalclient.Client, vaultClient *vault.Client, auth kitendpoint.Middleware) *ACMEAccount {
	svc := service.New(repositories, logger, temporalClient, vaultClient)
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
