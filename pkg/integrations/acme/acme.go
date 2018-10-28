package acme

import (
	integrationacme "powerssl.io/pkg/integration/acme"
)

type errorConst string

func (e errorConst) Error() string { return string(e) }

const ErrNotImplemented = errorConst("not implemented.")

type ACME struct{}

func New() integrationacme.Integration {
	return &ACME{}
}
