package acme

import (
	"context"
	"fmt"
	"time"

	"powerssl.dev/sdk/controller/api"
)

func (acme *ACME) CreateAccount(ctx context.Context, keyToken, directoryURL string, termsOfServiceAgreed bool, contacts []string) (*api.Account, error) {
	client, err := NewClient(ctx, directoryURL)
	if err != nil {
		return nil, err
	}
	signer, err := acme.vault.SignerFromKeyToken(keyToken)
	if err != nil {
		return nil, err
	}
	account, err := client.NewAccount(signer, false, termsOfServiceAgreed, contacts...)
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
