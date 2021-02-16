module powerssl.dev/apiserver

go 1.15

replace powerssl.dev/common => ../internal/common

replace powerssl.dev/backend => ../internal/backend

replace powerssl.dev/sdk => ../sdk

replace powerssl.dev/workflow => ../internal/workflow

require (
	github.com/ahmetb/govvv v0.3.0 // indirect
	github.com/freerware/work/v4 v4.0.0-beta.2
	github.com/go-kit/kit v0.10.0
	github.com/go-playground/validator/v10 v10.4.1
	github.com/gogo/protobuf v1.3.2
	github.com/gogo/status v1.1.0
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
	powerssl.dev/backend v0.0.0-00010101000000-000000000000
	powerssl.dev/common v0.0.0-00010101000000-000000000000
	powerssl.dev/sdk v0.0.0-00010101000000-000000000000
	powerssl.dev/workflow v0.0.0-00010101000000-000000000000
)
