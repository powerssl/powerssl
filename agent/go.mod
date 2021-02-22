module powerssl.dev/agent

go 1.15

replace powerssl.dev/api => ../api

replace powerssl.dev/common => ../internal/common

replace powerssl.dev/sdk => ../sdk

require (
	github.com/ahmetb/govvv v0.3.0
	github.com/go-kit/kit v0.10.0
	github.com/go-playground/validator/v10 v10.4.1
	github.com/opentracing/opentracing-go v1.2.0
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a
	powerssl.dev/common v0.0.0-00010101000000-000000000000
	powerssl.dev/sdk v0.0.0-00010101000000-000000000000
)
