package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	"powerssl.dev/backend/ctxkey"
	"powerssl.dev/common/log"

	"powerssl.dev/apiserver/internal/repository/acmeaccount"
	"powerssl.dev/apiserver/internal/repository/acmeserver"
	"powerssl.dev/apiserver/internal/repository/interface"
)

var transactionValue = ctxkey.New("dev.powerssl.apiserver.internal.repository")

type Repositories struct {
	db           *sqlx.DB
	ACMEAccounts _interface.ACMEAccountRepository
	ACMEServers  _interface.ACMEServerRepository
}

func NewRepositories(db *sqlx.DB, logger log.Logger) *Repositories {
	bar := &sqlxInterface{db: db}
	return &Repositories{
		db:           db,
		ACMEAccounts: acmeaccount.NewRepository(bar, logger),
		ACMEServers:  acmeserver.NewRepository(bar, logger),
	}
}

func (r *Repositories) Transaction(ctx context.Context, fn func(context.Context) error) (err error) {
	var tx *sqlx.Tx
	if tx, err = r.db.BeginTxx(ctx, nil); err != nil {
		return fmt.Errorf("failed to start transaction %w", err)
	}
	defer func() {
		if err == nil {
			return
		}

		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			err = fmt.Errorf("failed to rollback (%s) %w", rollbackErr.Error(), err)
		}
	}()

	ctx = context.WithValue(ctx, transactionValue, tx)
	if err = fn(ctx); err != nil {
		err = fmt.Errorf("failed to execute transaction %w", err)
		return err
	}

	if err = tx.Commit(); err != nil {
		err = fmt.Errorf("failed to commit transaction %w", err)
	}

	return err
}
