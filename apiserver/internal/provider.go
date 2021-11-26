package internal

import (
	"github.com/google/wire"

	"powerssl.dev/apiserver/internal/service"
	"powerssl.dev/backend/grpcserver"
	"powerssl.dev/backend/temporal/client"
	"powerssl.dev/common/interrupthandler"
	"powerssl.dev/common/log"
	"powerssl.dev/common/telemetry"
	"powerssl.dev/common/transport"
)

var Provider = wire.NewSet(
	ConfigFields,
	Provide,
	client.Provider,
	interrupthandler.Provider,
	log.Provider,
	service.Provider,
	grpcserver.Provider,
	telemetry.Provider,
	transport.Provider,
)

func Provide(interruptHandlerF interrupthandler.F, serverF grpcserver.F, telemetryF telemetry.F) []func() error {
	return []func() error{
		interruptHandlerF,
		serverF,
		telemetryF,
	}
}
