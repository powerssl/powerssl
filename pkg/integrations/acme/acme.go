package acme

import (
	"crypto/x509"

	"powerssl.io/pkg/controller/api"
	integrationacme "powerssl.io/pkg/integration/acme"
)

type errorConst string

func (e errorConst) Error() string { return string(e) }

const ErrNotImplemented = errorConst("not implemented.")

type ACME struct{}

func New() integrationacme.Integration {
	return &ACME{}
}

func (acme *ACME) CreateAccount(directoryURL string, termsOfServiceAgreed bool, contacts []string) (*api.Account, error) {
	return nil, ErrNotImplemented
}

func (acme *ACME) DeactivateAccount(accountURL string) (*api.Account, error) {
	return nil, ErrNotImplemented
}

func (acme *ACME) RekeyAccount(accountURL string, directoryURL string) (*api.Account, error) {
	return nil, ErrNotImplemented
}

func (acme *ACME) UpdateAccount(accountURL string, contacts []string) (*api.Account, error) {
	return nil, ErrNotImplemented
}

func (acme *ACME) CreateOrder(directoryURL string, accountURL string, identifiers []*api.Identifier, notBefore, notAfter string) (*api.Order, error) {
	return nil, ErrNotImplemented
}

func (acme *ACME) FinalizeOrder(accountURL string, orderURL string, certificateSigningRequest *x509.CertificateRequest) (*api.Order, error) {
	return nil, ErrNotImplemented
}

func (acme *ACME) GetOrder(accountURL string, orderURL string) (*api.Order, error) {
	return nil, ErrNotImplemented
}

func (acme *ACME) CreateAuthorization(directoryURL string, accountURL string, identifier *api.Identifier) (*api.Authorization, error) {
	return nil, ErrNotImplemented
}

func (acme *ACME) DeactivateAuthorization(accountURL string, authorizationURL string) (*api.Authorization, error) {
	return nil, ErrNotImplemented
}

func (acme *ACME) GetAuthorization(accountURL string, authorizationURL string) (*api.Authorization, error) {
	return nil, ErrNotImplemented
}

func (acme *ACME) GetChallenge(accountURL string, challengeURL string) (*api.Challenge, error) {
	return nil, ErrNotImplemented
}

func (acme *ACME) ValidateChallenge(accountURL string, challengeURL string) (*api.Challenge, error) {
	return nil, ErrNotImplemented
}

func (acme *ACME) GetCertificate(accountURL string, certificateURL string) ([]*x509.Certificate, error) {
	return nil, ErrNotImplemented
}

func (acme *ACME) RevokeCertificate(directoryURL string, accountURL string, certificate *x509.Certificate, reason *api.RevocationReason) error {
	return ErrNotImplemented
}
