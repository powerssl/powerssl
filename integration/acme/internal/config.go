package acme

import (
	"powerssl.dev/sdk/integration"
	"powerssl.dev/sdk/integration/vault"
)

type Config struct {
	Integration *integration.Config
	Vault       vault.Config
}

func NewConfig(name integration.IntegrationName) *Config {
	return &Config{
		Integration: &integration.Config{
			Integration: integration.IntegrationConfig{
				Name: name,
			},
		},
	}
}
