package internal

import (
	"github.com/go-playground/validator/v10"

	"powerssl.dev/common/transport"
	"powerssl.dev/common/util"
)

type APIServerClientConfig = transport.ClientConfig

type Config struct {
	APIServerClientConfig APIServerClientConfig `mapstructure:"apiserver"`
	Addr                  string                `validate:"required,hostname_port"`
	Metrics               struct {
		Addr string
	}
}

func (cfg *Config) Validate() error {
	validate := validator.New()
	validate.RegisterStructValidation(transport.ClientConfigValidator, transport.ClientConfig{})
	return util.ValidateConfig(validate, cfg)
}