package acme

import (
	"powerssl.io/pkg/controller/api"
)

func (acme *ACME) GetChallenge(accountURL string, challengeURL string) (*api.Challenge, error) {
	return nil, ErrNotImplemented
}

func (acme *ACME) ValidateChallenge(accountURL string, challengeURL string) (*api.Challenge, error) {
	return nil, ErrNotImplemented
}
