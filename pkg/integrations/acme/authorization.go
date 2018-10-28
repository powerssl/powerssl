package acme

import (
	"powerssl.io/pkg/controller/api"
)

func (acme *ACME) CreateAuthorization(directoryURL string, accountURL string, identifier *api.Identifier) (*api.Authorization, error) {
	return nil, ErrNotImplemented
}

func (acme *ACME) DeactivateAuthorization(accountURL string, authorizationURL string) (*api.Authorization, error) {
	return nil, ErrNotImplemented
}

func (acme *ACME) GetAuthorization(accountURL string, authorizationURL string) (*api.Authorization, error) {
	return nil, ErrNotImplemented
}
