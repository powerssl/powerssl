package acme

import "errors"

var ErrNotImplemented = errors.New("not implemented.")

type ACME struct{}

func New() ACME {
	return ACME{}
}
