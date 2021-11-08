//go:build wireinject

package integration

import (
	"context"

	"github.com/google/wire"

	"powerssl.dev/sdk/integration/acme"
	"powerssl.dev/sdk/integration/dns"
)

func InitializeACME(ctx context.Context, cfg *Config, handler acme.Integration) ([]func() error, func(), error) {
	panic(wire.Build(ProviderACME))
}

func InitializeDNS(ctx context.Context, cfg *Config, handler dns.Integration) ([]func() error, func(), error) {
	panic(wire.Build(ProviderDNS))
}
