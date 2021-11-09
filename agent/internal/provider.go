package internal

import (
	"github.com/google/wire"

	"powerssl.dev/common/interrupthandler"
	"powerssl.dev/common/log"
	"powerssl.dev/common/tracer"
	"powerssl.dev/sdk/apiserver"
)

var Provider = wire.NewSet(
	Provide,
	apiserver.Provider,
	interrupthandler.Provider,
	log.Provider,
	tracer.Provider,
	wire.FieldsOf(new(*Config), "APIServerClient", "Log", "Tracer"),
)

func Provide(interruptHandlerF interrupthandler.F, _ *apiserver.Client) []func() error {
	return []func() error{
		interruptHandlerF,
	}
}
