package apiserver

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"

	acmeaccountmodel "powerssl.dev/powerssl/internal/app/apiserver/acmeaccount/model"
	acmeservermodel "powerssl.dev/powerssl/internal/app/apiserver/acmeserver/model"
	certificatemodel "powerssl.dev/powerssl/internal/app/apiserver/certificate/model"
	usermodel "powerssl.dev/powerssl/internal/app/apiserver/user/model"
)

func createTables(db *pg.DB) error {
	models := []interface{}{
		(*acmeaccountmodel.ACMEAccount)(nil),
		(*acmeservermodel.ACMEServer)(nil),
		(*certificatemodel.Certificate)(nil),
		(*usermodel.User)(nil),
	}
	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
