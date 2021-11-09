package internal

import (
	"github.com/google/wire"

	"powerssl.dev/apiserver/internal/service"
	"powerssl.dev/backend/temporal/client"
	backendtransport "powerssl.dev/backend/transport"
	"powerssl.dev/common/interrupthandler"
	"powerssl.dev/common/log"
	"powerssl.dev/common/metrics"
	"powerssl.dev/common/tracer"
	"powerssl.dev/common/transport"
)

var Provider = wire.NewSet(
	Provide,
	client.Provider,
	interrupthandler.Provider,
	log.Provider,
	metrics.Provider,
	service.Provider,
	backendtransport.Provider,
	tracer.Provider,
	transport.Provider,
	wire.FieldsOf(new(*Config), "DB", "Log", "Metrics", "Server", "TemporalClient", "Tracer"),
	wire.FieldsOf(new(ConfigDB), "Connection"),
)

func Provide(interruptHandlerF interrupthandler.F, metricsServerF metrics.F, serverF backendtransport.F) []func() error {
	return []func() error{
		interruptHandlerF,
		metricsServerF,
		serverF,
	}
}
