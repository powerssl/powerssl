package service

import (
	"github.com/google/wire"

	apiv1 "powerssl.dev/api/controller/v1"
	"powerssl.dev/backend/transport"

	"powerssl.dev/controller/internal/service/acme"
	"powerssl.dev/controller/internal/service/integration"
)

var Provider = wire.NewSet(
	ProvideRegisterF,
	acme.Provider,
	integration.Provider,
)

func ProvideRegisterF(
	acmeServiceServer apiv1.ACMEServiceServer,
	integrationServiceServer apiv1.IntegrationServiceServer,
) transport.RegisterF {
	return func(srv *transport.Server) {
		srv.RegisterService(acme.ServiceDesc, acmeServiceServer)
		srv.RegisterService(integration.ServiceDesc, integrationServiceServer)
	}
}
