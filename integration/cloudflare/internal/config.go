package cloudflare

import (
	"powerssl.dev/sdk/integration"
)

type Config struct {
	Integration *integration.Config
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
