PROTOC := $(shell which protoc)

BIN_PATH := $(abspath bin)
PKG_PATH := $(abspath pkg)
PROTO_PATH := $(abspath api/protobuf-spec)

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

EXTERNAL_TOOLS=\
	github.com/ddollar/forego \
	github.com/jteeuwen/go-bindata/... \
	golang.org/x/tools/cmd/stringer


.DELETE_ON_ERROR:

.ALWAYS_REBUILD:
.PHONY: .ALWAYS_REBUILD

.DEFAULT_GOAL := all
all: build

# bootstrap the build by downloading additional tools
.PHONY: bootstrap
bootstrap:
	@for tool in  $(EXTERNAL_TOOLS) ; do \
		echo "Installing/Updating $$tool" ; \
		GO111MODULE=off go get -u $$tool; \
	done
	@echo
	@echo "Make sure you have installed Protocol Buffers - Protocol Compiler and Protobuf Go Runtime"
	@echo
	@echo "On MacOS this can be achieved this way:"
	@echo "$ brew install protobuf protoc-gen-go clang-format"

bin/protoc-gen-gogo:
	go build -o bin/protoc-gen-gogo $$(go mod download -json github.com/gogo/protobuf | grep '"Dir"' | cut -d '"' -f 4)/protoc-gen-gogo

bin/powerssl-agent: .ALWAYS_REBUILD
	COMPONENT=powerssl-agent scripts/build-go.sh

bin/powerssl-apiserver: .ALWAYS_REBUILD
	COMPONENT=powerssl-apiserver scripts/build-go.sh

bin/powerssl-auth: .ALWAYS_REBUILD
	COMPONENT=powerssl-auth scripts/build-go.sh

bin/powerssl-controller: .ALWAYS_REBUILD
	COMPONENT=powerssl-controller scripts/build-go.sh

bin/powerssl-integration-acme: .ALWAYS_REBUILD
	COMPONENT=powerssl-integration-acme scripts/build-go.sh

bin/powerssl-integration-cloudflare: .ALWAYS_REBUILD
	COMPONENT=powerssl-integration-cloudflare scripts/build-go.sh

bin/powerssl-signer: .ALWAYS_REBUILD
	COMPONENT=powerssl-signer scripts/build-go.sh

bin/powerssl-webapp: .ALWAYS_REBUILD
	COMPONENT=powerssl-webapp scripts/build-go.sh

bin/powerctl: .ALWAYS_REBUILD
	COMPONENT=powerctl scripts/build-go.sh

bin/powerutil: .ALWAYS_REBUILD
	COMPONENT=powerutil scripts/build-go.sh

.PHONY: build
build: bin/powerssl-agent bin/powerssl-apiserver bin/powerssl-auth bin/powerssl-controller bin/powerssl-integration-acme bin/powerssl-integration-cloudflare bin/powerssl-signer bin/powerssl-webapp bin/powerctl bin/powerutil

.PHONY: docs
docs:
	go run powerssl.io/tools/gendocs

.PHONY: install-agent
install-agent:
	COMPONENT=powerssl-agent scripts/install.sh

.PHONY: install-powerctl
install-powerctl:
	COMPONENT=powerctl scripts/install.sh

.PHONY: install-powerutil
install-powerutil:
	COMPONENT=powerutil scripts/install.sh

.PHONY: fmt
fmt:
	go fmt $$(go list ./...)
	clang-format -i --style=Google $(PROTOS)

.PHONY: vet
vet:
	go vet $$(go list ./...)

.PHONY: clean
clean:
	go clean powerssl.io/cmd/powerctl powerssl.io/cmd/powerssl-agent powerssl.io/cmd/powerssl-apiserver powerssl.io/cmd/powerssl-auth powerssl.io/cmd/powerssl-controller powerssl.io/cmd/powerssl-integration-acme powerssl.io/cmd/powerssl-integration-cloudflare powerssl.io/cmd/powerssl-signer powerssl.io/cmd/powerssl-webapp powerssl.io/cmd/powerutil
	rm -f bin/.go_protobuf_sources
	rm -f bin/powerctl bin/powerssl-agent bin/powerssl-apiserver bin/powerssl-auth bin/powerssl-controller bin/powerssl-integration-acme bin/powerssl-integration-cloudflare bin/powerssl-signer bin/powerssl-webapp bin/powerutil

.PHONY: protobuf
protobuf: bin/protoc-gen-gogo
	$(FIND_RELEVANT) -type f -name '*.pb.go' -exec rm {} +
	@rm -f powerssl.io && ln -s . powerssl.io
	set -e; for dir in $(sort $(dir $(PROTOS))); do \
		$(PROTOC) \
			-I$(PROTO_PATH):$(GOGO_GOOGLEAPIS_PATH):$(GOGO_PROTOBUF_PATH):$(PROTOBUF_PATH) \
			--gogo_out=$(PROTO_MAPPINGS),plugins=grpc:. \
			$$dir/*.proto; \
	done
	@rm -f powerssl.io

.PHONY: generate
generate:
	go generate $$(go list ./...)

.PHONY: images
images: agent-image apiserver-image auth-image builder-image controller-image envoy-image integration-acme-image integration-cloudflare-image powerctl-image signer-image webapp-image powerutil-image

.PHONY: agent-image
agent-image:
	COMPONENT=agent scripts/build-image.sh

.PHONY: apiserver-image
apiserver-image:
	COMPONENT=apiserver scripts/build-image.sh

.PHONY: auth-image
auth-image:
	COMPONENT=auth scripts/build-image.sh

.PHONY: builder-image
builder-image:
	COMPONENT=builder scripts/build-image.sh

.PHONY: controller-image
controller-image:
	COMPONENT=controller scripts/build-image.sh

.PHONY: envoy-image
envoy-image:
	COMPONENT=envoy scripts/build-image.sh

.PHONY: integration-acme-image
integration-acme-image:
	COMPONENT=integration-acme scripts/build-image.sh

.PHONY: integration-cloudflare-image
integration-cloudflare-image:
	COMPONENT=integration-cloudflare scripts/build-image.sh

.PHONY: powerctl-image
powerctl-image:
	COMPONENT=powerctl scripts/build-image.sh

.PHONY: powerutil-image
powerutil-image:
	COMPONENT=powerutil scripts/build-image.sh

.PHONY: signer-image
signer-image:
	COMPONENT=signer scripts/build-image.sh

.PHONY: webapp-image
webapp-image:
	COMPONENT=webapp scripts/build-image.sh
