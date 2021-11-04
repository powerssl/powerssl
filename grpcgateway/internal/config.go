package internal

import (
	"github.com/go-playground/validator/v10"

	"powerssl.dev/common/metrics"
	"powerssl.dev/common/transport"
	validator2 "powerssl.dev/common/validator"
	"powerssl.dev/grpcgateway/internal/server"
)

type Config struct {
	APIServerClient transport.ClientConfig
	Server          server.Config
	Metrics         metrics.Config
}

func (cfg *Config) Validate() error {
	validate := validator.New()
	validate.RegisterStructValidation(transport.ClientConfigValidator, transport.ClientConfig{})
	return validator2.ValidateConfig(validate, cfg)
}
