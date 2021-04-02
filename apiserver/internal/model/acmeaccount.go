package model

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/golang/protobuf/protoc-gen-go/generator"
	"github.com/google/uuid"
	"github.com/mennanov/fieldmask-utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"powerssl.dev/backend/auth"
	"powerssl.dev/sdk/apiserver/api"
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

func (a *ACMEAccount) UpdateWithMask(ctx context.Context, paths []string, acmeAccount *api.ACMEAccount) (_ map[string]interface{}, err error) {
	paths = a.sanitizeUpdateMask(paths, auth.IsInternal(ctx))
	var mask fieldmask_utils.Mask
	if mask, err = fieldmask_utils.MaskFromPaths(paths, generator.CamelCase); err != nil {
		return nil, err
	}
	if err = fieldmask_utils.StructToStruct(mask, acmeAccount, a); err != nil {
		return nil, err
	}
	clauses := make(map[string]interface{})
	for _, path := range paths {
		switch path {
		case "display_name":
			clauses[path] = a.DisplayName
		case "title":
			clauses[path] = a.Title
		case "description":
			clauses[path] = a.Description
		case "terms_of_service_agreed":
			clauses[path] = a.TermsOfServiceAgreed
		case "contacts":
			clauses[path] = a.Contacts
		case "account_url":
			clauses[path] = a.AccountURL
		}
	}
	return clauses, nil
}

func (a *ACMEAccount) sanitizeUpdateMask(paths []string, internal bool) []string {
	allowed := map[string]struct{}{
		"display_name": {},
		"title":        {},
		"description":  {},
		"contacts":     {},
	}
	if internal {
		allowed["terms_of_service_agreed"] = struct{}{}
		allowed["contacts"] = struct{}{}
		allowed["account_url"] = struct{}{}
	}
	n := 0
	for _, path := range paths {
		if _, ok := allowed[path]; ok {
			paths[n] = path
			n++
		}
	}
	return paths[:n]
}

type ACMEAccounts []*ACMEAccount

func (a ACMEAccounts) ToAPI() []*api.ACMEAccount {
	accounts := make([]*api.ACMEAccount, len(a))
	for i, account := range a {
		accounts[i] = account.ToAPI()
	}
	return accounts
}
