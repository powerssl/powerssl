package acme

import (
	"crypto/x509"

	"powerssl.io/pkg/controller/api"
)

func (acme *ACME) GetCertificate(accountURL string, certificateURL string) ([]*x509.Certificate, error) {
	return nil, ErrNotImplemented
}

func (acme *ACME) RevokeCertificate(directoryURL string, accountURL string, certificate *x509.Certificate, reason *api.RevocationReason) error {
	return ErrNotImplemented
}
