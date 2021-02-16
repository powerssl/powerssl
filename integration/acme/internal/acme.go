package acme

import (
	"errors"

	integrationacme "powerssl.dev/sdk/integration/acme"
)

var ErrNotImplemented = errors.New("not implemented")

type ACME struct{}

func New() integrationacme.Integration {
	return &ACME{}
}
