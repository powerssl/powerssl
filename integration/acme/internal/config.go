package acme

import (
	"github.com/go-playground/validator/v10"

	"powerssl.dev/sdk/integration"
	"powerssl.dev/sdk/integration/vault"
)

type Config struct {
	Integration integration.Config `flag:"integration"`
	Vault       vault.Config       `flag:"integration"`
}

func (cfg *Config) PreValidate(validate *validator.Validate) {
	cfg.Integration.Integration.Name = "acme"
	cfg.Vault.Address = "https://localhost:8200" // TODO
	cfg.Integration.PreValidate(validate)
	cfg.Vault.PreValidate(validate)
}
