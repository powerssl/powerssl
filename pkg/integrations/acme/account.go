package acme

import (
	"fmt"
	"time"

	"powerssl.io/pkg/controller/api"
)

func (acme *ACME) CreateAccount(directoryURL string, termsOfServiceAgreed bool, contacts []string) (*api.Account, error) {
	time.Sleep(1 * time.Second)

	fmt.Printf("directoryURL: %#v\n", directoryURL)
	fmt.Printf("termsOfServiceAgreed: %#v\n", termsOfServiceAgreed)
	fmt.Printf("contacts: %#v\n", contacts)
	// newAccountReq := struct {
	// 	TermsOfServiceAgreed bool     `json:"termsOfServiceAgreed"`
	// 	Contacts             []string `json:"contact"`
	// }{
	// 	TermsOfServiceAgreed: termsOfServiceAgreed,
	// 	Contacts:             contacts,
	// }
	account := &api.Account{
		Contacts:             []string{"mailto:bert@example.com"},
		Status:               api.AccountStatusRevoked,
		TermsOfServiceAgreed: true,
		URL:                  "https://example.com/acct/123",
	}
	return account, nil
}

func (acme *ACME) DeactivateAccount(accountURL string) (*api.Account, error) {
	time.Sleep(1 * time.Second)

	return nil, ErrNotImplemented
}

func (acme *ACME) RekeyAccount(accountURL string, directoryURL string) (*api.Account, error) {
	time.Sleep(1 * time.Second)

	return nil, ErrNotImplemented
}

func (acme *ACME) UpdateAccount(accountURL string, contacts []string) (*api.Account, error) {
	time.Sleep(1 * time.Second)

	return nil, ErrNotImplemented
}
