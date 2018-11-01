package acme

import (
	"crypto/x509"
	"time"

	"powerssl.io/pkg/controller/api"
)

func (acme *ACME) GetCertificate(accountURL string, certificateURL string) ([]*x509.Certificate, error) {
	time.Sleep(1 * time.Second)

	return nil, ErrNotImplemented
}

func (acme *ACME) RevokeCertificate(directoryURL string, accountURL string, certificate *x509.Certificate, reason *api.RevocationReason) error {
	time.Sleep(1 * time.Second)

	return ErrNotImplemented
}
