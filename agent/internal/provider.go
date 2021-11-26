package internal

import (
	"github.com/google/wire"
	"powerssl.dev/common/telemetry"

	"powerssl.dev/common/interrupthandler"
	"powerssl.dev/common/log"
	"powerssl.dev/sdk/apiserver"
)

var Provider = wire.NewSet(
	ConfigFields,
	Provide,
	apiserver.Provider,
	interrupthandler.Provider,
	log.Provider,
	telemetry.Provider,
)

func Provide(interruptHandlerF interrupthandler.F, _ *apiserver.Client) []func() error {
	return []func() error{
		interruptHandlerF,
	}
}
