package internal

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"

	"powerssl.dev/common/log"
	"powerssl.dev/temporal/internal/server"
)

var ConfigFields = wire.FieldsOf(new(*Config), "Log", "Server")

type Config struct {
	Log    log.Config    `flag:"log"`
	Server server.Config `flag:"server"`
}

func (cfg *Config) PreValidate(validate *validator.Validate) {
	cfg.Log.PreValidate(validate)
	cfg.Server.PreValidate(validate)
}
