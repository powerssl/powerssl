//go:build wireinject

package internal

import (
	"context"

	"github.com/google/wire"

	"powerssl.dev/backend/temporal/client"
	"powerssl.dev/backend/vault"
	"powerssl.dev/common/interrupthandler"
	"powerssl.dev/common/log"
	"powerssl.dev/common/metrics"
	"powerssl.dev/common/tracing"
	"powerssl.dev/sdk/apiserver"

	"powerssl.dev/worker/internal/worker"
)

func Initialize(ctx context.Context, cfg *Config) ([]func() error, func(), error) {
	panic(wire.Build(
		apiserver.Provider,
		client.Provider,
		interrupthandler.Provider,
		log.Provider,
		metrics.Provider,
		provideRunnerF,
		provideTemporalClientComponent,
		provideTracingComponent,
		tracing.Provider,
		vault.Provider,
		wire.FieldsOf(new(*Config),
			"APIServerClientConfig",
			"AuthToken",
			"Metrics",
			"TemporalClientConfig",
			"Tracer",
			"VaultClientConfig",
		),
		worker.Provide,
	))
}
