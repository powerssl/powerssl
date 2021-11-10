package acme

import (
	"context"
	"fmt"
	"time"

	apiv1 "powerssl.dev/api/controller/v1"
)

func (acme *ACME) CreateAccount(ctx context.Context, keyToken, directoryURL string, termsOfServiceAgreed bool, contacts []string) (*apiv1.Account, error) {
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
	return &apiv1.Account{
		Contacts:             account.Contact,
		Status:               apiv1.Account_REVOKED, // TODO
		TermsOfServiceAgreed: account.TermsOfServiceAgreed,
		Url:                  account.URL,
	}, nil
}

func (acme *ACME) DeactivateAccount(_ context.Context, accountURL string) (*apiv1.Account, error) {
	time.Sleep(1 * time.Second)

	return nil, ErrNotImplemented
}

func (acme *ACME) RekeyAccount(_ context.Context, accountURL string, directoryURL string) (*apiv1.Account, error) {
	time.Sleep(1 * time.Second)

	return nil, ErrNotImplemented
}

func (acme *ACME) UpdateAccount(_ context.Context, accountURL string, contacts []string) (*apiv1.Account, error) {
	time.Sleep(1 * time.Second)

	return nil, ErrNotImplemented
}
