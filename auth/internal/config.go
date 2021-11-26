package internal

import (
	"github.com/google/wire"

	"powerssl.dev/auth/internal/oauth2"
	"powerssl.dev/auth/internal/server"
	"powerssl.dev/common/log"
	"powerssl.dev/common/telemetry"
)

const component = "powerssl-auth"

var ConfigFields = wire.FieldsOf(new(*Config), "Log", "OAuth2", "Server", "Telemetry")

type Config struct {
	Log       log.Config       `flag:"log"`
	OAuth2    oauth2.Config    `flag:"oauth2"`
	Server    server.Config    `flag:"server"`
	Telemetry telemetry.Config `flag:"telemetry"`
}

func (cfg *Config) Defaults() {
	cfg.Telemetry.Component = component
}
