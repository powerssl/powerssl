package internal

import (
	"github.com/google/wire"

	"powerssl.dev/common/interrupthandler"
	"powerssl.dev/common/log"
	"powerssl.dev/common/telemetry"

	"powerssl.dev/auth/internal/oauth2"
	"powerssl.dev/auth/internal/server"
)

var Provider = wire.NewSet(
	Provide,
	ConfigFields,
	interrupthandler.Provider,
	log.Provider,
	oauth2.Provider,
	server.Provider,
	telemetry.Provider,
)

func Provide(interruptHandlerF interrupthandler.F, telemetryF telemetry.F, serverF server.F) []func() error {
	return []func() error{
		interruptHandlerF,
		serverF,
		telemetryF,
	}
}
