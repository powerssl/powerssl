package acme

import (
	"powerssl.io/pkg/controller/api"
	"powerssl.io/pkg/integration"
)

const Name = "ACME"

type ACME struct{}

func New() integration.CAIntegration {
	return &ACME{}
}

func (acme *ACME) GetName() string {
	return Name
}

func (acme *ACME) AuthorizeDomain(domain string) (challenges []*api.Challenge, err error) {
	return challenges, err
}

func (acme *ACME) RequestCertificate(csr string) (cert string, err error) {
	return cert, err
}

func (acme *ACME) RevokeCertificate(cert string) (err error) {
	return err
}

func (acme *ACME) VerifyDomain(domain string, challengeType api.ChallengeType) (err error) {
	return err
}
