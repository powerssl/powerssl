package grpcserver

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"time"

	"github.com/johanbrandhorst/certify"
	certifyvault "github.com/johanbrandhorst/certify/issuers/vault"
	"go.uber.org/zap"
	zapadapter "logur.dev/adapter/zap"

	"powerssl.dev/backend/vault"
)

type keyGeneratorFunc func() (crypto.PrivateKey, error)

func (kgf keyGeneratorFunc) Generate() (crypto.PrivateKey, error) {
	return kgf()
}

func makeCertify(cfg Config, logger *zap.SugaredLogger) (*tls.Config, error) {
	client, err := vault.New(cfg.Vault)
	if err != nil {
		return nil, err
	}

	c := &certify.Certify{
		CommonName:  cfg.CommonName,
		Issuer:      certifyvault.FromClient(client.Client(), cfg.VaultRole),
		RenewBefore: time.Hour,
		Cache:       certify.NewMemCache(),
		CertConfig: &certify.CertConfig{
			SubjectAlternativeNames: []string{cfg.CommonName},
			KeyGenerator: keyGeneratorFunc(func() (crypto.PrivateKey, error) {
				return rsa.GenerateKey(rand.Reader, 2048)
			}),
		},
		IssueTimeout: time.Minute,
		Logger:       certifyLogger(logger),
	}
	getCertificate := func(hello *tls.ClientHelloInfo) (cert *tls.Certificate, err error) {
		// TODO: ???
		hello.ServerName = cfg.CommonName
		if cert, err = c.GetCertificate(hello); err != nil {
			logger.Error(err)
		}
		return cert, err
	}
	// TODO: This was priming the server before the first request. Certify is now failing with this.
	// if _, err := getCertificate(&tls.ClientHelloInfo{ServerName: cfg.CommonName}); err != nil {
	// 	return err
	// }
	return &tls.Config{GetCertificate: getCertificate}, nil
}

func certifyLogger(logger *zap.SugaredLogger) certify.Logger {
	return zapadapter.New(logger.Desugar())
}
