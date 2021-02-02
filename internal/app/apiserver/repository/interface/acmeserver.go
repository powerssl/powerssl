package _interface

import (
	"context"

	"powerssl.dev/powerssl/internal/app/apiserver/model"
)

type ACMEServerRepository interface {
	// Generic
	Add(ctx context.Context, acmeServer *model.ACMEServer) (err error)
	AddRange(ctx context.Context, acmeServers *model.ACMEServers) (err error)
	Find(ctx context.Context, predicate string) (acmeServer *model.ACMEServer, err error)
	Get(ctx context.Context, id string) (acmeServer *model.ACMEServer, err error)
	GetAll(ctx context.Context) (acmeServers *model.ACMEServers, err error)
	GetRange(ctx context.Context, pageSize int, pageToken string) (acmeServers *model.ACMEServers, nextPageToken string, err error)
	Remove(ctx context.Context, acmeServer *model.ACMEServer) (err error)
	RemoveRange(ctx context.Context, acmeServers *model.ACMEServers) (err error)

	// Custom
	FindByName(ctx context.Context, name string) (acmeServer *model.ACMEServer, err error)
}
