PROTOC := $(shell which protoc)

BIN_PATH := $(abspath bin)
PKG_PATH := $(abspath pkg)
PROTO_PATH := $(abspath api/proto)

export PATH := $(BIN_PATH):$(PATH)

FIND_RELEVANT := find $(PKG_PATH)
FIND_PROTO := find $(PROTO_PATH)

GOGO_GOOGLEAPIS_PATH := $(shell go mod download -json github.com/gogo/googleapis | grep '"Dir"' | cut -d '"' -f 4)
GOGO_PROTOBUF_PATH := $(shell go mod download -json github.com/gogo/protobuf | grep '"Dir"' | cut -d '"' -f 4)
PROTOBUF_PATH := $(GOGO_PROTOBUF_PATH)/protobuf

PROTO_MAPPINGS :=
PROTO_MAPPINGS := $(PROTO_MAPPINGS)Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,
PROTO_MAPPINGS := $(PROTO_MAPPINGS)Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,
PROTO_MAPPINGS := $(PROTO_MAPPINGS)Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types,
PROTO_MAPPINGS := $(PROTO_MAPPINGS)Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,

PROTOS := $(sort $(shell $(FIND_PROTO) -type f -name '*.proto' -print))

.DELETE_ON_ERROR:

.ALWAYS_REBUILD:
.PHONY: .ALWAYS_REBUILD

.DEFAULT_GOAL := all
all: build

bin/protoc-gen-gogo:
	go build -o bin/protoc-gen-gogo $$(go mod download -json github.com/gogo/protobuf | grep '"Dir"' | cut -d '"' -f 4)/protoc-gen-gogo

bin/protoc-gen-grpc-web:
	rm -rf /tmp/grpc-web
	git clone --branch 0.4.0 https://github.com/grpc/grpc-web.git /tmp/grpc-web
	cd /tmp/grpc-web/javascript/net/grpc/web && \
		make protoc-gen-grpc-web && \
		install protoc-gen-grpc-web $(BIN_PATH)/protoc-gen-grpc-web

bin/powerssl-agent: .ALWAYS_REBUILD
	go build -o bin/powerssl-agent powerssl.io/cmd/powerssl-agent

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
build: bin/powerssl-agent bin/powerssl-apiserver bin/powerssl-auth bin/powerssl-controller bin/powerssl-integration-acme bin/powerssl-integration-cloudflare bin/powerssl-signer bin/powerctl

.PHONY: install-agent
install-agent:
	go install powerssl.io/cmd/powerssl-agent

.PHONY: install-powerctl
install-powerctl:
	go install powerssl.io/cmd/powerctl

.PHONY: fmt
fmt:
	go fmt $$(go list ./...)
	clang-format -i --style=google $(PROTOS)

.PHONY: vet
vet:
	go vet $$(go list ./...)

.PHONY: clean
clean:
	go clean powerssl.io/cmd/powerctl powerssl.io/cmd/powerssl-agent powerssl.io/cmd/powerssl-apiserver powerssl.io/cmd/powerssl-auth powerssl.io/cmd/powerssl-controller powerssl.io/cmd/powerssl-integration-acme powerssl.io/cmd/powerssl-integration-cloudflare powerssl.io/cmd/powerssl-signer
	rm -f bin/.go_protobuf_sources
	rm -f bin/powerctl bin/powerssl-agent bin/powerssl-apiserver bin/powerssl-auth bin/powerssl-controller bin/powerssl-integration-acme bin/powerssl-integration-cloudflare bin/powerssl-signer

.PHONY: protobuf
protobuf:
	$(FIND_RELEVANT) -type f -name '*.pb.go' -exec rm {} +
	set -e; for dir in $(sort $(dir $(PROTOS))); do \
		$(PROTOC) \
			-I$(PROTO_PATH):$(GOGO_GOOGLEAPIS_PATH):$(GOGO_PROTOBUF_PATH):$(PROTOBUF_PATH) \
			--gogo_out=$(PROTO_MAPPINGS),plugins=grpc:$(GOPATH)/src \
			$$dir/*.proto; \
	done

.PHONY: javascript-sdk
javascript-sdk: bin/protoc-gen-grpc-web
	set -e; for dir in $(sort $(dir $(PROTOS))); do \
		$(PROTOC) \
			-I$(PROTO_PATH):$(GOGO_GOOGLEAPIS_PATH):$(GOGO_PROTOBUF_PATH):$(PROTOBUF_PATH) \
			--js_out=import_style=commonjs:vendor/javascript-sdk \
			--grpc-web_out=import_style=commonjs,mode=grpcwebtext:vendor/javascript-sdk \
			$$dir/*.proto; \
	done

.PHONY: tools
tools:
	GO111MODULE=off go get golang.org/x/tools/cmd/stringer

.PHONY: generate
generate:
	go generate $$(go list ./...)

.PHONY: images
images: agent-image apiserver-image auth-image controller-image envoy-image integration-acme-image integration-cloudflare-image powerctl-image signer-image

.PHONY: agent-image
agent-image:
	docker build -f build/docker/agent/Dockerfile -t powerssl/agent .

.PHONY: apiserver-image
apiserver-image:
	docker build -f build/docker/apiserver/Dockerfile -t powerssl/apiserver .

.PHONY: auth-image
auth-image:
	docker build -f build/docker/auth/Dockerfile -t powerssl/auth .

.PHONY: controller-image
controller-image:
	docker build -f build/docker/controller/Dockerfile -t powerssl/controller .

.PHONY: envoy-image
envoy-image:
	docker build -f build/docker/envoy/Dockerfile -t powerssl/evnoy .

.PHONY: integration-acme-image
integration-acme-image:
	docker build -f build/docker/integration-acme/Dockerfile -t powerssl/integration-acme .

.PHONY: integration-cloudflare-image
integration-cloudflare-image:
	docker build -f build/docker/integration-cloudflare/Dockerfile -t powerssl/integration-cloudflare .

.PHONY: powerctl-image
powerctl-image:
	docker build -f build/docker/powerctl/Dockerfile -t powerssl/powerctl .

.PHONY: signer-image
signer-image:
	docker build -f build/docker/signer/Dockerfile -t powerssl/signer .
