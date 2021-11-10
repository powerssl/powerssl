package acme

import (
	"context"
	"crypto/x509"
	"time"

	apiv1 "powerssl.dev/api/controller/v1"
)

func (acme *ACME) CreateOrder(_ context.Context, directoryURL string, accountURL string, identifiers []*apiv1.Identifier, notBefore, notAfter string) (*apiv1.Order, error) {
	time.Sleep(1 * time.Second)

	return nil, ErrNotImplemented
}

func (acme *ACME) FinalizeOrder(_ context.Context, accountURL string, orderURL string, certificateSigningRequest *x509.CertificateRequest) (*apiv1.Order, error) {
	time.Sleep(1 * time.Second)

	return nil, ErrNotImplemented
}

func (acme *ACME) GetOrder(_ context.Context, accountURL string, orderURL string) (*apiv1.Order, error) {
	time.Sleep(1 * time.Second)

	return nil, ErrNotImplemented
}
