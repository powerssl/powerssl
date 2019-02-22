package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/gogo/status"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc/codes"

	"powerssl.io/internal/pkg/uid"
	"powerssl.io/pkg/apiserver/api"
)

type ACMEServer struct {
	ID        string `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	DisplayName     string
	DirectoryURL    string
	IntegrationName string
}

func (*ACMEServer) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uid.New())
	return nil
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

func FindACMEServerByName(name string, db *gorm.DB) (*ACMEServer, error) {
	s := strings.Split(name, "/")
	if len(s) != 2 {
		return nil, status.Error(codes.InvalidArgument, "malformed name")
	}

	acmeServer := &ACMEServer{}
	if db.Where("id = ?", s[1]).First(&acmeServer).RecordNotFound() {
		return nil, status.Error(codes.NotFound, "not found")
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
