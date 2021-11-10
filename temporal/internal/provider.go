package internal

import (
	"github.com/google/wire"

	"powerssl.dev/common/interrupthandler"
	"powerssl.dev/common/log"
	"powerssl.dev/common/metrics"
	"powerssl.dev/temporal/internal/server"
)

var Provider = wire.NewSet(
	Provide,
	interrupthandler.Provider,
	log.Provider,
	metrics.Provider,
	server.Provider,
	wire.FieldsOf(new(*Config), "Log", "Metrics", "Server"),
)

func Provide(interruptHandlerF interrupthandler.F, metricsF metrics.F, serverF server.F) []func() error {
	return []func() error{
		interruptHandlerF,
		metricsF,
		serverF,
	}
}
