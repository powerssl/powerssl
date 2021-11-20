package internal

import (
	"github.com/google/wire"

	"powerssl.dev/backend/grpcserver"
	"powerssl.dev/backend/temporal/client"
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
	ConfigFields,
	Provide,
	apiserver.Provider,
	grpcserver.Provider,
	client.Provider,
	interrupthandler.Provider,
	log.Provider,
	metrics.Provider,
	service.Provider,
	tracer.Provider,
	transport.Provider,
	vault.Provider,
	worker.Provider,
)

func Provide(interruptHandlerF interrupthandler.F, metricsF metrics.F, serverF grpcserver.F, workerF worker.F) []func() error {
	return []func() error{
		interruptHandlerF,
		metricsF,
		serverF,
		workerF,
	}
}
