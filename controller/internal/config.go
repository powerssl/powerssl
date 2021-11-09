package internal

import (
	"powerssl.dev/backend/temporal/client"
	backendtransport "powerssl.dev/backend/transport"
	"powerssl.dev/backend/vault"
	"powerssl.dev/common/log"
	"powerssl.dev/common/metrics"
	"powerssl.dev/common/tracer"
	"powerssl.dev/sdk/apiserver"
)

const component = "powerssl-controller"

type Config struct {
	APIServerClient apiserver.Config              `flag:"apiServerClient" validate:"required"`
	Log             log.Config                    `flag:"log"`
	Metrics         metrics.Config                `flag:"metrics"`
	Server          backendtransport.ServerConfig `flag:"server"`
	TemporalClient  client.Config                 `flag:"temporalClient"`
	Tracer          tracer.Config                 `flag:"tracer"`
	VaultClient     vault.ClientConfig            `flag:"vaultClient"`
}

func (cfg *Config) Defaults() {
	cfg.TemporalClient.Component = component
	cfg.Tracer.Component = component
}
