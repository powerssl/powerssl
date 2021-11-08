package internal

import (
	"powerssl.dev/backend/temporal/client"
	backendtransport "powerssl.dev/backend/transport"
	"powerssl.dev/common/interrupthandler"
	"powerssl.dev/common/metrics"
	"powerssl.dev/common/tracing"

	"powerssl.dev/controller/internal/worker"
)

const component = "powerssl-controller"

func provideRunnerF(interruptHandlerF interrupthandler.F, metricsServerF metrics.F, serverF backendtransport.ServerF, workerF worker.F) []func() error {
	return []func() error{
		interruptHandlerF,
		metricsServerF,
		serverF,
		workerF,
	}
}

func provideTemporalClientComponent() client.TemporalClientComponent {
	return component
}

func provideTracingComponent() tracing.TracerComponent {
	return component
}
