package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	apiv1 "powerssl.dev/api/apiserver/v1"
	context2 "powerssl.dev/backend/context"
)

var acmeServerMessageType *apiv1.ACMEServer
var acmeServerUpdateMaskSanitizer = NewUpdateMaskSanitizer(acmeServerMessageType).
	Allowed("display_name").
	Internal("directory_url", "integration_name")

func (a *AcmeServer) Name() string {
	return fmt.Sprintf("acmeServers/%s", a.ID)
}

func (a *AcmeServer) ToAPI() *apiv1.ACMEServer {
	return &apiv1.ACMEServer{
		Name:            a.Name(),
		CreateTime:      timestamppb.New(a.CreatedAt),
		UpdateTime:      timestamppb.New(a.UpdatedAt),
		DisplayName:     a.DisplayName,
		DirectoryUrl:    a.DirectoryUrl,
		IntegrationName: a.IntegrationName,
	}
}

func (q *Queries) CreateACMEServerFromAPI(ctx context.Context, acmeServer *apiv1.ACMEServer) (AcmeServer, error) {
	return q.CreateACMEServer(ctx, CreateACMEServerParams{
		DisplayName:     acmeServer.DisplayName,
		DirectoryUrl:    acmeServer.DirectoryUrl,
		IntegrationName: acmeServer.IntegrationName,
	})
}

func (q *Queries) UpdateACMEServerWithMask(ctx context.Context, id uuid.UUID, fm *fieldmaskpb.FieldMask, acmeServer *apiv1.ACMEServer) (AcmeServer, error) {
	fm = acmeServerUpdateMaskSanitizer.Sanitize(fm, context2.IsInternal(ctx))
	updateParams := UpdateACMEServerParams{ID: id}
	if err := setUpdateParams(fm, acmeServer, &updateParams); err != nil {
		return AcmeServer{}, err
	}
	return q.UpdateACMEServer(ctx, updateParams)
}

type AcmeServers []AcmeServer

func (a AcmeServers) ToAPI() []*apiv1.ACMEServer {
	acmeServers := make([]*apiv1.ACMEServer, len(a))
	for i, server := range a {
		acmeServers[i] = server.ToAPI()
	}
	return acmeServers
}
