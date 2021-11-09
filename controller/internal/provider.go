package internal

import (
	"github.com/google/wire"

	"powerssl.dev/backend/temporal/client"
	backendtransport "powerssl.dev/backend/transport"
	"powerssl.dev/backend/vault"
	"powerssl.dev/common/interrupthandler"
	"powerssl.dev/common/log"
	"powerssl.dev/common/metrics"
	"powerssl.dev/common/tracer"
	"powerssl.dev/common/transport"
	"powerssl.dev/controller/internal/service"
	"powerssl.dev/sdk/apiserver"

	"powerssl.dev/controller/internal/worker"
)

var Provider = wire.NewSet(
	Provide,
	apiserver.Provider,
	backendtransport.Provider,
	client.Provider,
	interrupthandler.Provider,
	log.Provider,
	metrics.Provider,
	service.Provider,
	tracer.Provider,
	transport.Provider,
	vault.Provider,
	wire.FieldsOf(new(*Config), "APIServerClient", "Log", "Metrics", "Server", "TemporalClient", "Tracer", "VaultClient"),
	worker.Provide,
)

func Provide(interruptHandlerF interrupthandler.F, metricsServerF metrics.F, serverF backendtransport.F, workerF worker.F) []func() error {
	return []func() error{
		interruptHandlerF,
		metricsServerF,
		serverF,
		workerF,
	}
}
