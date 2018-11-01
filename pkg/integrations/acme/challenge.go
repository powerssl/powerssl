package acme

import (
	"time"

	"powerssl.io/pkg/controller/api"
)

func (acme *ACME) GetChallenge(accountURL string, challengeURL string) (*api.Challenge, error) {
	time.Sleep(1 * time.Second)

	return nil, ErrNotImplemented
}

func (acme *ACME) ValidateChallenge(accountURL string, challengeURL string) (*api.Challenge, error) {
	time.Sleep(1 * time.Second)

	return nil, ErrNotImplemented
}
