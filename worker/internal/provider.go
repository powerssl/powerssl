package internal

import (
	"powerssl.dev/backend/temporal/client"
	"powerssl.dev/common/interrupthandler"
	"powerssl.dev/common/metrics"
	"powerssl.dev/common/tracing"

	"powerssl.dev/worker/internal/worker"
)

const component = "powerssl-worker"

func provideRunnerF(interruptHandlerF interrupthandler.F, metricsServerF metrics.F, workerF worker.F) []func() error {
	return []func() error{
		interruptHandlerF,
		metricsServerF,
		workerF,
	}
}

func provideTemporalClientComponent() client.TemporalClientComponent {
	return component
}

func provideTracingComponent() tracing.TracerComponent {
	return component
}
