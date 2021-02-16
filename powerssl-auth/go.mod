module powerssl.dev/auth

go 1.15

replace powerssl.dev/common => ../internal/common

require (
	github.com/ahmetb/govvv v0.3.0 // indirect
	github.com/arschles/assert v2.0.0+incompatible // indirect
	github.com/arschles/go-bindata-html-template v0.0.0-20170123182818-839a6918b9ff
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-kit/kit v0.10.0
	github.com/go-playground/validator/v10 v10.4.1
	github.com/pborman/uuid v1.2.1
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a
	gopkg.in/square/go-jose.v2 v2.5.1
	powerssl.dev/common v0.0.0-00010101000000-000000000000
)
