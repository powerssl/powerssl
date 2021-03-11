package internal

import (
	"github.com/go-playground/validator/v10"

	temporalclient "powerssl.dev/backend/temporal/client"
	"powerssl.dev/backend/transport"
	"powerssl.dev/backend/vault"
	"powerssl.dev/common"
)

type ServerConfig = transport.ServerConfig
type TemporalClientConfig = temporalclient.Config
type VaultClientConfig = vault.ClientConfig

type Config struct {
	CAFile string `mapstructure:"ca-file"`
	DB     struct {
		Connection string `validate:"required"`
		Dialect    string `validate:"required"`
	}
	JWKS struct {
		InsecureSkipTLSVerify bool   `mapstructure:"insecure-skip-tls-verify"`
		ServerNameOverride    string `mapstructure:"server-name-override"`
		URL                   string `validate:"required"`
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
	return common.ValidateConfig(validate, cfg)
}
