package acme

import (
	"context"
	"crypto/x509"
	"time"

	"powerssl.io/pkg/controller/api"
)

func (acme *ACME) CreateOrder(_ context.Context, directoryURL string, accountURL string, identifiers []*api.Identifier, notBefore, notAfter string) (*api.Order, error) {
	time.Sleep(1 * time.Second)

	return nil, ErrNotImplemented
}

func (acme *ACME) FinalizeOrder(_ context.Context, accountURL string, orderURL string, certificateSigningRequest *x509.CertificateRequest) (*api.Order, error) {
	time.Sleep(1 * time.Second)

	return nil, ErrNotImplemented
}

func (acme *ACME) GetOrder(_ context.Context, accountURL string, orderURL string) (*api.Order, error) {
	time.Sleep(1 * time.Second)

	return nil, ErrNotImplemented
}
