package acmeserver

import (
	"context"

	"github.com/freerware/work/v4/unit"
	"go.uber.org/zap"

	"powerssl.dev/powerssl/internal/app/apiserver/model"
)

func acmeServers(entities []interface{}) model.ACMEServers {
	var acmeServers []*model.ACMEServer
	for _, entity := range entities {
		acmeServers = append(acmeServers, entity.(*model.ACMEServer))
	}
	return acmeServers
}

type DataMapper struct {
	logger *zap.Logger
}

var _ unit.DataMapper = &DataMapper{}

func NewDataMapper(logger *zap.Logger) *DataMapper {
	return &DataMapper{
		logger: logger,
	}
}

func (m DataMapper) Insert(ctx context.Context, mCtx unit.MapperContext, entities ...interface{}) error {
	for _, acmeServer := range acmeServers(entities) {
		if _, err := mCtx.Tx.ExecContext(ctx,
			`insert into acme_servers (id, display_name, directory_url, integration_name) values ($1, $2, $3, $4)`,
			acmeServer.ID, acmeServer.DisplayName, acmeServer.DirectoryURL, acmeServer.IntegrationName); err != nil {
			return err
		}
	}
	return nil
}

func (m DataMapper) Update(ctx context.Context, mCtx unit.MapperContext, entities ...interface{}) error {
	for _, acmeServer := range acmeServers(entities) {
		if _, err := mCtx.Tx.ExecContext(ctx,
			`update acme_servers set display_name = $1, directory_url = $2, integration_name = $3, updated_at = now() where id = $4`,
			acmeServer.DisplayName, acmeServer.DirectoryURL, acmeServer.IntegrationName, acmeServer.ID); err != nil {
			return err
		}
	}
	return nil
}

func (m DataMapper) Delete(ctx context.Context, mCtx unit.MapperContext, entities ...interface{}) error {
	for _, acmeServer := range acmeServers(entities) {
		if _, err := mCtx.Tx.ExecContext(ctx,
			`update acme_servers set deleted_at = now() where id = $1`,
			acmeServer.ID); err != nil {
			return err
		}
	}
	return nil
}
