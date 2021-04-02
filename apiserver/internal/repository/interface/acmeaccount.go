package _interface

import (
	"context"

	"powerssl.dev/apiserver/internal/model"
)

type ACMEAccountRepository interface {
	Delete(ctx context.Context, acmeAccounts ...*model.ACMEAccount) (err error)
	FindAll(ctx context.Context) (acmeAccounts model.ACMEAccounts, err error)
	FindAllByParent(ctx context.Context, parent string) (acmeAccounts model.ACMEAccounts, err error)
	FindOneByName(ctx context.Context, name string) (acmeAccount *model.ACMEAccount, err error)
	Insert(ctx context.Context, acmeAccounts ...*model.ACMEAccount) (err error)
	Update(ctx context.Context, acmeAccount *model.ACMEAccount, clauses map[string]interface{}) (err error)
}
