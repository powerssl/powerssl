PROTOC := $(shell which protoc)

PKG_PATH := $(abspath pkg)
PROTO_PATH := $(abspath api/protobuf-spec)

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
	github.com/gogo/protobuf/protoc-gen-gogo \
	github.com/jteeuwen/go-bindata/... \
	golang.org/x/tools/cmd/stringer


.DELETE_ON_ERROR:

.ALWAYS_REBUILD:
.PHONY: .ALWAYS_REBUILD

.DEFAULT_GOAL := all
all: build

bin/%: .ALWAYS_REBUILD
	@$(MAKE) build-${*}

.PHONY: bootstrap
bootstrap:
	@for tool in  $(EXTERNAL_TOOLS) ; do \
		echo "Installing/Updating $$tool" ; \
		GO111MODULE=off go get -u $$tool; \
	done

.PHONY: build
build: build-dev-runner build-powerctl build-powerssl-agent build-powerssl-apiserver build-powerssl-auth build-powerssl-controller build-powerssl-integration-acme build-powerssl-integration-cloudflare build-powerssl-signer build-powerssl-webapp build-powerutil

.PHONY: build-%
build-%:
	COMPONENT=${*} scripts/build-go.sh

.PHONY: build-dev-runner
build-dev-runner:
	go build -o bin/dev-runner powerssl.io/tools/dev-runner

.PHONY: clean
clean: clean-dev-runner clean-powerctl clean-powerssl-agent clean-powerssl-apiserver clean-powerssl-auth clean-powerssl-controller clean-powerssl-integration-acme clean-powerssl-integration-cloudflare clean-powerssl-signer clean-powerssl-webapp clean-powerutil

.PHONY: clean-%
clean-%:
	go clean powerssl.io/cmd/${*}
	rm -f bin/${*}

.PHONY: clean-dev-runner
clean-dev-runner:
	go clean powerssl.io/tools/dev-runner
	rm -f bin/dev-runner

.PHONY: clear-local-dev
clear-local-dev:
	rm -f local/certs/*.pem local/certs/*.csr
	rm -f local/powerssl.sqlite3
	rm -rf local/vault

.PHONY: fmt
fmt:
	go fmt $$(go list ./...)
	clang-format -i --style=Google $(PROTOS)

.PHONY: generate
generate:
	go generate $$(go list ./...)
	@$(MAKE) fmt
	@if [ "${SKIP_DOCS}" == "" ]; then \
		$(MAKE) generate-docs; \
	fi
	@if [ "${SKIP_PROTOBUF}" == "" ]; then \
		$(MAKE) generate-protobuf; \
	fi

.PHONY: generate-docs
generate-docs:
	go run powerssl.io/tools/gendocs

.PHONY: generate-protobuf
generate-protobuf:
	$(FIND_RELEVANT) -type f -name '*.pb.go' -exec rm {} +
	@rm -f powerssl.io && ln -s . powerssl.io
	set -e; for dir in $(sort $(dir $(PROTOS))); do \
		$(PROTOC) \
			-I$(PROTO_PATH):$(GOGO_GOOGLEAPIS_PATH):$(GOGO_PROTOBUF_PATH):$(PROTOBUF_PATH) \
			--gogo_out=$(PROTO_MAPPINGS),plugins=grpc:. \
			$$dir/*.proto; \
	done
	@rm -f powerssl.io

.PHONY: image-%
image-%:
	COMPONENT=${*} scripts/build-image.sh

.PHONY: images
images: image-builder image-agent image-apiserver image-auth image-controller image-envoy image-integration-acme image-integration-cloudflare image-powerctl image-signer image-webapp image-powerutil

.PHONY: install
install: install-powerctl install-powerssl-agent install-powerutil

.PHONY: install-%
install-%:
	COMPONENT=${*} scripts/install.sh

.PHONY: prepare-local-dev
prepare-local-dev:
	$(MAKE) -C local/certs

.PHONY: run
run: bin/dev-runner
	@bin/dev-runner

.PHONY: vet
vet:
	go vet $$(go list ./...)
