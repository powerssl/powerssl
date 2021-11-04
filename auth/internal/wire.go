//go:build wireinject

package internal

import (
	"context"

	"github.com/google/wire"

	"powerssl.dev/common/interrupthandler"
	"powerssl.dev/common/log"
	"powerssl.dev/common/metrics"

	"powerssl.dev/auth/internal/oauth2"
	"powerssl.dev/auth/internal/server"
)

func Initialize(ctx context.Context, cfg *Config) ([]func() error, func(), error) {
	panic(wire.Build(
		interrupthandler.Provider,
		log.Provider,
		metrics.Provider,
		oauth2.Provider,
		provideRunnerF,
		server.Provider,
		wire.FieldsOf(new(*Config), "Server", "OAuth2", "Metrics"),
	))
}
