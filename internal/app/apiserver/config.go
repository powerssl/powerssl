package apiserver

import (
	"gopkg.in/go-playground/validator.v9"

	"powerssl.dev/powerssl/internal/pkg/transport"
)

type ControllerClientConfig = transport.ClientConfig
type ServerConfig = transport.ServerConfig

type VaultClientConfig struct {
	Token  string
	URL    string
	CAFile string
}

type Config struct {
	AuthToken              string
	ControllerClientConfig *ControllerClientConfig
	DBConnection           string `validate:"required"`
	DBDialect              string `validate:"required"`
	JWKSURL                string `validate:"required"`
	MetricsAddr            string
	ServerConfig           *ServerConfig
	Tracer                 string
	VaultClientConfig      *VaultClientConfig
}

func (cfg *Config) Validate() error {
	validate := validator.New()
	validate.RegisterStructValidation(transport.ClientConfigValidator, transport.ClientConfig{})
	validate.RegisterStructValidation(transport.ServerConfigValidator, transport.ServerConfig{})
	return validate.Struct(cfg)
}
