package acme

import (
	"errors"

	integrationacme "powerssl.dev/sdk/integration/acme"
	"powerssl.dev/sdk/integration/vault"
)

var ErrNotImplemented = errors.New("not implemented")

type ACME struct {
	vault *vault.Client
}

func New(vaultConfig vault.Config) (integrationacme.Integration, error) {
	vaultClient, err := vault.NewClient(vaultConfig)
	if err != nil {
		return nil, err
	}
	return &ACME{
		vault: vaultClient,
	}, nil
}
