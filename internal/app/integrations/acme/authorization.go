package acme

import (
	"context"
	"time"

	"powerssl.io/powerssl/pkg/controller/api"
)

func (acme *ACME) CreateAuthorization(_ context.Context, directoryURL string, accountURL string, identifier *api.Identifier) (*api.Authorization, error) {
	time.Sleep(1 * time.Second)

	return nil, ErrNotImplemented
}

func (acme *ACME) DeactivateAuthorization(_ context.Context, accountURL string, authorizationURL string) (*api.Authorization, error) {
	time.Sleep(1 * time.Second)

	return nil, ErrNotImplemented
}

func (acme *ACME) GetAuthorization(_ context.Context, accountURL string, authorizationURL string) (*api.Authorization, error) {
	time.Sleep(1 * time.Second)

	return nil, ErrNotImplemented
}
