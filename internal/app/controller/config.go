package controller

import (
	"gopkg.in/go-playground/validator.v9"

	"powerssl.dev/powerssl/internal/pkg/transport"
)

type APIServerClientConfig = transport.ClientConfig
type ServerConfig = transport.ServerConfig

type VaultClientConfig struct {
	Token  string
	URL    string
	CAFile string
}

type Config struct {
	APIServerClientConfig *APIServerClientConfig
	AuthToken             string
	JWKSURL               string `validate:"required"`
	MetricsAddr           string
	ServerConfig          *ServerConfig
	Tracer                string
	VaultClientConfig     *VaultClientConfig
}

func (cfg *Config) Validate() error {
	validate := validator.New()
	validate.RegisterStructValidation(transport.ClientConfigValidator, transport.ClientConfig{})
	validate.RegisterStructValidation(transport.ServerConfigValidator, transport.ServerConfig{})
	return validate.Struct(cfg)
}
