package internal

import (
	"github.com/google/wire"

	"powerssl.dev/backend/temporal/client"
	"powerssl.dev/backend/vault"
	"powerssl.dev/common/interrupthandler"
	"powerssl.dev/common/log"
	"powerssl.dev/common/metrics"
	"powerssl.dev/common/tracer"
	"powerssl.dev/sdk/apiserver"
	"powerssl.dev/worker/internal/worker"
)

var Provider = wire.NewSet(
	ConfigFields,
	Provide,
	apiserver.Provider,
	client.Provider,
	interrupthandler.Provider,
	log.Provider,
	metrics.Provider,
	tracer.Provider,
	vault.Provider,
	worker.Provider,
)

func Provide(interruptHandlerF interrupthandler.F, metricsF metrics.F, workerF worker.F) []func() error {
	return []func() error{
		interruptHandlerF,
		metricsF,
		workerF,
	}
}
