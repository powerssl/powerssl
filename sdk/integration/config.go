package integration // import "powerssl.dev/sdk/integration"

import (
	"github.com/go-playground/validator/v10"

	"powerssl.dev/common/metrics"
	"powerssl.dev/common/tracing"
	"powerssl.dev/common/transport"
	"powerssl.dev/sdk/integration/internal"

	"powerssl.dev/sdk/controller"
)

type IntegrationConfig = internal.IntegrationConfig

type IntegrationName = internal.IntegrationName

type Config struct {
	Integration            internal.IntegrationConfig
	AuthToken              controller.AuthToken   `mapstructure:"auth-token" validate:"required"`
	ControllerClientConfig transport.ClientConfig `mapstructure:"controller"`
	Metrics                metrics.Config
	Tracer                 tracing.TracerImplementation
}

// TODO: Improve config validation errors
func (cfg *Config) Validate() error {
	validate := validator.New()
	validate.RegisterStructValidation(transport.ClientConfigValidator, transport.ClientConfig{})
	return validate.Struct(cfg)
}
