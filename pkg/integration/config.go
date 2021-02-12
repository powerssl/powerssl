package integration // import "powerssl.dev/powerssl/pkg/integration"

import (
	"github.com/go-playground/validator/v10"

	"powerssl.dev/powerssl/internal/pkg/transport"
	"powerssl.dev/powerssl/internal/pkg/util"
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

func (cfg *Config) Validate() error {
	validate := validator.New()
	validate.RegisterStructValidation(transport.ClientConfigValidator, transport.ClientConfig{})
	return util.ValidateConfig(validate, cfg)
}
