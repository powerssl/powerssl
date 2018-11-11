package acmeaccount

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gogo/status"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc/codes"

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
		Name: fmt.Sprintf("acme-accounts/%d", a.ID),

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

func FindACMEAccountByName(name string, db *gorm.DB) (*ACMEAccount, error) {
	s := strings.Split(name, "/")
	if len(s) != 2 {
		return nil, status.Error(codes.InvalidArgument, "malformed name")
	}
	id, err := strconv.Atoi(s[1])
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "malformed name")
	}

	acmeAccount := &ACMEAccount{}
	if db.Where("id = ?", id).First(&acmeAccount).RecordNotFound() {
		return nil, status.Error(codes.NotFound, "not found")
	}
	return acmeAccount, nil
}

func NewACMEAccountFromAPI(acmeAccount *api.ACMEAccount) *ACMEAccount {
	return &ACMEAccount{
		TermsOfServiceAgreed: acmeAccount.TermsOfServiceAgreed,
		Contacts:             strings.Join(acmeAccount.Contacts, ","),
		AccountURL:           acmeAccount.AccountURL,
		DirectoryURL:         acmeAccount.DirectoryURL,
		IntegrationName:      acmeAccount.IntegrationName,
	}
}
