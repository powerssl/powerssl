package model

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"

	acmeserver "powerssl.dev/powerssl/internal/app/apiserver/acmeserver/model"
	"powerssl.dev/powerssl/internal/pkg/uid"
	"powerssl.dev/powerssl/pkg/apiserver/api"
)

type ACMEAccount struct {
	ID        string `pg:",pk"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `pg:",soft_delete"`

	DisplayName          string
	Title                string
	Description          string
	ACMEServerID         string
	TermsOfServiceAgreed bool
	Contacts             string
	AccountURL           string

	ACMEServer *acmeserver.ACMEServer `pg:"rel:has-one"`
}

var _ pg.BeforeInsertHook = (*ACMEAccount)(nil)

func (acmeAccount *ACMEAccount) BeforeInsert(ctx context.Context) (context.Context, error) {
	acmeAccount.ID = uid.New()
	return ctx, nil
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
		Name: a.Name(),

		CreateTime:  a.CreatedAt,
		UpdateTime:  a.UpdatedAt,
		DisplayName: a.DisplayName,
		Title:       a.Title,
		Description: a.Description,
		Labels:      map[string]string{"not": "implemented"},

		TermsOfServiceAgreed: a.TermsOfServiceAgreed,
		Contacts:             contacts,
		AccountURL:           a.AccountURL,
	}
}

// TODO: Find better way
// TODO: Doesn't get called
func (a *ACMEAccount) Validate(db *pg.DB) (map[string]error, bool) {
	var errors map[string]error

	if !a.TermsOfServiceAgreed {
		errors["TermsOfServiceAgreed"] = status.Error(codes.InvalidArgument, "terms of service need to be agreed")
	}

	if a.ACMEServer == nil {
		errors["TermsOfServiceAgreed"] = status.Error(codes.NotFound, "ACME server not found")
	}

	return errors, len(errors) == 0
}

type ACMEAccounts []*ACMEAccount

func (a ACMEAccounts) ToAPI() []*api.ACMEAccount {
	accounts := make([]*api.ACMEAccount, len(a))
	for i, account := range a {
		accounts[i] = account.ToAPI()
	}
	return accounts
}

func FindACMEAccountByName(name string, db *pg.DB) (*ACMEAccount, error) {
	s := strings.Split(name, "/")
	if len(s) != 4 || s[0] != "acmeServers" || s[2] != "acmeAccounts" {
		return nil, status.Error(codes.InvalidArgument, "malformed name")
	}
	acmeServerID, acmeAccountID := s[1], s[3]

	var acmeAccount ACMEAccount
	q := db.Model(&acmeAccount).Where("id = ?", acmeAccountID).Limit(1)
	if s[1] != "-" {
		var acmeServer acmeserver.ACMEServer
		if err := db.Model(&acmeServer).Column("id").Where("id = ?", acmeServerID).Limit(1).Select(); err != nil {
			if err == pg.ErrNoRows {
				return nil, status.Error(codes.NotFound, "parent not found")
			}
			return nil, err
		}
		q = q.Where("acme_server_id = ?", acmeServer.ID)
	}

	if err := q.Select(); err != nil {
		if err == pg.ErrNoRows {
			return nil, status.Error(codes.NotFound, "not found")
		}
		return nil, err
	}
	return &acmeAccount, nil
}

func NewACMEAccountFromAPI(parent string, acmeAccount *api.ACMEAccount) (*ACMEAccount, error) {
	s := strings.Split(parent, "/")
	if len(s) != 2 || s[0] != "acmeServers" {
		return nil, status.Error(codes.InvalidArgument, "malformed parent")
	}
	acmeServerID := s[1]

	return &ACMEAccount{
		ACMEServerID:         acmeServerID,
		TermsOfServiceAgreed: acmeAccount.TermsOfServiceAgreed,
		Contacts:             strings.Join(acmeAccount.Contacts, ","),
		AccountURL:           acmeAccount.AccountURL,
	}, nil
}
