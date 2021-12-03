package grpcserver

import (
	"github.com/go-playground/validator/v10"

	"powerssl.dev/backend/vault"
)

type Config struct {
	Addr       string       `flag:"addr" flag-desc:"server addr" validate:"required,hostname_port"`
	CertFile   string       `flag:"certFile" flag-desc:"server Cert file"`
	CommonName string       `flag:"commonName" flag-desc:"server common name"`
	Insecure   bool         `flag:"insecure" flag-desc:"server insecure"`
	KeyFile    string       `flag:"keyFile" flag-desc:"server key file"`
	Vault      vault.Config `flag:"vault"`
	VaultRole  string       `flag:"-"`
}

func (cfg *Config) PreValidate(_ *validator.Validate) {}
