package cloudflare

import (
	"powerssl.dev/sdk/integration"
)

type Config struct {
	Integration integration.Config `flag:"integration"`
}

func (cfg *Config) Defaults() {
	cfg.Integration.Integration.Name = "cloudflare"
}
