// +build tools

package tools

import (
	_ "github.com/ahmetb/govvv"
	_ "github.com/gogo/protobuf/protoc-gen-gogo"
	_ "github.com/golang/protobuf/protoc-gen-go"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger"
	_ "github.com/jteeuwen/go-bindata"
	_ "golang.org/x/tools/cmd/stringer"
)
