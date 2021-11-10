//go:build wireinject

package internal

import (
	"context"

	"github.com/google/wire"
)

func Initialize(ctx context.Context, cfg *Config) ([]func() error, func(), error) {
	panic(wire.Build(Provider))
}
