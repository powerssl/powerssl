package worker

import (
	"github.com/go-playground/validator/v10"

	temporalclient "powerssl.dev/powerssl/internal/pkg/temporal/client"
	"powerssl.dev/powerssl/internal/pkg/transport"
	"powerssl.dev/powerssl/internal/pkg/util"
	"powerssl.dev/powerssl/internal/pkg/vault"
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
	return util.ValidateConfig(validate, cfg)
}
