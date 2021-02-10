package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/mennanov/fieldmask-utils"

	"powerssl.dev/powerssl/pkg/apiserver/api"
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

func (a *ACMEServer) Update(updateMask []string, acmeServer *ACMEServer) error {
	mask := fieldmask_utils.MaskFromString(strings.Join(updateMask, ","))
	return fieldmask_utils.StructToStruct(mask, acmeServer, a)
}

type ACMEServers []*ACMEServer

func (a ACMEServers) ToAPI() []*api.ACMEServer {
	acmeServers := make([]*api.ACMEServer, len(a))
	for i, server := range a {
		acmeServers[i] = server.ToAPI()
	}
	return acmeServers
}
