module powerssl.dev/temporal

go 1.16

replace powerssl.dev/common => ../internal/common

replace github.com/apache/thrift => github.com/apache/thrift v0.0.0-20161221203622-b2a4d4ae21c7

require (
	github.com/ahmetb/govvv v0.3.0
	github.com/go-playground/validator/v10 v10.4.1
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1
	go.temporal.io/api v1.4.0
	go.temporal.io/sdk v1.3.0
	go.temporal.io/server v1.6.4
	powerssl.dev/common v0.0.0-00010101000000-000000000000
)
