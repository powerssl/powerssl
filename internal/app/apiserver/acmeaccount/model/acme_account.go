package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/gogo/status"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc/codes"

	acmeserver "powerssl.dev/powerssl/internal/app/apiserver/acmeserver/model"
	"powerssl.dev/powerssl/internal/pkg/uid"
	"powerssl.dev/powerssl/pkg/apiserver/api"
)

type ACMEAccount struct {
	ID        string `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	DisplayName          string
	Title                string
	Description          string
	ACMEServerID         string
	TermsOfServiceAgreed bool
	Contacts             string
	AccountURL           string
}

func (a *ACMEAccount) ACMEServer(db *gorm.DB, s string) (*acmeserver.ACMEServer, error) {
	acmeServer := &acmeserver.ACMEServer{}
	if db.Model(a).Select(s).Related(&acmeServer).RecordNotFound() {
		return nil, fmt.Errorf("ACME server not found")
	}
	return acmeServer, nil
}

func (*ACMEAccount) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uid.New())
	return nil
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

func (a *ACMEAccount) Validate(db *gorm.DB) {
	if !a.TermsOfServiceAgreed {
		db.AddError(status.Error(codes.InvalidArgument, "terms of service need to be agreed"))
	}

	if _, err := a.ACMEServer(db, "id"); err != nil {
		db.AddError(err)
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
	if len(s) != 4 || s[0] != "acmeServers" || s[2] != "acmeAccounts" {
		return nil, status.Error(codes.InvalidArgument, "malformed name")
	}
	acmeServerID, acmeAccountID := s[1], s[3]

	q := db.Where("id = ?", acmeAccountID)
	if s[1] != "-" {
		acmeServer := &acmeserver.ACMEServer{}
		if db.Select("id").Where("id = ?", acmeServerID).First(&acmeServer).RecordNotFound() {
			return nil, status.Error(codes.NotFound, "parent not found")
		}
		q = q.Where("acme_server_id = ?", acmeServer.ID)
	}

	acmeAccount := &ACMEAccount{}
	if q.First(&acmeAccount).RecordNotFound() {
		return nil, status.Error(codes.NotFound, "not found")
	}
	return acmeAccount, nil
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
