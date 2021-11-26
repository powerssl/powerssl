module powerssl.dev/common/telemetry

go 1.17

replace powerssl.dev/common => ./..

require (
	github.com/google/wire v0.5.0
	github.com/prometheus/client_golang v1.11.0
	go.opentelemetry.io/otel/exporters/prometheus v0.25.0
	go.opentelemetry.io/otel/exporters/stdout/stdoutmetric v0.25.0
	go.opentelemetry.io/otel/metric v0.25.0
	go.opentelemetry.io/otel/sdk v1.2.0
	go.opentelemetry.io/otel/sdk/export/metric v0.25.0
	go.opentelemetry.io/otel/sdk/metric v0.25.0
	go.opentelemetry.io/otel/trace v1.2.0
	powerssl.dev/common v0.0.0-00010101000000-000000000000
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.26.0 // indirect
	github.com/prometheus/procfs v0.6.0 // indirect
	go.opentelemetry.io/otel v1.2.0 // indirect
	go.opentelemetry.io/otel/internal/metric v0.25.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.19.1 // indirect
	golang.org/x/sys v0.0.0-20210823070655-63515b42dcdf // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	logur.dev/adapter/zap v0.5.0 // indirect
	logur.dev/logur v0.16.1 // indirect
)
