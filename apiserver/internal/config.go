package internal

import (
	"github.com/go-playground/validator/v10"

	"powerssl.dev/apiserver/internal/repository"
	"powerssl.dev/backend/temporal/client"
	"powerssl.dev/backend/transport"
	"powerssl.dev/backend/vault"
	"powerssl.dev/common/metrics"
	"powerssl.dev/common/tracing"
	validator2 "powerssl.dev/common/validator"
)

type (
	Config struct {
		CAFile               ConfigCAFile `mapstructure:"ca-file"`
		DB                   ConfigDB
		JWKS                 ConfigJWKS
		Metrics              metrics.Config
		ServerConfig         transport.ServerConfig `mapstructure:",squash"`
		TemporalClientConfig client.Config          `mapstructure:"temporal"`
		Tracer               tracing.TracerImplementation
		VaultClientConfig    vault.ClientConfig `mapstructure:"vault"`
	}
	ConfigCAFile string
	ConfigDB     struct {
		Connection repository.ConnString `validate:"required"`
		Dialect    ConfigDBDialect       `validate:"required"`
	}
	ConfigDBDialect string
	ConfigJWKS      struct {
		InsecureSkipTLSVerify ConfigJWKSInsecureSkipTLSVerify `mapstructure:"insecure-skip-tls-verify"`
		ServerNameOverride    ConfigJWKSServerNameOverride    `mapstructure:"server-name-override"`
		URL                   ConfigJWKSURL                   `validate:"required"`
	}
	ConfigJWKSInsecureSkipTLSVerify bool
	ConfigJWKSServerNameOverride    string
	ConfigJWKSURL                   string
)

func (cfg *Config) Validate() error {
	validate := validator.New()
	validate.RegisterStructValidation(transport.ServerConfigValidator, transport.ServerConfig{})
	return validator2.ValidateConfig(validate, cfg)
}
