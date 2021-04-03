module powerssl.dev/apiserver

go 1.16

replace powerssl.dev/api => ../api

replace powerssl.dev/backend => ../backend

replace powerssl.dev/common => ../common

replace powerssl.dev/sdk => ../sdk

replace powerssl.dev/workflow => ../workflow

require (
	github.com/Masterminds/squirrel v1.5.0
	github.com/ahmetb/govvv v0.3.0
	github.com/go-bindata/go-bindata v3.1.2+incompatible
	github.com/go-kit/kit v0.10.0
	github.com/go-playground/validator/v10 v10.4.1
	github.com/gogo/status v1.1.0
	github.com/golang-migrate/migrate v3.5.4+incompatible
	github.com/golang-migrate/migrate/v4 v4.14.1
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.1.4
	github.com/jmoiron/sqlx v1.3.1
	github.com/lib/pq v1.9.0
	github.com/mennanov/fieldmask-utils v0.3.3
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.9.0
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1
	go.temporal.io/sdk v1.4.1
	go.uber.org/zap v1.16.0
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a
	google.golang.org/grpc v1.35.0
	powerssl.dev/api v0.0.0-00010101000000-000000000000
	powerssl.dev/backend v0.0.0-00010101000000-000000000000
	powerssl.dev/common v0.0.0-00010101000000-000000000000
	powerssl.dev/sdk v0.0.0-00010101000000-000000000000
	powerssl.dev/workflow v0.0.0-00010101000000-000000000000
)
