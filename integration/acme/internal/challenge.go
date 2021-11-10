package acme

import (
	"context"
	"time"

	apiv1 "powerssl.dev/api/controller/v1"
)

func (acme *ACME) GetChallenge(_ context.Context, accountURL string, challengeURL string) (*apiv1.Challenge, error) {
	time.Sleep(1 * time.Second)

	return nil, ErrNotImplemented
}

func (acme *ACME) ValidateChallenge(_ context.Context, accountURL string, challengeURL string) (*apiv1.Challenge, error) {
	time.Sleep(1 * time.Second)

	return nil, ErrNotImplemented
}
