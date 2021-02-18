module powerssl.dev/backend

go 1.15

replace powerssl.dev/common => ../common

replace powerssl.dev/backend => ../backend

replace powerssl.dev/sdk => ../../sdk

require (
	github.com/go-kit/kit v0.10.0
	github.com/go-playground/validator/v10 v10.4.1
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/hashicorp/vault/api v1.0.5-0.20200117231345-460d63e36490
	github.com/johanbrandhorst/certify v1.8.1
	github.com/lib/pq v1.3.0 // indirect
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pborman/uuid v1.2.1
	github.com/uber-go/tally v3.3.17+incompatible
	go.temporal.io/sdk v1.4.1
	go.uber.org/zap v1.16.0
	google.golang.org/grpc v1.35.0
	powerssl.dev/sdk v0.0.0-00010101000000-000000000000
)
