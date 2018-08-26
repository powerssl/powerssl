PROTOC := $(shell which protoc)
JQ := $(shell which jq)

BIN_PATH := $(abspath bin)
PKG_PATH := $(abspath pkg)

export PATH := $(BIN_PATH):$(PATH)

FIND_RELEVANT := find $(PKG_PATH)

GOGO_PROTOBUF_PATH := $(shell go mod download -json github.com/gogo/protobuf | $(JQ) -r '.Dir')
PROTOBUF_PATH := $(GOGO_PROTOBUF_PATH)/protobuf

PROTO_MAPPINGS :=
PROTO_MAPPINGS := $(PROTO_MAPPINGS)Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,

GO_PROTOS := $(sort $(shell $(FIND_RELEVANT) -type f -name '*.proto' -print))
GO_SOURCES := $(GO_PROTOS:%.proto=%.pb.go)

PROTOBUF_TARGETS := bin/.go_protobuf_sources

.DELETE_ON_ERROR:

.ALWAYS_REBUILD:
.PHONY: .ALWAYS_REBUILD

.DEFAULT_GOAL := all
all: build

bin/.go_protobuf_sources: bin/protoc-gen-gogo
	$(FIND_RELEVANT) -type f -name '*.pb.go' -exec rm {} +
	set -e; for dir in $(sort $(dir $(GO_PROTOS))); do \
		$(PROTOC) -I$(PKG_PATH):$(GOGO_PROTOBUF_PATH):$(PROTOBUF_PATH) --gogo_out=$(PROTO_MAPPINGS),plugins=grpc:$(GOPATH)/src $$dir/*.proto; \
	done
	gofmt -s -w $(GO_SOURCES)
	touch $@

bin/protoc-gen-gogo:
	go build -o bin/protoc-gen-gogo $$(go mod download -json github.com/gogo/protobuf | $(JQ) -r '.Dir')/protoc-gen-gogo

bin/powerssl-apiserver: .ALWAYS_REBUILD
	go build -o bin/powerssl-apiserver ./cmd/powerssl-apiserver

bin/powerctl: .ALWAYS_REBUILD
	go build -o bin/powerctl ./cmd/powerctl

.PHONY: build
build: bin/powerssl-apiserver bin/powerctl

.PHONY: fmt
fmt:
	go fmt $$(go list ./...)
	clang-format -i --style=google $(GO_PROTOS)

.PHONY: vet
vet:
	go vet $$(go list ./...)

.PHONY: clean
clean:
	go clean ./cmd/powerssl-api ./cmd/powerctl
	rm -f bin/.go_protobuf_sources
	rm -f bin/powerssl-apiserver bin/powerctl

.PHONY: protobuf
protobuf: $(PROTOBUF_TARGETS)

# build-linux:
# 	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/powerctl -v ./cmd/powerctl
# docker-build:
# 	docker run --rm -it -e GO111MODULE=on -v $$(pwd):/go/src -v $$(pwd)/bin:/go/bin -w /go/src golang:1.11rc1 go build -o /go/bin/powerctl -v ./cmd/powerctl
