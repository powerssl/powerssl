package signer

import (
	"gopkg.in/go-playground/validator.v9"

	"powerssl.dev/powerssl/internal/pkg/transport"
)

type ServerConfig = transport.ServerConfig

type Config struct {
	MetricsAddr  string
	ServerConfig *ServerConfig
	Tracer       string
}

func (cfg *Config) Validate() error {
	validate := validator.New()
	validate.RegisterStructValidation(transport.ServerConfigValidator, transport.ServerConfig{})
	return validate.Struct(cfg)
}
