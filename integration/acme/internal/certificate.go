package acme

import (
	"context"
	"crypto/x509"
	"time"

	apiv1 "powerssl.dev/api/controller/v1"
)

func (acme *ACME) GetCertificate(_ context.Context, accountURL string, certificateURL string) ([]*x509.Certificate, error) {
	time.Sleep(1 * time.Second)

	return nil, ErrNotImplemented
}

func (acme *ACME) RevokeCertificate(_ context.Context, directoryURL string, accountURL string, certificate *x509.Certificate, reason *apiv1.RevocationReason) error {
	time.Sleep(1 * time.Second)

	return ErrNotImplemented
}
