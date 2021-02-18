package pki

import (
	cfsslcligenkey "github.com/cloudflare/cfssl/cli/genkey"
	cfsslconfig "github.com/cloudflare/cfssl/config"
	cfsslcsr "github.com/cloudflare/cfssl/csr"
	cfsslsigner "github.com/cloudflare/cfssl/signer"
	cfsslsignerlocal "github.com/cloudflare/cfssl/signer/local"
)

func Gen(ca, caKey, hostname, keyAlgo string, keySize int) ([]byte, []byte, []byte, error) {
	req := cfsslcsr.CertificateRequest{
		KeyRequest: &cfsslcsr.KeyRequest{
			A: keyAlgo,
			S: keySize,
		},
		Hosts: []string{hostname},
	}

	g := &cfsslcsr.Generator{Validator: cfsslcligenkey.Validator}
	csr, key, err := g.ProcessRequest(&req)
	if err != nil {
		return nil, nil, nil, err
	}

	policy := &cfsslconfig.Signing{Default: cfsslconfig.DefaultConfig()}

	s, err := cfsslsignerlocal.NewSignerFromFile(ca, caKey, policy)
	if err != nil {
		return nil, nil, nil, err
	}

	signReq := cfsslsigner.SignRequest{Request: string(csr)}
	cert, err := s.Sign(signReq)
	if err != nil {
		return nil, nil, nil, err
	}

	return cert, csr, key, nil
}
