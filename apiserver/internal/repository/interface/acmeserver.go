package _interface

import (
	"context"

	"powerssl.dev/apiserver/internal/model"
)

type ACMEServerRepository interface {
	Delete(ctx context.Context, acmeServers ...*model.ACMEServer) (err error)
	FindAll(ctx context.Context, pageSize int, pageToken string) (acmeServers model.ACMEServers, nextPageToken string, err error)
	FindOneByName(ctx context.Context, name string) (acmeServer *model.ACMEServer, err error)
	Insert(ctx context.Context, acmeServers ...*model.ACMEServer) (err error)
	Update(ctx context.Context, acmeServer *model.ACMEServer, clauses map[string]interface{}) (err error)
}
