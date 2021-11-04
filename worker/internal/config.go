package internal

import (
	"github.com/go-playground/validator/v10"

	temporalclient "powerssl.dev/backend/temporal/client"
	"powerssl.dev/backend/vault"
	"powerssl.dev/common/transport"
	validator2 "powerssl.dev/common/validator"
)

type APIServerClientConfig = transport.ClientConfig
type TemporalClientConfig = temporalclient.Config
type VaultClientConfig = vault.ClientConfig

type Config struct {
	APIServerClientConfig APIServerClientConfig `mapstructure:"apiserver"`
	AuthToken             string                `mapstructure:"auth-token" validate:"required"`
	Metrics               struct {
		Addr string
	}
	TemporalClientConfig TemporalClientConfig `mapstructure:"temporal"`
	Tracer               string
	VaultClientConfig    VaultClientConfig `mapstructure:"vault"`
}

func (cfg *Config) Validate() error {
	validate := validator.New()
	validate.RegisterStructValidation(transport.ClientConfigValidator, transport.ClientConfig{})
	return validator2.ValidateConfig(validate, cfg)
}
