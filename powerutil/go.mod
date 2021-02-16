module powerssl.dev/powerutil

go 1.15

replace powerssl.dev/common => ../internal/common

replace powerssl.dev/backend => ../internal/backend

replace powerssl.dev/sdk => ../sdk

require (
	github.com/ahmetb/govvv v0.3.0 // indirect
	github.com/cloudflare/cfssl v1.5.0
	github.com/ghodss/yaml v1.0.0
	github.com/golang-migrate/migrate/v4 v4.14.1
	github.com/hashicorp/vault/api v1.0.5-0.20200117231345-460d63e36490
	github.com/spf13/cobra v1.1.3
	powerssl.dev/backend v0.0.0-00010101000000-000000000000
	powerssl.dev/common v0.0.0-00010101000000-000000000000
)
