//go:build wireinject

package internal

import (
	"context"

	"github.com/google/wire"

	"powerssl.dev/backend/temporal/client"
	"powerssl.dev/backend/transport"
	"powerssl.dev/common/interrupthandler"
	"powerssl.dev/common/log"
	"powerssl.dev/common/metrics"
	"powerssl.dev/common/tracing"

	"powerssl.dev/apiserver/internal/service"
)

func Initialize(ctx context.Context, cfg *Config) ([]func() error, func(), error) {
	panic(wire.Build(
		client.Provider,
		interrupthandler.Provider,
		log.Provider,
		metrics.Provider,
		provideRunnerF,
		provideTemporalClientComponent,
		provideTracingComponent,
		service.Provider,
		tracing.Provider,
		transport.Provider,
		wire.FieldsOf(new(*Config), "TemporalClientConfig", "ServerConfig", "DB", "Tracer", "Metrics"),
		wire.FieldsOf(new(ConfigDB), "Connection"),
	))
}
