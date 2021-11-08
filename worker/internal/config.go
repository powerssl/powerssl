package internal

import (
	"github.com/go-playground/validator/v10"

	"powerssl.dev/backend/temporal/client"
	"powerssl.dev/backend/vault"
	"powerssl.dev/common/metrics"
	"powerssl.dev/common/tracing"
	"powerssl.dev/common/transport"
	validator2 "powerssl.dev/common/validator"
	"powerssl.dev/sdk/apiserver"
)

type Config struct {
	APIServerClientConfig transport.ClientConfig `mapstructure:"apiserver"`
	AuthToken             apiserver.AuthToken    `mapstructure:"auth-token" validate:"required"`
	Metrics               metrics.Config
	TemporalClientConfig  client.Config `mapstructure:"temporal"`
	Tracer                tracing.TracerImplementation
	VaultClientConfig     vault.ClientConfig `mapstructure:"vault"`
}

func (cfg *Config) Validate() error {
	validate := validator.New()
	validate.RegisterStructValidation(transport.ClientConfigValidator, transport.ClientConfig{})
	return validator2.ValidateConfig(validate, cfg)
}
