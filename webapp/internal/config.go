package internal

import (
	"github.com/go-playground/validator/v10"

	"powerssl.dev/common/metrics"
	validator2 "powerssl.dev/common/validator"
	"powerssl.dev/webapp/internal/server"
)

type Config struct {
	Metrics metrics.Config
	Server  server.Config
}

func (cfg *Config) Validate() error {
	validate := validator.New()
	return validator2.ValidateConfig(validate, cfg)
}
