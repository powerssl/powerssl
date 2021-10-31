package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"powerssl.dev/backend/auth"
	"powerssl.dev/sdk/apiserver/api"
)

var acmeServerUpdateMaskSanitizer = NewUpdateMaskSanitizer().
	Allowed("display_name").
	Internal("directory_url", "integration_name")

func (a *AcmeServer) Name() string {
	return fmt.Sprintf("acmeServers/%s", a.ID)
}

func (a *AcmeServer) ToAPI() *api.ACMEServer {
	return &api.ACMEServer{
		Name:            a.Name(),
		CreateTime:      a.CreatedAt,
		UpdateTime:      a.UpdatedAt,
		DisplayName:     a.DisplayName,
		DirectoryURL:    a.DirectoryUrl,
		IntegrationName: a.IntegrationName,
	}
}

func (q *Queries) CreateACMEServerFromAPI(ctx context.Context, acmeServer *api.ACMEServer) (AcmeServer, error) {
	return q.CreateACMEServer(ctx, CreateACMEServerParams{
		DisplayName:     acmeServer.DisplayName,
		DirectoryUrl:    acmeServer.DirectoryURL,
		IntegrationName: acmeServer.IntegrationName,
	})
}

func (q *Queries) UpdateACMEServerWithMask(ctx context.Context, id uuid.UUID, paths []string, acmeServer *api.ACMEServer) (AcmeServer, error) {
	paths = acmeServerUpdateMaskSanitizer.Sanitize(paths, auth.IsInternal(ctx))
	updateParams := UpdateACMEServerParams{ID: id}
	if err := setUpdateParams(paths, acmeServer, &updateParams); err != nil {
		return AcmeServer{}, err
	}
	return q.UpdateACMEServer(ctx, updateParams)
}

type AcmeServers []AcmeServer

func (a AcmeServers) ToAPI() []*api.ACMEServer {
	acmeServers := make([]*api.ACMEServer, len(a))
	for i, server := range a {
		acmeServers[i] = server.ToAPI()
	}
	return acmeServers
}
