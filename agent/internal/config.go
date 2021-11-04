package internal

import (
	"github.com/go-playground/validator/v10"

	"powerssl.dev/common/tracing"
	"powerssl.dev/common/transport"
	validator2 "powerssl.dev/common/validator"
	"powerssl.dev/sdk/apiserver"
)

type Config struct {
	APIServerClientConfig transport.ClientConfig `mapstructure:"apiserver"`
	AuthToken             apiserver.AuthToken    `validate:"required"`
	Tracer                tracing.TracerImplementation
}

func (cfg *Config) Validate() error {
	validate := validator.New()
	return validator2.ValidateConfig(validate, cfg)
}
