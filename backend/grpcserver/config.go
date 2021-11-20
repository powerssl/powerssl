package grpcserver

import "powerssl.dev/backend/vault"

type Config struct {
	Addr       string       `flag:"addr;;;server addr" validate:"required,hostname_port"`
	CertFile   string       `flag:"certFile;;;server Cert file"`
	CommonName string       `flag:"commonName;;;server common name"`
	Insecure   bool         `flag:"insecure;;;server insecure"`
	KeyFile    string       `flag:"keyFile;;;server key file"`
	Vault      vault.Config `flag:"vault"`
	VaultRole  string       `flag:"-"`
}
