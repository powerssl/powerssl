module powerssl.dev/grpcgateway

go 1.16

replace powerssl.dev/api => ../api

replace powerssl.dev/backend => ../backend

replace powerssl.dev/common => ../common

replace powerssl.dev/sdk => ../sdk

require (
	github.com/ahmetb/govvv v0.3.0
	github.com/go-bindata/go-bindata v3.1.2+incompatible
	github.com/go-kit/kit v0.10.0
	github.com/go-playground/validator/v10 v10.4.1
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.2.0
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1
	golang.org/x/net v0.0.0-20210119194325-5f4716e94777
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a
	google.golang.org/grpc v1.35.0
	powerssl.dev/api v0.0.0-00010101000000-000000000000
	powerssl.dev/backend v0.0.0-00010101000000-000000000000
	powerssl.dev/common v0.0.0-00010101000000-000000000000
)
