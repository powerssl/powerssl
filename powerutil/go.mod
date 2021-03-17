module powerssl.dev/powerutil

go 1.16

replace powerssl.dev/api => ../api

replace powerssl.dev/common => ../internal/common

replace powerssl.dev/backend => ../internal/backend

replace powerssl.dev/sdk => ../sdk

require (
	github.com/ahmetb/govvv v0.3.0
	github.com/cloudflare/cfssl v1.5.0
	github.com/ghodss/yaml v1.0.0
	github.com/go-bindata/go-bindata v3.1.2+incompatible
	github.com/hashicorp/vault/api v1.0.5-0.20210316211753-9e6de0762492
	github.com/spf13/cobra v1.1.3
	powerssl.dev/backend v0.0.0-00010101000000-000000000000
	powerssl.dev/common v0.0.0-00010101000000-000000000000
)
