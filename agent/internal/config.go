package internal

import (
	"github.com/google/wire"

	"powerssl.dev/common/log"
	"powerssl.dev/common/tracer"
	"powerssl.dev/sdk/apiserver"
)

const component = "powerssl-agent"

var ConfigFields = wire.FieldsOf(new(*Config), "APIServerClient", "Log", "Tracer")

type Config struct {
	APIServerClient apiserver.Config `flag:"apiServerClient"`
	Log             log.Config       `flag:"log"`
	Tracer          tracer.Config    `flag:"tracer"`
}

func (cfg *Config) Defaults() {
	cfg.Tracer.Component = component
}
