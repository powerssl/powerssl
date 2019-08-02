package pki

import (
	"github.com/cloudflare/cfssl/csr"
	"github.com/cloudflare/cfssl/initca"
)

func Init(keyAlgo string, keySize int) ([]byte, []byte, []byte, error) {
	req := csr.CertificateRequest{
		KeyRequest: &csr.KeyRequest{
			A: keyAlgo,
			S: keySize,
		},
		Names: []csr.Name{
			{
				O: "PowerSSL Root Authority",
			},
		},
	}

	cert, csr, key, err := initca.New(&req)
	if err != nil {
		return nil, nil, nil, err
	}

	return cert, csr, key, nil
}
