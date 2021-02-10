package signer

import (
	"github.com/go-playground/validator/v10"

	"powerssl.dev/powerssl/internal/pkg/transport"
)

type ServerConfig = transport.ServerConfig

type Config struct {
	Metrics struct {
		Addr string
	}
	ServerConfig ServerConfig `mapstructure:",squash"`
	Tracer       string
}

func (cfg *Config) Validate() error {
	validate := validator.New()
	validate.RegisterStructValidation(transport.ServerConfigValidator, transport.ServerConfig{})
	return validate.Struct(cfg)
}
