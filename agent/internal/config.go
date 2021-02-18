package internal

import (
	"github.com/go-playground/validator/v10"

	"powerssl.dev/common"
	"powerssl.dev/common/transport"
)

type APIServerClientConfig = transport.ClientConfig

type Config struct {
	APIServerClientConfig APIServerClientConfig `mapstructure:"apiserver"`
	AuthToken             string                `validate:"required"`
}

func (cfg *Config) Validate() error {
	validate := validator.New()
	return common.ValidateConfig(validate, cfg)
}
