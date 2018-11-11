package acmeaccount

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"

	"powerssl.io/pkg/apiserver/api"
)

type ACMEAccount struct {
	gorm.Model

	DisplayName          string
	Title                string
	Description          string
	TermsOfServiceAgreed bool
	Contacts             string
	AccountURL           string
	DirectoryURL         string
	IntegrationName      string
}

func (a *ACMEAccount) ToAPI() *api.ACMEAccount {
	return &api.ACMEAccount{
		Name: fmt.Sprint("acme-accounts/", a.ID),

		CreateTime:  a.CreatedAt,
		UpdateTime:  a.UpdatedAt,
		DisplayName: a.DisplayName,
		Title:       a.Title,
		Description: a.Description,
		Labels:      map[string]string{"not": "implemented"},

		TermsOfServiceAgreed: a.TermsOfServiceAgreed,
		Contacts:             strings.Split(a.Contacts, ","),
		AccountURL:           a.AccountURL,
		DirectoryURL:         a.DirectoryURL,
		IntegrationName:      a.IntegrationName,
	}
}

type ACMEAccounts []*ACMEAccount

func (a ACMEAccounts) ToAPI() []*api.ACMEAccount {
	accounts := make([]*api.ACMEAccount, len(a))
	for i, account := range a {
		accounts[i] = account.ToAPI()
	}
	return accounts
}
