package internal

import (
	"github.com/google/wire"
	"powerssl.dev/common/interrupthandler"
	"powerssl.dev/common/log"
	"powerssl.dev/common/telemetry"
	"powerssl.dev/grpcgateway/internal/server"
)

var Provider = wire.NewSet(
	ConfigFields,
	Provide,
	interrupthandler.Provider,
	log.Provider,
	server.Provider,
	telemetry.Provider,
)

func Provide(interruptHandlerF interrupthandler.F, serverF server.F, telemetryF telemetry.F) []func() error {
	return []func() error{
		interruptHandlerF,
		serverF,
		telemetryF,
	}
}
