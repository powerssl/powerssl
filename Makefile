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
	docker build -f build/docker/agent/Dockerfile -t powerssl/agent .

.PHONY: apiserver-image
apiserver-image:
	docker build -f build/docker/apiserver/Dockerfile -t powerssl/apiserver .

.PHONY: auth-image
auth-image:
	docker build -f build/docker/auth/Dockerfile -t powerssl/auth .

.PHONY: builder-image
builder-image:
	docker build -f build/docker/builder/Dockerfile -t powerssl/builder .

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

.PHONY: powerutil-image
powerutil-image:
	docker build -f build/docker/powerutil/Dockerfile -t powerssl/powerutil .

.PHONY: signer-image
signer-image:
	docker build -f build/docker/signer/Dockerfile -t powerssl/signer .

.PHONY: webapp-image
webapp-image:
	docker build -f build/docker/webapp/Dockerfile -t powerssl/webapp .
