package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/gogo/status"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"

	"powerssl.dev/powerssl/pkg/apiserver/api"
)

// TODO: Doesn't get called
//func (a *ACMEAccount) Validate(db *pg.DB) (map[string]error, bool) {
//	var errors map[string]error
//
//	if !a.TermsOfServiceAgreed {
//		errors["TermsOfServiceAgreed"] = status.Error(codes.InvalidArgument, "terms of service need to be agreed")
//	}
//
//	if a.ACMEServer == nil {
//		errors["TermsOfServiceAgreed"] = status.Error(codes.NotFound, "ACME server not found")
//	}
//
//	return errors, len(errors) == 0
//}

type ACMEAccount struct {
	// Generic
	ID        string     `db:"id"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`

	// Custom
	DisplayName          string `db:"display_name"`
	Title                string `db:"title"`
	Description          string `db:"description"`
	TermsOfServiceAgreed bool   `db:"terms_of_service_agreed"`
	Contacts             string `db:"contacts"`
	AccountURL           string `db:"account_url"`

	// Relations
	ACMEServerID string `db:"acme_server_id"`
	ACMEServer   *ACMEServer
}

func NewACMEAccountFromAPI(parent string, acmeAccount *api.ACMEAccount, id string) (*ACMEAccount, error) {
	s := strings.Split(parent, "/")
	if len(s) != 2 || s[0] != "acmeServers" {
		return nil, status.Error(codes.InvalidArgument, "malformed parent")
	}

	if id == "" {
		id = uuid.New().String()
	}

	return &ACMEAccount{
		ID:                   id,
		ACMEServerID:         s[1],
		TermsOfServiceAgreed: acmeAccount.TermsOfServiceAgreed,
		Contacts:             strings.Join(acmeAccount.Contacts, ","),
		AccountURL:           acmeAccount.AccountURL,
		ACMEServer: &ACMEServer{
			ID: s[1],
		},
	}, nil
}

func (a *ACMEAccount) Name() string {
	return fmt.Sprintf("acmeServers/%s/acmeAccounts/%s", a.ACMEServerID, a.ID)
}

func (a *ACMEAccount) ToAPI() *api.ACMEAccount {
	var contacts []string
	if a.Contacts != "" {
		contacts = strings.Split(a.Contacts, ",")
	}
	return &api.ACMEAccount{
		Name:                 a.Name(),
		CreateTime:           a.CreatedAt,
		UpdateTime:           a.UpdatedAt,
		DisplayName:          a.DisplayName,
		Title:                a.Title,
		Description:          a.Description,
		Labels:               map[string]string{"not": "implemented"},
		TermsOfServiceAgreed: a.TermsOfServiceAgreed,
		Contacts:             contacts,
		AccountURL:           a.AccountURL,
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
