package acme

import (
	"time"

	"powerssl.io/pkg/controller/api"
)

func (acme *ACME) CreateAuthorization(directoryURL string, accountURL string, identifier *api.Identifier) (*api.Authorization, error) {
	time.Sleep(1 * time.Second)

	return nil, ErrNotImplemented
}

func (acme *ACME) DeactivateAuthorization(accountURL string, authorizationURL string) (*api.Authorization, error) {
	time.Sleep(1 * time.Second)

	return nil, ErrNotImplemented
}

func (acme *ACME) GetAuthorization(accountURL string, authorizationURL string) (*api.Authorization, error) {
	time.Sleep(1 * time.Second)

	return nil, ErrNotImplemented
}
