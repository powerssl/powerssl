package datamapper

import (
	"context"

	"github.com/freerware/work/v4/unit"
	"go.uber.org/zap"

	"powerssl.dev/powerssl/internal/app/apiserver/model"
)

func acmeAccounts(entities []interface{}) model.ACMEAccounts {
	var acmeAccounts []*model.ACMEAccount
	for _, entity := range entities {
		acmeAccounts = append(acmeAccounts, entity.(*model.ACMEAccount))
	}
	return acmeAccounts
}

type acmeAccountDataMapper struct {
	logger *zap.Logger
}

var _ unit.DataMapper = &acmeAccountDataMapper{}

func NewACMEAccountDataMapper(logger *zap.Logger) *acmeAccountDataMapper {
	return &acmeAccountDataMapper{
		logger: logger,
	}
}

func (m acmeAccountDataMapper) Insert(ctx context.Context, mCtx unit.MapperContext, entities ...interface{}) error {
	for _, acmeAccount := range acmeAccounts(entities) {
		m.logger.Sugar().Debug("DEBUG", acmeAccount)
		if _, err := mCtx.Tx.ExecContext(ctx,
			`insert into acme_accounts (id, acme_server_id, display_name, title, description, terms_of_service_agreed, contacts, account_url) values ($1, $2, $3, $4, $5, $6, $7, $8)`,
			acmeAccount.ID, acmeAccount.ACMEServerID, acmeAccount.DisplayName, acmeAccount.Title, acmeAccount.Description, acmeAccount.TermsOfServiceAgreed, acmeAccount.Contacts, acmeAccount.AccountURL); err != nil {
			return err
		}
	}
	return nil
}

func (m acmeAccountDataMapper) Update(ctx context.Context, mCtx unit.MapperContext, entities ...interface{}) error {
	for _, acmeAccount := range acmeAccounts(entities) {
		if _, err := mCtx.Tx.ExecContext(ctx,
			`update acme_accounts set display_name = $1, title = $2, description = $3, terms_of_service_agreed = $4, contacts = $5, account_url = $6, updated_at = now() where id = $4`,
			acmeAccount.DisplayName, acmeAccount.Title, acmeAccount.Description, acmeAccount.TermsOfServiceAgreed, acmeAccount.Contacts, acmeAccount.AccountURL, acmeAccount.ID); err != nil {
			return err
		}
	}
	return nil
}

func (m acmeAccountDataMapper) Delete(ctx context.Context, mCtx unit.MapperContext, entities ...interface{}) error {
	for _, acmeAccount := range acmeAccounts(entities) {
		if _, err := mCtx.Tx.ExecContext(ctx,
			`update acme_accounts set deleted_at = now() where id = $1`,
			acmeAccount.ID); err != nil {
			return err
		}
	}
	return nil
}
