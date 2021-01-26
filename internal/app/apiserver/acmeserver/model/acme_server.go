package model

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"

	"powerssl.dev/powerssl/internal/pkg/uid"
	"powerssl.dev/powerssl/pkg/apiserver/api"
)

type ACMEServer struct {
	ID        string `pg:",pk"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `pg:",soft_delete"`

	DisplayName     string
	DirectoryURL    string
	IntegrationName string
}

var _ pg.BeforeInsertHook = (*ACMEServer)(nil)

func (acmeServer *ACMEServer) BeforeInsert(ctx context.Context) (context.Context, error) {
	acmeServer.ID = uid.New()
	return ctx, nil
}

func (a *ACMEServer) Name() string {
	return fmt.Sprintf("acmeServers/%s", a.ID)
}

func (a *ACMEServer) ToAPI() *api.ACMEServer {
	return &api.ACMEServer{
		Name: a.Name(),

		CreateTime:  a.CreatedAt,
		UpdateTime:  a.UpdatedAt,
		DisplayName: a.DisplayName,

		DirectoryURL:    a.DirectoryURL,
		IntegrationName: a.IntegrationName,
	}
}

type ACMEServers []*ACMEServer

func (a ACMEServers) ToAPI() []*api.ACMEServer {
	servers := make([]*api.ACMEServer, len(a))
	for i, server := range a {
		servers[i] = server.ToAPI()
	}
	return servers
}

func FindACMEServerByName(name string, db *pg.DB) (*ACMEServer, error) {
	s := strings.Split(name, "/")
	if len(s) != 2 {
		return nil, status.Error(codes.InvalidArgument, "malformed name")
	}

	acmeServer := &ACMEServer{}
	if err := db.Model(acmeServer).Where("id = ?", s[1]).Limit(1).Select(); err != nil {
		if err == pg.ErrNoRows {
			return nil, status.Error(codes.NotFound, "not found")
		}
		return nil, err
	}
	return acmeServer, nil
}

func NewACMEServerFromAPI(acmeServer *api.ACMEServer) *ACMEServer {
	return &ACMEServer{
		DisplayName:     acmeServer.DisplayName,
		DirectoryURL:    acmeServer.DirectoryURL,
		IntegrationName: acmeServer.IntegrationName,
	}
}
