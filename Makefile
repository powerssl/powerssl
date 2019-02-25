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

bin/protoc-gen-gogo:
	go build -o bin/protoc-gen-gogo $$(go mod download -json github.com/gogo/protobuf | grep '"Dir"' | cut -d '"' -f 4)/protoc-gen-gogo

bin/powerssl-agent: .ALWAYS_REBUILD
	COMPONENT=powerssl-agent scripts/build.sh

bin/powerssl-apiserver: .ALWAYS_REBUILD
	COMPONENT=powerssl-apiserver scripts/build.sh

bin/powerssl-auth: .ALWAYS_REBUILD
	COMPONENT=powerssl-auth scripts/build.sh

bin/powerssl-controller: .ALWAYS_REBUILD
	COMPONENT=powerssl-controller scripts/build.sh

bin/powerssl-integration-acme: .ALWAYS_REBUILD
	COMPONENT=powerssl-integration-acme scripts/build.sh

bin/powerssl-integration-cloudflare: .ALWAYS_REBUILD
	COMPONENT=powerssl-integration-cloudflare scripts/build.sh

bin/powerssl-signer: .ALWAYS_REBUILD
	COMPONENT=powerssl-signer scripts/build.sh

bin/powerssl-webapp: .ALWAYS_REBUILD
	COMPONENT=powerssl-webapp scripts/build.sh

bin/powerctl: .ALWAYS_REBUILD
	COMPONENT=powerctl scripts/build.sh

bin/powerutil: .ALWAYS_REBUILD
	COMPONENT=powerutil scripts/build.sh

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
	clang-format -i --style=google $(PROTOS)

.PHONY: vet
vet:
	go vet $$(go list ./...)

.PHONY: clean
clean:
	go clean powerssl.io/cmd/powerctl powerssl.io/cmd/powerssl-agent powerssl.io/cmd/powerssl-apiserver powerssl.io/cmd/powerssl-auth powerssl.io/cmd/powerssl-controller powerssl.io/cmd/powerssl-integration-acme powerssl.io/cmd/powerssl-integration-cloudflare powerssl.io/cmd/powerssl-signer powerssl.io/cmd/powerssl-webapp powerssl.io/cmd/powerutil
	rm -f bin/.go_protobuf_sources
	rm -f bin/powerctl bin/powerssl-agent bin/powerssl-apiserver bin/powerssl-auth bin/powerssl-controller bin/powerssl-integration-acme bin/powerssl-integration-cloudflare bin/powerssl-signer bin/powerssl-webapp bin/powerutil

.PHONY: protobuf
protobuf:
	$(FIND_RELEVANT) -type f -name '*.pb.go' -exec rm {} +
	set -e; for dir in $(sort $(dir $(PROTOS))); do \
		$(PROTOC) \
			-I$(PROTO_PATH):$(GOGO_GOOGLEAPIS_PATH):$(GOGO_PROTOBUF_PATH):$(PROTOBUF_PATH) \
			--gogo_out=$(PROTO_MAPPINGS),plugins=grpc:$(GOPATH)/src \
			$$dir/*.proto; \
	done

.PHONY: tools
tools:
	GO111MODULE=off go get golang.org/x/tools/cmd/stringer

.PHONY: generate
generate:
	go generate $$(go list ./...)

.PHONY: images
images: agent-image apiserver-image auth-image builder-image controller-image envoy-image integration-acme-image integration-cloudflare-image powerctl-image signer-image webapp-image powerutil-image

.PHONY: agent-image
agent-image:
	COMPONENT=agent scripts/docker-build.sh

.PHONY: apiserver-image
apiserver-image:
	COMPONENT=apiserver scripts/docker-build.sh

.PHONY: auth-image
auth-image:
	COMPONENT=auth scripts/docker-build.sh

.PHONY: builder-image
builder-image:
	COMPONENT=builder scripts/docker-build.sh

.PHONY: controller-image
controller-image:
	COMPONENT=controller scripts/docker-build.sh

.PHONY: envoy-image
envoy-image:
	COMPONENT=envoy scripts/docker-build.sh

.PHONY: integration-acme-image
integration-acme-image:
	COMPONENT=integration-acme scripts/docker-build.sh

.PHONY: integration-cloudflare-image
integration-cloudflare-image:
	COMPONENT=integration-cloudflare scripts/docker-build.sh

.PHONY: powerctl-image
powerctl-image:
	COMPONENT=powerctl scripts/docker-build.sh

.PHONY: powerutil-image
powerutil-image:
	COMPONENT=powerutil scripts/docker-build.sh

.PHONY: signer-image
signer-image:
	COMPONENT=signer scripts/docker-build.sh

.PHONY: webapp-image
webapp-image:
	COMPONENT=webapp scripts/docker-build.sh
