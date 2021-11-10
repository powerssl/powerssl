package acme

import (
	"context"
	"time"

	apiv1 "powerssl.dev/api/controller/v1"
)

func (acme *ACME) CreateAuthorization(_ context.Context, directoryURL string, accountURL string, identifier *apiv1.Identifier) (*apiv1.Authorization, error) {
	time.Sleep(1 * time.Second)

	return nil, ErrNotImplemented
}

func (acme *ACME) DeactivateAuthorization(_ context.Context, accountURL string, authorizationURL string) (*apiv1.Authorization, error) {
	time.Sleep(1 * time.Second)

	return nil, ErrNotImplemented
}

func (acme *ACME) GetAuthorization(_ context.Context, accountURL string, authorizationURL string) (*apiv1.Authorization, error) {
	time.Sleep(1 * time.Second)

	return nil, ErrNotImplemented
}
