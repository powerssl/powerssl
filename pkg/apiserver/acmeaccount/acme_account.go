package acmeaccount

import (
	"fmt"
	"strings"
	"time"

	"github.com/gogo/status"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc/codes"

	"powerssl.io/pkg/apiserver/api"
)

type ACMEAccount struct {
	ID        string `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	DisplayName          string
	Title                string
	Description          string
	ACMEServer           string
	TermsOfServiceAgreed bool
	Contacts             string
	AccountURL           string

	ACMEServerID string
}

func (*ACMEAccount) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New().String())
	return nil
}

func (a *ACMEAccount) Name() string {
	return fmt.Sprintf("acme-accounts/%s", a.ID)
}

func (a *ACMEAccount) ToAPI() *api.ACMEAccount {
	return &api.ACMEAccount{
		Name: a.Name(),

		CreateTime:  a.CreatedAt,
		UpdateTime:  a.UpdatedAt,
		DisplayName: a.DisplayName,
		Title:       a.Title,
		Description: a.Description,
		Labels:      map[string]string{"not": "implemented"},

		ACMEServer:           a.ACMEServer,
		TermsOfServiceAgreed: a.TermsOfServiceAgreed,
		Contacts:             strings.Split(a.Contacts, ","),
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

func FindACMEAccountByName(name string, db *gorm.DB) (*ACMEAccount, error) {
	s := strings.Split(name, "/")
	if len(s) != 2 {
		return nil, status.Error(codes.InvalidArgument, "malformed name")
	}

	acmeAccount := &ACMEAccount{}
	if db.Where("id = ?", s[1]).First(&acmeAccount).RecordNotFound() {
		return nil, status.Error(codes.NotFound, "not found")
	}
	return acmeAccount, nil
}

func NewACMEAccountFromAPI(acmeAccount *api.ACMEAccount) *ACMEAccount {
	return &ACMEAccount{
		ACMEServer:           acmeAccount.ACMEServer,
		TermsOfServiceAgreed: acmeAccount.TermsOfServiceAgreed,
		Contacts:             strings.Join(acmeAccount.Contacts, ","),
		AccountURL:           acmeAccount.AccountURL,
	}
}
