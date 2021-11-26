package internal

import (
	"github.com/google/wire"

	"powerssl.dev/backend/grpcserver"
	"powerssl.dev/backend/temporal/client"
	"powerssl.dev/backend/vault"
	"powerssl.dev/common/log"
	"powerssl.dev/common/telemetry"
	"powerssl.dev/sdk/apiserver"
)

const component = "powerssl-controller"

var ConfigFields = wire.FieldsOf(new(*Config), "APIServerClient", "Log", "Server", "TemporalClient", "Telemetry", "VaultClient")

type Config struct {
	APIServerClient apiserver.Config  `flag:"apiServerClient" validate:"required"`
	Log             log.Config        `flag:"log"`
	Server          grpcserver.Config `flag:"server"`
	TemporalClient  client.Config     `flag:"temporalClient"`
	Telemetry       telemetry.Config  `flag:"telemetry"`
	VaultClient     vault.Config      `flag:"vaultClient"`
}

func (cfg *Config) Defaults() {
	cfg.Server.VaultRole = component
	cfg.TemporalClient.Component = component
	cfg.Telemetry.Component = component
}
