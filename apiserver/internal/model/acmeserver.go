package model

import (
	"context"
	"fmt"
	"time"

	"github.com/golang/protobuf/protoc-gen-go/generator"
	"github.com/google/uuid"
	"github.com/mennanov/fieldmask-utils"

	"powerssl.dev/backend/auth"
	"powerssl.dev/sdk/apiserver/api"
)

type ACMEServer struct {
	// Generic
	ID        string     `db:"id"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`

	// Custom
	DisplayName     string `db:"display_name"`
	DirectoryURL    string `db:"directory_url"`
	IntegrationName string `db:"integration_name"`

	// Relations
	ACMEAccounts []ACMEAccount
}

func NewACMEServerFromAPI(apiACMEServer *api.ACMEServer, id string) *ACMEServer {
	if id == "" {
		id = uuid.New().String()
	}

	return &ACMEServer{
		ID:              id,
		DisplayName:     apiACMEServer.DisplayName,
		DirectoryURL:    apiACMEServer.DirectoryURL,
		IntegrationName: apiACMEServer.IntegrationName,
	}
}

func (a *ACMEServer) Name() string {
	return fmt.Sprintf("acmeServers/%s", a.ID)
}

func (a *ACMEServer) ToAPI() *api.ACMEServer {
	return &api.ACMEServer{
		Name:            a.Name(),
		CreateTime:      a.CreatedAt,
		UpdateTime:      a.UpdatedAt,
		DisplayName:     a.DisplayName,
		DirectoryURL:    a.DirectoryURL,
		IntegrationName: a.IntegrationName,
	}
}

func (a *ACMEServer) UpdateWithMask(ctx context.Context, paths []string, acmeServer *api.ACMEServer) (_ map[string]interface{}, err error) {
	paths = a.sanitizeUpdateMask(paths, auth.IsInternal(ctx))
	var mask fieldmask_utils.Mask
	if mask, err = fieldmask_utils.MaskFromPaths(paths, generator.CamelCase); err != nil {
		return nil, err
	}
	if err = fieldmask_utils.StructToStruct(mask, acmeServer, a); err != nil {
		return nil, err
	}
	clauses := make(map[string]interface{})
	for _, path := range paths {
		switch path {
		case "display_name":
			clauses[path] = a.DisplayName
		case "directory_url":
			clauses[path] = a.DirectoryURL
		case "integration_name":
			clauses[path] = a.IntegrationName
		}
	}
	return clauses, nil
}

func (a *ACMEServer) sanitizeUpdateMask(paths []string, internal bool) []string {
	allowed := map[string]struct{}{
		"display_name": {},
	}
	if internal {
		allowed["directory_url"] = struct{}{}
		allowed["integration_name"] = struct{}{}
	}
	n := 0
	for _, path := range paths {
		if _, ok := allowed[path]; ok {
			paths[n] = path
			n++
		}
	}
	return paths[:n]
}

type ACMEServers []*ACMEServer

func (a ACMEServers) ToAPI() []*api.ACMEServer {
	acmeServers := make([]*api.ACMEServer, len(a))
	for i, server := range a {
		acmeServers[i] = server.ToAPI()
	}
	return acmeServers
}
