package internal

import (
	"github.com/google/wire"

	"powerssl.dev/backend/grpcserver"
	"powerssl.dev/backend/temporal/client"
	"powerssl.dev/backend/vault"
	"powerssl.dev/common/log"
	"powerssl.dev/common/metrics"
	"powerssl.dev/common/tracer"
	"powerssl.dev/sdk/apiserver"
)

const component = "powerssl-controller"

var ConfigFields = wire.FieldsOf(new(*Config), "APIServerClient", "Log", "Metrics", "Server", "TemporalClient", "Tracer", "VaultClient")

type Config struct {
	APIServerClient apiserver.Config        `flag:"apiServerClient" validate:"required"`
	Log             log.Config              `flag:"log"`
	Metrics         metrics.Config          `flag:"metrics"`
	Server          grpcserver.Config `flag:"server"`
	TemporalClient  client.Config           `flag:"temporalClient"`
	Tracer          tracer.Config           `flag:"tracer"`
	VaultClient     vault.Config            `flag:"vaultClient"`
}

func (cfg *Config) Defaults() {
	cfg.TemporalClient.Component = component
	cfg.Tracer.Component = component
	cfg.Server.VaultRole = component
}
