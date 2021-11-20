package internal

import (
	"github.com/google/wire"

	"powerssl.dev/common/interrupthandler"
	"powerssl.dev/common/log"
	"powerssl.dev/common/tracer"
	"powerssl.dev/sdk/apiserver"
)

var Provider = wire.NewSet(
	ConfigFields,
	Provide,
	apiserver.Provider,
	interrupthandler.Provider,
	log.Provider,
	tracer.Provider,
)

func Provide(interruptHandlerF interrupthandler.F, _ *apiserver.Client) []func() error {
	return []func() error{
		interruptHandlerF,
	}
}
