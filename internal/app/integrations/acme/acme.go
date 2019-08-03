package acme

import (
	"errors"

	integrationacme "powerssl.dev/powerssl/pkg/integration/acme"
)

var ErrNotImplemented = errors.New("not implemented")

type ACME struct{}

func New() integrationacme.Integration {
	return &ACME{}
}
