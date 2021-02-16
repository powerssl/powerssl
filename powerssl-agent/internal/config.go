package internal

import (
	"github.com/go-playground/validator/v10"

	"powerssl.dev/common/transport"
	"powerssl.dev/common/util"
)

type APIServerClientConfig = transport.ClientConfig

type Config struct {
	APIServerClientConfig APIServerClientConfig `mapstructure:"apiserver"`
	AuthToken             string                `validate:"required"`
}

func (cfg *Config) Validate() error {
	validate := validator.New()
	return util.ValidateConfig(validate, cfg)
}