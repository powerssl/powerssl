package acme

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"time"

	"powerssl.io/powerssl/pkg/controller/api"
)

func (acme *ACME) CreateAccount(ctx context.Context, directoryURL string, termsOfServiceAgreed bool, contacts []string) (*api.Account, error) {
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("error creating private key: %v", err)
	}
	client, err := NewClient(ctx, directoryURL)
	if err != nil {
		return nil, err
	}
	account, err := client.NewAccount(privKey, false, termsOfServiceAgreed, contacts...)
	if err != nil {
		return nil, fmt.Errorf("error creating new account: %v", err)
	}
	return &api.Account{
		Contacts:             account.Contact,
		Status:               api.AccountStatusRevoked, // TODO
		TermsOfServiceAgreed: account.TermsOfServiceAgreed,
		URL:                  account.URL,
	}, nil
}

func (acme *ACME) DeactivateAccount(_ context.Context, accountURL string) (*api.Account, error) {
	time.Sleep(1 * time.Second)

	return nil, ErrNotImplemented
}

func (acme *ACME) RekeyAccount(_ context.Context, accountURL string, directoryURL string) (*api.Account, error) {
	time.Sleep(1 * time.Second)

	return nil, ErrNotImplemented
}

func (acme *ACME) UpdateAccount(_ context.Context, accountURL string, contacts []string) (*api.Account, error) {
	time.Sleep(1 * time.Second)

	return nil, ErrNotImplemented
}
