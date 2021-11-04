package internal

import (
	"powerssl.dev/backend/temporal/client"
	backendtransport "powerssl.dev/backend/transport"
	"powerssl.dev/common/interrupthandler"
	"powerssl.dev/common/metrics"
	"powerssl.dev/common/tracing"
)

const component = "powerssl-apiserver"

func provideRunnerF(interruptHandlerF interrupthandler.F, metricsServerF metrics.F, serverF backendtransport.ServerF) []func() error {
	return []func() error{
		interruptHandlerF,
		metricsServerF,
		serverF,
	}
}

func provideTemporalClientComponent() client.TemporalClientComponent {
	return component
}

func provideTracingComponent() tracing.TracerComponent {
	return component
}
