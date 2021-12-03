package cloudflare

import (
	"github.com/go-playground/validator/v10"

	"powerssl.dev/sdk/integration"
)

type Config struct {
	Integration integration.Config `flag:"integration"`
}

func (cfg *Config) PreValidate(validate *validator.Validate) {
	cfg.Integration.Integration.Name = "cloudflare"
	cfg.Integration.PreValidate(validate)
}
