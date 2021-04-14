module powerssl.dev/auth

go 1.16

replace powerssl.dev/api => ../api

replace powerssl.dev/backend => ../backend

replace powerssl.dev/common => ../common

replace powerssl.dev/sdk => ../sdk

require (
	github.com/ahmetb/govvv v0.3.0
	github.com/arschles/assert v2.0.0+incompatible // indirect
	github.com/arschles/go-bindata-html-template v0.0.0-20170123182818-839a6918b9ff
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-bindata/go-bindata v3.1.2+incompatible
	github.com/go-playground/validator/v10 v10.4.1
	github.com/google/go-github v17.0.0+incompatible
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/pborman/uuid v1.2.1
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1
	golang.org/x/oauth2 v0.0.0-20210201163806-010130855d6c
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a
	gopkg.in/square/go-jose.v2 v2.5.1
	powerssl.dev/backend v0.0.0-00010101000000-000000000000
	powerssl.dev/common v0.0.0-00010101000000-000000000000
)
