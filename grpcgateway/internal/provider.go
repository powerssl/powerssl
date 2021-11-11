package internal

import (
	"github.com/google/wire"

	"powerssl.dev/common/interrupthandler"
	"powerssl.dev/common/log"
	"powerssl.dev/common/metrics"
	"powerssl.dev/common/transport"

	"powerssl.dev/grpcgateway/internal/server"
)

var Provider = wire.NewSet(
	Provide,
	interrupthandler.Provider,
	log.Provider,
	metrics.Provider,
	server.Provider,
	transport.Provider,
	wire.FieldsOf(new(*Config), "APIServerClient", "Log", "Metrics", "Server"),
)

func Provide(interruptHandlerF interrupthandler.F, metricsF metrics.F, serverF server.F) []func() error {
	return []func() error{
		interruptHandlerF,
		metricsF,
		serverF,
	}
}
