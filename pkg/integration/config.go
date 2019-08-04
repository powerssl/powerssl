package integration // import "powerssl.dev/powerssl/pkg/integration"

import (
	"gopkg.in/go-playground/validator.v9"

	"powerssl.dev/powerssl/internal/pkg/transport"
)

type ControllerClientConfig = transport.ClientConfig

type Config struct {
	AuthToken              string `validate:"required"`
	ControllerClientConfig *ControllerClientConfig
	MetricsAddr            string
	Tracer                 string
}

func (cfg *Config) Validate() error {
	validate := validator.New()
	validate.RegisterStructValidation(transport.ClientConfigValidator, transport.ClientConfig{})
	return validate.Struct(cfg)
}
