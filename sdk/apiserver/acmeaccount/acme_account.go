package acmeaccount // import "powerssl.dev/sdk/apiserver/acmeaccount"

import (
	"context"

	"powerssl.dev/sdk/apiserver/api"
)

type Service interface {
	Create(ctx context.Context, parent string, acmeAccount *api.ACMEAccount) (*api.ACMEAccount, error)
	Delete(ctx context.Context, name string) error
	Get(ctx context.Context, name string) (*api.ACMEAccount, error)
	List(ctx context.Context, parent string, pageSize int, pageToken string) ([]*api.ACMEAccount, string, error)
	Update(ctx context.Context, name string, updateMask []string, acmeAccount *api.ACMEAccount) (*api.ACMEAccount, error)
}
