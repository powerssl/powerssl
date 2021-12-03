package internal

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"

	"powerssl.dev/common/log"
	"powerssl.dev/common/telemetry"

	"powerssl.dev/webapp/internal/server"
)

const component = "powerssl-webapp"

var ConfigFields = wire.FieldsOf(new(*Config), "Log", "Server", "Telemetry")

type Config struct {
	Log       log.Config       `flag:"log"`
	Server    server.Config    `flag:"server"`
	Telemetry telemetry.Config `flag:"telemetry"`
}

func (cfg *Config) PreValidate(validate *validator.Validate) {
	cfg.Telemetry.Component = component
	cfg.Log.PreValidate(validate)
	cfg.Server.PreValidate(validate)
	cfg.Telemetry.PreValidate(validate)
}
