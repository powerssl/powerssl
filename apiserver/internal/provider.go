package internal

import (
	"github.com/google/wire"

	"powerssl.dev/apiserver/internal/service"
	"powerssl.dev/backend/grpcserver"
	"powerssl.dev/backend/temporal/client"
	"powerssl.dev/common/interrupthandler"
	"powerssl.dev/common/log"
	"powerssl.dev/common/metrics"
	"powerssl.dev/common/tracer"
	"powerssl.dev/common/transport"
)

var Provider = wire.NewSet(
	ConfigFields,
	Provide,
	client.Provider,
	interrupthandler.Provider,
	log.Provider,
	metrics.Provider,
	service.Provider,
	grpcserver.Provider,
	tracer.Provider,
	transport.Provider,
)

func Provide(interruptHandlerF interrupthandler.F, metricsF metrics.F, serverF grpcserver.F) []func() error {
	return []func() error{
		interruptHandlerF,
		metricsF,
		serverF,
	}
}
