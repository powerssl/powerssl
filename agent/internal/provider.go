package internal

import (
	"powerssl.dev/common/interrupthandler"
	"powerssl.dev/common/tracing"
	"powerssl.dev/sdk/apiserver"
)

const component = "powerssl-agent"

func provideRunnerF(interruptHandlerF interrupthandler.F, _ *apiserver.Client) []func() error {
	return []func() error{
		interruptHandlerF,
	}
}

func provideTracingComponent() tracing.TracerComponent {
	return component
}
