package pki

import (
	"github.com/cloudflare/cfssl/cli/genkey"
	"github.com/cloudflare/cfssl/config"
	"github.com/cloudflare/cfssl/csr"
	"github.com/cloudflare/cfssl/signer"
	"github.com/cloudflare/cfssl/signer/local"
)

func Gen(ca, caKey, hostname, keyAlgo string, keySize int) ([]byte, []byte, []byte, error) {
	req := csr.CertificateRequest{
		KeyRequest: &csr.BasicKeyRequest{
			A: keyAlgo,
			S: keySize,
		},
		Hosts: []string{hostname},
	}

	g := &csr.Generator{Validator: genkey.Validator}
	csr, key, err := g.ProcessRequest(&req)
	if err != nil {
		return nil, nil, nil, err
	}

	policy := &config.Signing{Default: config.DefaultConfig()}

	s, err := local.NewSignerFromFile(ca, caKey, policy)
	if err != nil {
		return nil, nil, nil, err
	}

	signReq := signer.SignRequest{Request: string(csr)}
	cert, err := s.Sign(signReq)
	if err != nil {
		return nil, nil, nil, err
	}

	return cert, csr, key, nil
}
