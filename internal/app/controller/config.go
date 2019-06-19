package controller

import (
	"gopkg.in/go-playground/validator.v9"

	"powerssl.io/powerssl/internal/pkg/util"
)

type APIServerClientConfig = util.ClientConfig
type ServerConfig = util.ServerConfig

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
	validate.RegisterStructValidation(util.ClientConfigValidator, util.ClientConfig{})
	validate.RegisterStructValidation(util.ServerConfigValidator, util.ServerConfig{})
	return validate.Struct(cfg)
}
