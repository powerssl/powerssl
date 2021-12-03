package internal

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"

	"powerssl.dev/backend/temporal/client"
	"powerssl.dev/backend/vault"
	"powerssl.dev/common/log"
	"powerssl.dev/common/telemetry"
	"powerssl.dev/sdk/apiserver"
)

const component = "powerssl-worker"

var ConfigFields = wire.FieldsOf(new(*Config), "APIServerClient", "Log", "TemporalClient", "Telemetry", "VaultClient")

type Config struct {
	APIServerClient apiserver.Config `flag:"apiServerClient"`
	TemporalClient  client.Config    `flag:"temporalClient"`
	Telemetry       telemetry.Config `flag:"telemetry"`
	Log             log.Config       `flag:"log"`
	VaultClient     vault.Config     `flag:"vaultClient"`
}

func (cfg *Config) PreValidate(validate *validator.Validate) {
	cfg.TemporalClient.Component = component
	cfg.Telemetry.Component = component
	cfg.APIServerClient.PreValidate(validate)
	cfg.TemporalClient.PreValidate(validate)
	cfg.Telemetry.PreValidate(validate)
	cfg.Log.PreValidate(validate)
	cfg.Log.PreValidate(validate)
}
