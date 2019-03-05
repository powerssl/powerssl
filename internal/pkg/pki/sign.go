package pki

import (
	"github.com/cloudflare/cfssl/config"
	"github.com/cloudflare/cfssl/helpers"
	"github.com/cloudflare/cfssl/signer"
	"github.com/cloudflare/cfssl/signer/local"
)

func Sign(ca, caKey, csr string) ([]byte, error) {
	policy := &config.Signing{
		Default: &config.SigningProfile{
			Expiry: helpers.OneYear,
			CAConstraint: config.CAConstraint{
				IsCA:       true,
				MaxPathLen: 1,
			},
			Usage: []string{
				"digital signature",
				"signing",
				"key encipherment",
				"cert sign",
				"crl sign",
			},
		},
	}

	s, err := local.NewSignerFromFile(ca, caKey, policy)
	if err != nil {
		return nil, err
	}

	req := signer.SignRequest{Request: csr}
	cert, err := s.Sign(req)
	if err != nil {
		return nil, err
	}

	return cert, nil
}
