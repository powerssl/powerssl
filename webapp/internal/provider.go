package internal

import (
	"github.com/google/wire"

	"powerssl.dev/common/interrupthandler"
	"powerssl.dev/common/log"
	"powerssl.dev/common/metrics"

	"powerssl.dev/webapp/internal/server"
)

var Provider = wire.NewSet(
	ConfigFields,
	Provide,
	interrupthandler.Provider,
	log.Provider,
	metrics.Provider,
	server.Provider,
)

func Provide(interruptHandlerF interrupthandler.F, metricsF metrics.F, serverF server.F) []func() error {
	return []func() error{
		interruptHandlerF,
		metricsF,
		serverF,
	}
}
