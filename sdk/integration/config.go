package integration // import "powerssl.dev/sdk/integration"

import (
	"github.com/go-playground/validator/v10"

	"powerssl.dev/common/transport"
)

type ControllerClientConfig = transport.ClientConfig

type Config struct {
	AuthToken              string                 `mapstructure:"auth-token" validate:"required"`
	ControllerClientConfig ControllerClientConfig `mapstructure:"controller"`
	Metrics                struct {
		Addr string
	}
	Tracer string
}

// TODO: Improve config validation errors
func (cfg *Config) Validate() error {
	validate := validator.New()
	validate.RegisterStructValidation(transport.ClientConfigValidator, transport.ClientConfig{})
	return validate.Struct(cfg)
}
