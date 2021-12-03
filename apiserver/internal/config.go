package internal

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"

	"powerssl.dev/backend/grpcserver"
	"powerssl.dev/backend/temporal/client"
	"powerssl.dev/common/log"
	"powerssl.dev/common/telemetry"

	"powerssl.dev/apiserver/internal/repository"
)

const component = "powerssl-apiserver"

var ConfigFields = wire.FieldsOf(new(*Config), "DB", "Log", "Server", "TemporalClient", "Telemetry")

type Config struct {
	DB             repository.Config `flag:"db"`
	Log            log.Config        `flag:"log"`
	Server         grpcserver.Config `flag:"server"`
	Telemetry      telemetry.Config  `flag:"telemetry"`
	TemporalClient client.Config     `flag:"temporalClient"`
}

func (cfg *Config) PreValidate(validate *validator.Validate) {
	cfg.Server.VaultRole = component
	cfg.Telemetry.Component = component
	cfg.DB.PreValidate(validate)
	cfg.Log.PreValidate(validate)
	cfg.Server.PreValidate(validate)
	cfg.Telemetry.PreValidate(validate)
	cfg.TemporalClient.PreValidate(validate)
}
