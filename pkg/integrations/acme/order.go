package acme

import (
	"crypto/x509"
	"time"

	"powerssl.io/pkg/controller/api"
)

func (acme *ACME) CreateOrder(directoryURL string, accountURL string, identifiers []*api.Identifier, notBefore, notAfter string) (*api.Order, error) {
	time.Sleep(1 * time.Second)

	return nil, ErrNotImplemented
}

func (acme *ACME) FinalizeOrder(accountURL string, orderURL string, certificateSigningRequest *x509.CertificateRequest) (*api.Order, error) {
	time.Sleep(1 * time.Second)

	return nil, ErrNotImplemented
}

func (acme *ACME) GetOrder(accountURL string, orderURL string) (*api.Order, error) {
	time.Sleep(1 * time.Second)

	return nil, ErrNotImplemented
}
