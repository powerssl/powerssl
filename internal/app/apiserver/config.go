package apiserver

import (
	"github.com/go-playground/validator/v10"

	temporalclient "powerssl.dev/powerssl/internal/pkg/temporal/client"
	"powerssl.dev/powerssl/internal/pkg/transport"
	"powerssl.dev/powerssl/internal/pkg/util"
	"powerssl.dev/powerssl/internal/pkg/vault"
)

type ServerConfig = transport.ServerConfig
type TemporalClientConfig = temporalclient.Config
type VaultClientConfig = vault.ClientConfig

type Config struct {
	DB struct {
		Connection string `validate:"required"`
		Dialect    string `validate:"required"`
	}
	JWKS struct {
		URL string `validate:"required"`
	}
	Metrics struct {
		Addr string
	}
	ServerConfig         ServerConfig         `mapstructure:",squash"`
	TemporalClientConfig TemporalClientConfig `mapstructure:"temporal"`
	Tracer               string
	VaultClientConfig    VaultClientConfig `mapstructure:"vault"`
}

func (cfg *Config) Validate() error {
	validate := validator.New()
	validate.RegisterStructValidation(transport.ServerConfigValidator, transport.ServerConfig{})
	return util.ValidateConfig(validate, cfg)
}
