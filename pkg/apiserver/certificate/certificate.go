package certificate // import "powerssl.dev/powerssl/pkg/apiserver/certificate"

import (
	"context"

	"powerssl.dev/powerssl/pkg/apiserver/api"
)

type Service interface {
	Create(ctx context.Context, certificate *api.Certificate) (*api.Certificate, error)
	Delete(ctx context.Context, name string) error
	Get(ctx context.Context, name string) (*api.Certificate, error)
	List(ctx context.Context, pageSize int, pageToken string) ([]*api.Certificate, string, error)
	Update(ctx context.Context, name string, certificate *api.Certificate) (*api.Certificate, error)
}
