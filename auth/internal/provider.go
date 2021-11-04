package internal

import (
	"powerssl.dev/auth/internal/server"
	"powerssl.dev/common/interrupthandler"
	"powerssl.dev/common/metrics"
)

const component = "powerssl-auth"

func provideRunnerF(interruptHandlerF interrupthandler.F, metricsServerF metrics.F, serverF server.F) []func() error {
	return []func() error{
		interruptHandlerF,
		metricsServerF,
		serverF,
	}
}
