package agent

import (
	"gopkg.in/go-playground/validator.v9"

	"powerssl.io/powerssl/internal/pkg/transport"
)

type APIServerClientConfig = transport.ClientConfig

type Config struct {
	APIServerClientConfig *APIServerClientConfig
	AuthToken             string `validate:"required"`
}

func (cfg *Config) Validate() error {
	validate := validator.New()
	validate.RegisterStructValidation(transport.ClientConfigValidator, transport.ClientConfig{})
	return validate.Struct(cfg)
}
