module powerssl.dev/integration/acme

go 1.15

replace powerssl.dev/powerssl => ../..

require (
	github.com/eggsampler/acme/v2 v2.0.1
	github.com/opentracing/opentracing-go v1.2.0
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1
	powerssl.dev/powerssl v0.0.0-00010101000000-000000000000
)
