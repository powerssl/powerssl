package internal

import (
	"powerssl.dev/common/interrupthandler"
	"powerssl.dev/common/metrics"

	"powerssl.dev/grpcgateway/internal/server"
)

const component = "powerssl-grpcgateway"

func provideRunnerF(interruptHandlerF interrupthandler.F, metricsServerF metrics.F, serverF server.F) []func() error {
	return []func() error{
		interruptHandlerF,
		metricsServerF,
		serverF,
	}
}
