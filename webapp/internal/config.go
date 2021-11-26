package internal

import (
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

func (cfg *Config) Defaults() {
	cfg.Telemetry.Component = component
}
