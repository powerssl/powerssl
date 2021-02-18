package pki

import (
	cfsslconfig "github.com/cloudflare/cfssl/config"
	cfsslhelpers "github.com/cloudflare/cfssl/helpers"
	cfsslsigner "github.com/cloudflare/cfssl/signer"
	cfsslsignerlocal "github.com/cloudflare/cfssl/signer/local"
)

func Sign(ca, caKey, csr string) ([]byte, error) {
	policy := &cfsslconfig.Signing{
		Default: &cfsslconfig.SigningProfile{
			Expiry: cfsslhelpers.OneYear,
			CAConstraint: cfsslconfig.CAConstraint{
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

	s, err := cfsslsignerlocal.NewSignerFromFile(ca, caKey, policy)
	if err != nil {
		return nil, err
	}

	req := cfsslsigner.SignRequest{Request: csr}
	cert, err := s.Sign(req)
	if err != nil {
		return nil, err
	}

	return cert, nil
}
