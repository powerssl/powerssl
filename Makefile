PROTOC := $(shell which protoc)

BIN_PATH := $(abspath bin)
PKG_PATH := $(abspath pkg)

export PATH := $(BIN_PATH):$(PATH)

FIND_RELEVANT := find $(PKG_PATH)

GOGO_GOOGLEAPIS_PATH := $(shell go mod download -json github.com/gogo/googleapis | grep '"Dir"' | cut -d '"' -f 4)
GOGO_PROTOBUF_PATH := $(shell go mod download -json github.com/gogo/protobuf | grep '"Dir"' | cut -d '"' -f 4)
PROTOBUF_PATH := $(GOGO_PROTOBUF_PATH)/protobuf

PROTO_MAPPINGS :=
PROTO_MAPPINGS := $(PROTO_MAPPINGS)Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,
PROTO_MAPPINGS := $(PROTO_MAPPINGS)Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,
PROTO_MAPPINGS := $(PROTO_MAPPINGS)Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,

GO_PROTOS := $(sort $(shell $(FIND_RELEVANT) -type f -name '*.proto' -print))
GO_SOURCES := $(GO_PROTOS:%.proto=%.pb.go)

PROTOBUF_TARGETS := bin/.go_protobuf_sources

.DELETE_ON_ERROR:

.ALWAYS_REBUILD:
.PHONY: .ALWAYS_REBUILD

.DEFAULT_GOAL := all
all: build

bin/.go_protobuf_sources: bin/protoc-gen-gogo bin/protoc-gen-gotemplate bin/protoc-gen-grpc-web
	$(FIND_RELEVANT) -type f -name '*.pb.go' -exec rm {} +
	set -e; for dir in $(sort $(dir $(GO_PROTOS))); do \
		$(PROTOC) \
			-I$(PKG_PATH):$(GOGO_GOOGLEAPIS_PATH):$(GOGO_PROTOBUF_PATH):$(PROTOBUF_PATH) \
			--gogo_out=$(PROTO_MAPPINGS),plugins=grpc:$(GOPATH)/src \
			--gotemplate_out=$(PKG_PATH)/resource/generated \
			--js_out=import_style=commonjs:vendor/javascript \
			--grpc-web_out=import_style=commonjs,mode=grpcwebtext:vendor/javascript \
			$$dir/*.proto; \
	done
	gofmt -s -w $(PKG_PATH)/resource/generated
	find pkg/resource/generated -type d -depth 1 | cut -d '/' -f 4 | xargs -I '{}' sh -c "eval $$(echo mv -n $(PKG_PATH)/resource/generated/'{}'/service/\* $(PKG_PATH)/resource/'{}'/)"
	rm -rf pkg/resource/generated/*/service
	gofmt -s -w $(GO_SOURCES)
	touch $@

.PHONY: controllerproto
controllerproto: bin/protoc-gen-gogo
	# $(FIND_RELEVANT) -type f -name '*.pb.go' -exec rm {} +
	set -e; for dir in $(sort $(dir $(GO_PROTOS))); do \
		echo $(PROTOC) \
			-I$(PKG_PATH):$(GOGO_GOOGLEAPIS_PATH):$(GOGO_PROTOBUF_PATH):$(PROTOBUF_PATH) \
			--gogo_out=$(PROTO_MAPPINGS),plugins=grpc:$(GOPATH)/src \
			$$dir/*.proto; \
	done
	# gofmt -s -w $(GO_SOURCES)

bin/protoc-gen-gotemplate:
	go build -o bin/protoc-gen-gotemplate $$(go mod download -json moul.io/protoc-gen-gotemplate | grep '"Dir"' | cut -d '"' -f 4)

bin/protoc-gen-gogo:
	go build -o bin/protoc-gen-gogo $$(go mod download -json github.com/gogo/protobuf | grep '"Dir"' | cut -d '"' -f 4)/protoc-gen-gogo

# Not used just as a reference
bin/protoc-gen-grpc-web:
	rm -rf /tmp/grpc-web
	git clone --branch 0.4.0 https://github.com/grpc/grpc-web.git /tmp/grpc-web
	cd /tmp/grpc-web/javascript/net/grpc/web && \
		make protoc-gen-grpc-web && \
		install protoc-gen-grpc-web $(BIN_PATH)/protoc-gen-grpc-web

bin/powerssl-apiserver: .ALWAYS_REBUILD
	go build -o bin/powerssl-apiserver powerssl.io/cmd/powerssl-apiserver

bin/powerssl-auth: .ALWAYS_REBUILD
	go build -o bin/powerssl-auth powerssl.io/cmd/powerssl-auth

bin/powerssl-controller: .ALWAYS_REBUILD
	go build -o bin/powerssl-controller powerssl.io/cmd/powerssl-controller

bin/powerssl-integration-acme: .ALWAYS_REBUILD
	go build -o bin/powerssl-integration-acme powerssl.io/cmd/powerssl-integration-acme

bin/powerssl-integration-cloudflare: .ALWAYS_REBUILD
	go build -o bin/powerssl-integration-cloudflare powerssl.io/cmd/powerssl-integration-cloudflare

bin/powerssl-signer: .ALWAYS_REBUILD
	go build -o bin/powerssl-signer powerssl.io/cmd/powerssl-signer

bin/powerctl: .ALWAYS_REBUILD
	go build -o bin/powerctl powerssl.io/cmd/powerctl

.PHONY: build
build: bin/powerssl-apiserver bin/powerssl-auth bin/powerssl-controller bin/powerssl-integration-acme bin/powerssl-integration-cloudflare bin/powerssl-signer bin/powerctl

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
#

.PHONY: tools
tools:
	GO111MODULE=off go get golang.org/x/tools/cmd/stringer

.PHONY: generate
generate:
	go generate $$(go list ./...)

.PHONY: images
images: apiserver-image auth-image controller-image envoy-image integration-acme-image integration-cloudflare-image powerctl-image signer-image

.PHONY: apiserver-image
apiserver-image:
	docker build -f dockerfiles/apiserver/Dockerfile -t powerssl/apiserver .

.PHONY: auth-image
auth-image:
	docker build -f dockerfiles/auth/Dockerfile -t powerssl/auth .

.PHONY: controller-image
controller-image:
	docker build -f dockerfiles/controller/Dockerfile -t powerssl/controller .

.PHONY: envoy-image
envoy-image:
	docker build -f dockerfiles/envoy/Dockerfile -t powerssl/evnoy .

.PHONY: integration-acme-image
integration-acme-image:
	docker build -f dockerfiles/integration-acme/Dockerfile -t powerssl/integration-acme .

.PHONY: integration-cloudflare-image
integration-cloudflare-image:
	docker build -f dockerfiles/integration-cloudflare/Dockerfile -t powerssl/integration-cloudflare .

.PHONY: powerctl-image
powerctl-image:
	docker build -f dockerfiles/powerctl/Dockerfile -t powerssl/powerctl .

.PHONY: signer-image
signer-image:
	docker build -f dockerfiles/signer/Dockerfile -t powerssl/signer .
