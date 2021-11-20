package internal

import (
	"github.com/google/wire"

	"powerssl.dev/backend/temporal/client"
	"powerssl.dev/backend/vault"
	"powerssl.dev/common/log"
	"powerssl.dev/common/metrics"
	"powerssl.dev/common/tracer"
	"powerssl.dev/sdk/apiserver"
)

const component = "powerssl-worker"

var ConfigFields = wire.FieldsOf(new(*Config), "APIServerClient", "Log", "Metrics", "TemporalClient", "Tracer", "VaultClient")

type Config struct {
	APIServerClient apiserver.Config `flag:"apiServerClient"`
	Metrics         metrics.Config   `flag:"metrics"`
	TemporalClient  client.Config    `flag:"temporalClient"`
	Tracer          tracer.Config    `flag:"tracer"`
	Log             log.Config       `flag:"log"`
	VaultClient     vault.Config     `flag:"vaultClient"`
}

func (cfg *Config) Defaults() {
	cfg.TemporalClient.Component = component
	cfg.Tracer.Component = component
}
