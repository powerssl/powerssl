package internal

import (
	"powerssl.dev/apiserver/internal/repository"
	"powerssl.dev/backend/temporal/client"
	"powerssl.dev/backend/transport"
	"powerssl.dev/common/log"
	"powerssl.dev/common/metrics"
	"powerssl.dev/common/tracer"
)

const component = "powerssl-apiserver"

type Config struct {
	DB             repository.Config `flag:"db"`
	Log            log.Config        `flag:"log"`
	Metrics        metrics.Config    `flag:"metrics"`
	Server         transport.Config  `flag:"server"`
	TemporalClient client.Config     `flag:"temporalClient"`
	Tracer         tracer.Config     `flag:"tracer"`
}

func (cfg *Config) Defaults() {
	cfg.TemporalClient.Component = component
	cfg.Tracer.Component = component
	cfg.Server.VaultRole = component
}
