package internal

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"

	"powerssl.dev/common/log"
	"powerssl.dev/common/telemetry"

	"powerssl.dev/auth/internal/oauth2"
	"powerssl.dev/auth/internal/server"
)

const component = "powerssl-auth"

var ConfigFields = wire.FieldsOf(new(*Config), "Log", "OAuth2", "Server", "Telemetry")

type Config struct {
	Log       log.Config       `flag:"log"`
	OAuth2    oauth2.Config    `flag:"oauth2"`
	Server    server.Config    `flag:"server"`
	Telemetry telemetry.Config `flag:"telemetry"`
}

func (cfg *Config) PreValidate(validate *validator.Validate) {
	cfg.Telemetry.Component = component
	cfg.Log.PreValidate(validate)
	cfg.OAuth2.PreValidate(validate)
	cfg.Server.PreValidate(validate)
	cfg.Telemetry.PreValidate(validate)
}
