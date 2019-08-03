package grpcgateway

import (
	"gopkg.in/go-playground/validator.v9"

	"powerssl.dev/powerssl/internal/pkg/transport"
)

type APIServerClientConfig = transport.ClientConfig

type Config struct {
	APIServerClientConfig *APIServerClientConfig
	Addr                  string `validate:"required"`
	MetricsAddr           string
}

func (cfg *Config) Validate() error {
	validate := validator.New()
	validate.RegisterStructValidation(transport.ClientConfigValidator, transport.ClientConfig{})
	return validate.Struct(cfg)
}
