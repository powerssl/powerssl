module powerssl.dev/sdk

go 1.16

replace powerssl.dev/api => ../api

replace powerssl.dev/common => ../internal/common

require (
	github.com/go-kit/kit v0.10.0
	github.com/go-playground/validator/v10 v10.4.1
	github.com/gogo/protobuf v1.3.2
	github.com/hashicorp/vault/api v1.0.4
	github.com/kenshaw/pemutil v0.1.0
	github.com/opentracing/opentracing-go v1.2.0
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a
	google.golang.org/grpc v1.35.0
	google.golang.org/protobuf v1.25.0
	powerssl.dev/api v0.0.0-00010101000000-000000000000
	powerssl.dev/common v0.0.0-00010101000000-000000000000
)
