package acme

import (
	"powerssl.dev/sdk/integration"
	"powerssl.dev/sdk/integration/vault"
)

type Config struct {
	Integration integration.Config `flag:"integration"`
	Vault       vault.Config       `flag:"integration"`
}

func (cfg *Config) Defaults() {
	cfg.Integration.Integration.Name = "acme"
	cfg.Vault.Address = "https://localhost:8200" // TODO
}
