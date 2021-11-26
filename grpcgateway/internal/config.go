package internal

import (
	"github.com/google/wire"

	"powerssl.dev/common/log"
	"powerssl.dev/common/telemetry"
	"powerssl.dev/common/transport"
	"powerssl.dev/grpcgateway/internal/server"
)

const component = "powerssl-grpcgateway"

var ConfigFields = wire.FieldsOf(new(*Config), "APIServerClient", "Log", "Server", "Telemetry")

type Config struct {
	APIServerClient transport.Config `flag:"apiServerClient"`
	Log             log.Config       `flag:"log"`
	Server          server.Config    `flag:"server"`
	Telemetry       telemetry.Config `flag:"telemetry"`
}

func (cfg *Config) Defaults() {
	cfg.Telemetry.Component = component
}
