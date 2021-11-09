package internal

import (
	"powerssl.dev/apiserver/internal/repository"
	"powerssl.dev/backend/temporal/client"
	"powerssl.dev/backend/transport"
	"powerssl.dev/backend/vault"
	"powerssl.dev/common/log"
	"powerssl.dev/common/metrics"
	"powerssl.dev/common/tracer"
)

const component = "powerssl-apiserver"

type (
	Config struct {
		CAFile            string                 `flag:"caFile"`
		DB                ConfigDB               `flag:"db"`
		Log               log.Config             `flag:"log"`
		JWKS              ConfigJWKS             `flag:"jwks"`
		Metrics           metrics.Config         `flag:"metrics"`
		Server            transport.ServerConfig `flag:"server"`
		TemporalClient    client.Config          `flag:"temporalClient"`
		Tracer            tracer.Config          `flag:"tracer"`
		VaultClientConfig vault.ClientConfig     `flag:"vaultClient"`
	}
	ConfigDB struct {
		Connection repository.ConnString `flag:"connection" validate:"required"`
	}
	ConfigJWKS struct {
		InsecureSkipTLSVerify bool   `flag:"insecureSkipTLSVerify" mapstructure:"insecure-skip-tls-verify"`
		ServerNameOverride    string `flag:"serverNameOverride" mapstructure:"server-name-override"`
		URL                   string `flag:"url" validate:"required"`
	}
)

func (cfg *Config) Defaults() {
	cfg.TemporalClient.Component = component
	cfg.Tracer.Component = component
}
