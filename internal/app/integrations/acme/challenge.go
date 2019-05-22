package acme

import (
	"context"
	"time"

	"powerssl.io/powerssl/pkg/controller/api"
)

func (acme *ACME) GetChallenge(_ context.Context, accountURL string, challengeURL string) (*api.Challenge, error) {
	time.Sleep(1 * time.Second)

	return nil, ErrNotImplemented
}

func (acme *ACME) ValidateChallenge(_ context.Context, accountURL string, challengeURL string) (*api.Challenge, error) {
	time.Sleep(1 * time.Second)

	return nil, ErrNotImplemented
}
