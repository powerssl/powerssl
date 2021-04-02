package acmeserver // import "powerssl.dev/sdk/apiserver/acmeserver"

import (
	"context"

	"powerssl.dev/sdk/apiserver/api"
)

type Service interface {
	Create(ctx context.Context, acmeServer *api.ACMEServer) (*api.ACMEServer, error)
	Delete(ctx context.Context, name string) error
	Get(ctx context.Context, name string) (*api.ACMEServer, error)
	List(ctx context.Context, pageSize int, pageToken string) ([]*api.ACMEServer, string, error)
	Update(ctx context.Context, name string, updateMask []string, acmeServer *api.ACMEServer) (*api.ACMEServer, error)
}
