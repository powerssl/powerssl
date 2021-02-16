package internal

import (
	"github.com/go-playground/validator/v10"

	temporalclient "powerssl.dev/backend/temporal/client"
	backendtransport "powerssl.dev/backend/transport"
	"powerssl.dev/backend/vault"
	"powerssl.dev/common/transport"
	"powerssl.dev/common/util"
)

type APIServerClientConfig = transport.ClientConfig
type ServerConfig = backendtransport.ServerConfig
type TemporalClientConfig = temporalclient.Config
type VaultClientConfig = vault.ClientConfig

type Config struct {
	APIServerClientConfig APIServerClientConfig `mapstructure:"apiserver"`
	AuthToken             string                `mapstructure:"auth-token" validate:"required"`
	Metrics               struct {
		Addr string
	}
	ServerConfig         ServerConfig         `mapstructure:",squash"`
	TemporalClientConfig TemporalClientConfig `mapstructure:"temporal"`
	Tracer               string
	VaultClientConfig    VaultClientConfig `mapstructure:"vault"`
}

func (cfg *Config) Validate() error {
	validate := validator.New()
	validate.RegisterStructValidation(transport.ClientConfigValidator, transport.ClientConfig{})
	validate.RegisterStructValidation(backendtransport.ServerConfigValidator, backendtransport.ServerConfig{})
	return util.ValidateConfig(validate, cfg)
}
