module powerssl.dev/workflow

go 1.15

replace powerssl.dev/common => ../common

replace powerssl.dev/sdk => ../../sdk

require (
	go.temporal.io/sdk v1.4.1 // indirect
	powerssl.dev/sdk v0.0.0-00010101000000-000000000000
)
