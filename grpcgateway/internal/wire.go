//go:build wireinject

package internal

import (
	"context"

	"github.com/google/wire"

	"powerssl.dev/common/interrupthandler"
	"powerssl.dev/common/log"
	"powerssl.dev/common/metrics"
	"powerssl.dev/common/transport"
	"powerssl.dev/grpcgateway/internal/server"
)

func Initialize(ctx context.Context, cfg *Config) ([]func() error, func(), error) {
	panic(wire.Build(
		interrupthandler.Provider,
		log.Provider,
		metrics.Provider,
		provideRunnerF,
		server.Provider,
		transport.Provider,
		wire.FieldsOf(new(*Config), "APIServerClient", "Server", "Metrics"),
	))
}
