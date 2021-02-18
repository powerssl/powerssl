package pki

import (
	cfsslcsr "github.com/cloudflare/cfssl/csr"
	cfsslinitca "github.com/cloudflare/cfssl/initca"
)

func Init(keyAlgo string, keySize int) ([]byte, []byte, []byte, error) {
	req := cfsslcsr.CertificateRequest{
		KeyRequest: &cfsslcsr.KeyRequest{
			A: keyAlgo,
			S: keySize,
		},
		Names: []cfsslcsr.Name{
			{
				O: "PowerSSL Root Authority",
			},
		},
	}

	cert, csr, key, err := cfsslinitca.New(&req)
	if err != nil {
		return nil, nil, nil, err
	}

	return cert, csr, key, nil
}
