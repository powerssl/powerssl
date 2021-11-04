//go:build wireinject

package internal

import (
	"context"

	"github.com/google/wire"

	"powerssl.dev/common/interrupthandler"
	"powerssl.dev/common/log"
	"powerssl.dev/common/tracing"
	"powerssl.dev/sdk/apiserver"
)

func Initialize(ctx context.Context, cfg *Config) ([]func() error, func(), error) {
	panic(wire.Build(
		apiserver.Provider,
		interrupthandler.Provider,
		log.Provider,
		provideRunnerF,
		provideTracingComponent,
		tracing.Provider,
		wire.FieldsOf(new(*Config),
			"APIServerClientConfig",
			"AuthToken",
			"Tracer",
		),
	))
}
