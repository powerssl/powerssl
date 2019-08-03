package user // import "powerssl.dev/powerssl/pkg/apiserver/user"

import (
	"context"

	"powerssl.dev/powerssl/pkg/apiserver/api"
)

type Service interface {
	Create(ctx context.Context, user *api.User) (*api.User, error)
	Delete(ctx context.Context, name string) error
	Get(ctx context.Context, name string) (*api.User, error)
	List(ctx context.Context, pageSize int, pageToken string) ([]*api.User, string, error)
	Update(ctx context.Context, name string, user *api.User) (*api.User, error)
}
