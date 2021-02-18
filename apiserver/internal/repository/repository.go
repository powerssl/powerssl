package repository

import (
	"github.com/freerware/work/v4/unit"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"powerssl.dev/apiserver/internal/model"
	"powerssl.dev/apiserver/internal/repository/acmeaccount"
	"powerssl.dev/apiserver/internal/repository/acmeserver"
	"powerssl.dev/apiserver/internal/repository/interface"
)

type Repositories struct {
	ACMEAccounts _interface.ACMEAccountRepository
	ACMEServers  _interface.ACMEServerRepository
	uniter       unit.Uniter
}

func NewRepositories(db *sqlx.DB, logger *zap.Logger) *Repositories {
	dataMappers := map[unit.TypeName]unit.DataMapper{
		unit.TypeNameOf(&model.ACMEAccount{}): acmeaccount.NewDataMapper(logger),
		unit.TypeNameOf(&model.ACMEServer{}):  acmeserver.NewDataMapper(logger),
	}
	opts := []unit.Option{
		unit.DB(db.DB),
		unit.DataMappers(dataMappers),
		unit.Logger(logger),
		//unit.Scope(),
	}
	return &Repositories{
		ACMEAccounts: acmeaccount.NewRepository(db, logger),
		ACMEServers:  acmeserver.NewRepository(db, logger),
		uniter:       unit.NewUniter(opts...),
	}
}

func (r *Repositories) UnitOfWork() (unit.Unit, error) {
	return r.uniter.Unit()
}
