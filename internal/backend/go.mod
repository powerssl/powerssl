module powerssl.dev/backend

go 1.16

replace powerssl.dev/api => ../../api

replace powerssl.dev/common => ../common

replace powerssl.dev/backend => ../backend

replace powerssl.dev/sdk => ../../sdk

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-kit/kit v0.10.0
	github.com/go-playground/validator/v10 v10.4.1
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/hashicorp/vault/api v1.0.5-0.20210316211753-9e6de0762492
	github.com/johanbrandhorst/certify v1.8.1
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pborman/uuid v1.2.1
	github.com/uber-go/tally v3.3.17+incompatible
	go.temporal.io/sdk v1.4.1
	go.uber.org/zap v1.16.0
	google.golang.org/grpc v1.35.0
	gopkg.in/square/go-jose.v2 v2.5.1
	powerssl.dev/common v0.0.0-00010101000000-000000000000
	powerssl.dev/sdk v0.0.0-00010101000000-000000000000
)
