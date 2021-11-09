package service

import (
	"github.com/google/wire"

	apiv1 "powerssl.dev/api/apiserver/v1"
	"powerssl.dev/apiserver/internal/repository"
	"powerssl.dev/apiserver/internal/service/acmeaccount"
	"powerssl.dev/apiserver/internal/service/acmeserver"
	"powerssl.dev/apiserver/internal/service/certificate"
	"powerssl.dev/apiserver/internal/service/user"
	"powerssl.dev/backend/transport"
)

var Provider = wire.NewSet(
	Provide,
	acmeaccount.Provider,
	acmeserver.Provider,
	certificate.Provider,
	repository.Provider,
	user.Provider,
)

func Provide(
	acmeAccountServiceServer apiv1.ACMEAccountServiceServer,
	acmeServerServiceServer apiv1.ACMEServerServiceServer,
	certificateServiceServer apiv1.CertificateServiceServer,
	userServiceServer apiv1.UserServiceServer,
) transport.Register {
	return func(srv *transport.Server) {
		srv.RegisterService(acmeaccount.ServiceDesc, acmeAccountServiceServer)
		srv.RegisterService(acmeserver.ServiceDesc, acmeServerServiceServer)
		srv.RegisterService(certificate.ServiceDesc, certificateServiceServer)
		srv.RegisterService(user.ServiceDesc, userServiceServer)
	}
}
