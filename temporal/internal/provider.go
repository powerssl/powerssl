package internal

import (
	"github.com/google/wire"

	"powerssl.dev/common/interrupthandler"
	"powerssl.dev/common/log"
	"powerssl.dev/temporal/internal/server"
)

var Provider = wire.NewSet(
	ConfigFields,
	Provide,
	interrupthandler.Provider,
	log.Provider,
	server.Provider,
)

func Provide(interruptHandlerF interrupthandler.F, serverF server.F) []func() error {
	return []func() error{
		interruptHandlerF,
		serverF,
	}
}
