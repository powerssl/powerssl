package internal

import (
	"github.com/google/wire"

	"powerssl.dev/backend/grpcserver"
	"powerssl.dev/backend/temporal/client"
	"powerssl.dev/backend/vault"
	"powerssl.dev/common/interrupthandler"
	"powerssl.dev/common/log"
	"powerssl.dev/common/telemetry"
	"powerssl.dev/common/transport"
	"powerssl.dev/controller/internal/service"
	"powerssl.dev/sdk/apiserver"

	"powerssl.dev/controller/internal/worker"
)

var Provider = wire.NewSet(
	ConfigFields,
	Provide,
	apiserver.Provider,
	grpcserver.Provider,
	client.Provider,
	interrupthandler.Provider,
	log.Provider,
	service.Provider,
	telemetry.Provider,
	transport.Provider,
	vault.Provider,
	worker.Provider,
)

func Provide(interruptHandlerF interrupthandler.F, serverF grpcserver.F, telemetryF telemetry.F, workerF worker.F) []func() error {
	return []func() error{
		interruptHandlerF,
		serverF,
		telemetryF,
		workerF,
	}
}
