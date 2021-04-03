module powerssl.dev/integration/cloudflare

go 1.16

replace powerssl.dev/api => ../../api

replace powerssl.dev/common => ../../common

replace powerssl.dev/sdk => ../../sdk

require (
	github.com/ahmetb/govvv v0.3.0
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1
	powerssl.dev/common v0.0.0-00010101000000-000000000000
	powerssl.dev/sdk v0.0.0-00010101000000-000000000000
)
