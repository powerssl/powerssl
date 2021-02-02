package _interface

import (
	"context"

	"powerssl.dev/powerssl/internal/app/apiserver/model"
)

type ACMEAccountRepository interface {
	// Generic
	Add(ctx context.Context, acmeAccount *model.ACMEAccount) (err error)
	AddRange(ctx context.Context, acmeAccounts *model.ACMEAccounts) (err error)
	Find(ctx context.Context, predicate string) (acmeAccount *model.ACMEAccount, err error)
	Get(ctx context.Context, id string) (acmeAccount *model.ACMEAccount, err error)
	GetAll(ctx context.Context) (acmeAccounts *model.ACMEAccounts, err error)
	Remove(ctx context.Context, acmeAccount *model.ACMEAccount) (err error)
	RemoveRange(ctx context.Context, acmeAccounts *model.ACMEAccounts) (err error)

	// Custom
	FindByName(ctx context.Context, name string) (acmeAccount *model.ACMEAccount, err error)
}
