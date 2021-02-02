package apiserver

import (
	"go.uber.org/zap"
	"powerssl.dev/powerssl/internal/app/apiserver/datamapper"
	"powerssl.dev/powerssl/internal/app/apiserver/model"

	"github.com/freerware/work/v4/unit"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/jmoiron/sqlx"
	stdopentracing "github.com/opentracing/opentracing-go"
	temporalclient "go.temporal.io/sdk/client"

	"powerssl.dev/powerssl/internal/app/apiserver/acmeaccount"
	"powerssl.dev/powerssl/internal/app/apiserver/acmeserver"
	"powerssl.dev/powerssl/internal/app/apiserver/certificate"
	"powerssl.dev/powerssl/internal/app/apiserver/repository"
	acmeaccountrepository "powerssl.dev/powerssl/internal/app/apiserver/repository/acmeaccount"
	acmeserverrepository "powerssl.dev/powerssl/internal/app/apiserver/repository/acmeserver"
	"powerssl.dev/powerssl/internal/app/apiserver/user"
	"powerssl.dev/powerssl/internal/pkg/auth"
	"powerssl.dev/powerssl/internal/pkg/transport"
	"powerssl.dev/powerssl/internal/pkg/vault"
)

func makeServices(db *sqlx.DB, logger log.Logger, tracer stdopentracing.Tracer, duration metrics.Histogram, temporalClient temporalclient.Client, vaultClient *vault.Client, jwtPublicKeyFile string) ([]transport.Service, error) {
	zapLogger, _ := zap.NewDevelopment()
	dataMapper := map[unit.TypeName]unit.DataMapper{
		unit.TypeNameOf(&model.ACMEAccount{}): datamapper.NewACMEAccountDataMapper(zapLogger),
		unit.TypeNameOf(&model.ACMEServer{}): datamapper.NewACMEServerDataMapper(zapLogger),
	}
	opts := []unit.Option{
		unit.DB(db.DB),
		unit.DataMappers(dataMapper),
		unit.Logger(zapLogger),
		//unit.Scope(),
	}
	uniter := unit.NewUniter(opts...)

	repositories := repository.NewRepositories()
	repositories.ACMEAccounts = acmeaccountrepository.NewRepository(db)
	repositories.ACMEServers = acmeserverrepository.NewRepository(db)
	auth, err := auth.NewParser(jwtPublicKeyFile)
	if err != nil {
		return nil, err
	}
	return []transport.Service{
		acmeaccount.New(uniter, repositories, logger, tracer, duration, temporalClient, vaultClient, auth),
		acmeserver.New(uniter, repositories, logger, tracer, duration, auth),
		certificate.New(db, logger, tracer, duration, auth),
		user.New(db, logger, tracer, duration, auth),
	}, nil
}
