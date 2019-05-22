EXTERNAL_TOOLS=\
	github.com/gogo/protobuf/protoc-gen-gogo \
	github.com/jteeuwen/go-bindata/... \
	golang.org/x/tools/cmd/stringer

PROTO_MAPPINGS :=
PROTO_MAPPINGS := $(PROTO_MAPPINGS)Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,
PROTO_MAPPINGS := $(PROTO_MAPPINGS)Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,
PROTO_MAPPINGS := $(PROTO_MAPPINGS)Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types,
PROTO_MAPPINGS := $(PROTO_MAPPINGS)Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,

.DELETE_ON_ERROR:

.ALWAYS_REBUILD:
.PHONY: .ALWAYS_REBUILD

define delete_files
$(shell find $(1) -type f -name '$(2)' -exec rm {} +)
endef

define files
$(sort $(shell find $(1) -type f -name '$(2)' -print))
endef

define go_mod_dir
$(shell go mod download -json $(1) | grep \"Dir\" | cut -d \" -f 4)
endef

define proto_dirs
$(sort $(dir $(call proto_files)))
endef

define proto_files
$(call files,api/protobuf-spec,*.proto)
endef

define strip_powerssl
$(shell echo ${*} | sed 's/powerssl-//')
endef

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
	go build -o bin/dev-runner powerssl.io/powerssl/tools/dev-runner

.PHONY: clean
clean: clean-dev-runner clean-powerctl clean-powerssl-agent clean-powerssl-apiserver clean-powerssl-auth clean-powerssl-controller clean-powerssl-integration-acme clean-powerssl-integration-cloudflare clean-powerssl-signer clean-powerssl-webapp clean-powerutil

.PHONY: clean-%
clean-%:
	go clean powerssl.io/powerssl/cmd/${*}
	rm -f bin/${*}

.PHONY: clean-dev-runner
clean-dev-runner:
	go clean powerssl.io/powerssl/tools/dev-runner
	rm -f bin/dev-runner

.PHONY: clear-local-dev
clear-local-dev:
	rm -f local/certs/*.pem local/certs/*.csr
	rm -f local/powerssl.sqlite3
	rm -rf local/vault

.PHONY: fmt
fmt:
	go fmt $$(go list ./...)
	clang-format -i --style=Google $(call proto_files)

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
	go run powerssl.io/powerssl/tools/gendocs

.PHONY: generate-protobuf
generate-protobuf:
	$(call delete_files,pkg,*.pb.go)
	@rm -f powerssl.io && ln -s . powerssl.io
	set -e; for dir in $(call proto_dirs); do \
		protoc \
			-Iapi/protobuf-spec:$(call go_mod_dir,github.com/gogo/googleapis):$(call go_mod_dir,github.com/gogo/protobuf):$(call go_mod_dir,github.com/gogo/protobuf)/protobuf \
			--gogo_out=$(PROTO_MAPPINGS),plugins=grpc:. \
			$$dir/*.proto; \
	done
	@rm -f powerssl.io

.PHONY: image-%
image-%:
	COMPONENT=${*} TAG=powerssl/$(call strip_powerssl,${*}):latest scripts/build-image.sh

.PHONY: images
images: image-envoy image-powerctl image-powerssl-agent image-powerssl-apiserver image-powerssl-auth image-powerssl-controller image-powerssl-integration-acme image-powerssl-integration-cloudflare image-powerssl-signer image-powerssl-webapp image-powerutil

.PHONY: install
install: install-powerctl install-powerssl-agent install-powerutil

.PHONY: install-%
install-%:
	COMPONENT=${*} scripts/install.sh

.PHONY: prepare-local-dev
prepare-local-dev:
	$(MAKE) -C local/certs

.PHONY: release-image-%
release-image-%:
	docker push powerssl/$(call strip_powerssl,${*}):latest

.PHONY: run
run: bin/dev-runner
	@bin/dev-runner

.PHONY: vet
vet:
	go vet $$(go list ./...)
