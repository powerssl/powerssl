package internal

import (
	"github.com/go-playground/validator/v10"

	"powerssl.dev/auth/internal/oauth2"
	"powerssl.dev/auth/internal/server"
	"powerssl.dev/common/metrics"
	validator2 "powerssl.dev/common/validator"
)

type Config struct {
	Server  server.Config
	OAuth2  oauth2.Config
	Metrics metrics.Config
}

func (cfg *Config) Validate() error {
	validate := validator.New()
	return validator2.ValidateConfig(validate, cfg)
}
