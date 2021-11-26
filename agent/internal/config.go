package internal

import (
	"github.com/google/wire"

	"powerssl.dev/common/log"
	"powerssl.dev/common/telemetry"
	"powerssl.dev/sdk/apiserver"
)

const component = "powerssl-agent"

var ConfigFields = wire.FieldsOf(new(*Config), "APIServerClient", "Log", "Telemetry")

type Config struct {
	APIServerClient apiserver.Config `flag:"apiServerClient"`
	Log             log.Config       `flag:"log"`
	Telemetry       telemetry.Config `flag:"telemetry"`
}

func (cfg *Config) Defaults() {
	cfg.Telemetry.Component = component
}
