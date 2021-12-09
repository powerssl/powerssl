package model

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	apiv1 "powerssl.dev/api/apiserver/v1"
	"powerssl.dev/apiserver/internal/repository"
)

var acmeServerMessageType *apiv1.ACMEServer
var acmeServerUpdateMaskSanitizer = NewUpdateMaskSanitizer(acmeServerMessageType).
	Allowed("display_name").
	Internal("directory_url", "integration_name")

func AcmeServerUpdateParams(ctx context.Context, id uuid.UUID, fm *fieldmaskpb.FieldMask, acmeServer *apiv1.ACMEServer) (updateParams repository.UpdateACMEServerParams, err error) {
	updateParams.ID = id
	err = setUpdateParams(acmeServerUpdateMaskSanitizer.Sanitize(ctx, fm), acmeServer, &updateParams)
	return updateParams, err
}

func ParseAcmeServerID(name string) (id uuid.UUID, err error) {
	n := strings.Split(name, "/")
	if len(n) != 2 || n[0] != "acmeServers" {
		return id, fmt.Errorf("acme server name format mismatch")
	}
	return uuid.Parse(n[1])
}

type AcmeServer struct {
	repository.AcmeServer
}

func NewAcmeServer(acmeServer repository.AcmeServer) AcmeServer {
	return AcmeServer{
		AcmeServer: acmeServer,
	}
}

func (a AcmeServer) Encode() *apiv1.ACMEServer {
	return &apiv1.ACMEServer{
		Name:            a.Name(),
		CreateTime:      timestamppb.New(a.CreatedAt),
		UpdateTime:      timestamppb.New(a.UpdatedAt),
		DisplayName:     a.DisplayName,
		DirectoryUrl:    a.DirectoryUrl,
		IntegrationName: a.IntegrationName,
	}
}

func (a AcmeServer) Name() string {
	return fmt.Sprintf("acmeServers/%s", a.ID)
}

type AcmeServers []AcmeServer

func NewAcmeServers(acmeServers []repository.AcmeServer) AcmeServers {
	as := make(AcmeServers, len(acmeServers))
	for i, acmeServer := range acmeServers {
		as[i] = NewAcmeServer(acmeServer)
	}
	return as
}

func (a AcmeServers) Encode() []*apiv1.ACMEServer {
	acmeServers := make([]*apiv1.ACMEServer, len(a))
	for i, server := range a {
		acmeServers[i] = server.Encode()
	}
	return acmeServers
}
