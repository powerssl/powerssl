//go:build wireinject

package internal

import (
	"context"

	"github.com/google/wire"

	"powerssl.dev/sdk/apiserver"
)

func newClient(ctx context.Context, cfg *Config) (*apiserver.Client, func(), error) {
	panic(wire.Build(Provider))
}
