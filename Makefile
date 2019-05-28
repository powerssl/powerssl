.DEFAULT_GOAL := all
.DELETE_ON_ERROR:

PROTO_MAPPINGS :=
PROTO_MAPPINGS := $(PROTO_MAPPINGS)Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,
PROTO_MAPPINGS := $(PROTO_MAPPINGS)Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,
PROTO_MAPPINGS := $(PROTO_MAPPINGS)Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types,
PROTO_MAPPINGS := $(PROTO_MAPPINGS)Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,

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

bin/%:
	@$(MAKE) build-${*}

local/certs/%-key.pem local/certs/%.csr local/certs/%.pem: bin/powerutil local/certs/ca-key.pem local/certs/ca.pem
	cd local/certs && $(PWD)/bin/powerutil ca gen --ca ca.pem --ca-key ca-key.pem --hostname ${*}

local/certs/ca-key.pem local/certs/ca.csr local/certs/ca.pem: bin/powerutil
	mkdir -p local/certs
	cd local/certs && $(PWD)/bin/powerutil ca init

.PHONY: all
all: build

.PHONY: build
build: build-dev-runner build-powerctl build-powerssl-agent build-powerssl-apiserver build-powerssl-auth build-powerssl-controller build-powerssl-integration-acme build-powerssl-integration-cloudflare build-powerssl-signer build-powerssl-webapp build-powerutil

.PHONY: build-%
build-%:
	COMPONENT=${*} scripts/build-go.sh

.PHONY: build-dev-runner
build-dev-runner:
	go build -o bin/dev-runner powerssl.io/powerssl/tools/dev-runner

.PHONY: build-powerssl-apiserver
build-powerssl-apiserver:
ifneq ($(OS),Windows_NT)
ifeq ($(shell uname -s),Darwin)
	$(eval STATIC_ENABLED := 0)
endif
endif
	# NOTE: Due to sqlite3 dependency in apiserver CGO must be enabled
	CGO_ENABLED=1 COMPONENT=powerssl-apiserver STATIC_ENABLED=$(STATIC_ENABLED) scripts/build-go.sh

.PHONY: check-scripts
check-scripts:
	shellcheck scripts/*.sh

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

.PHONY: clear
clear:
	rm -rf local

.PHONY: fmt
fmt:
	go fmt $$(go list ./...)
	clang-format -i --style=Google $(call proto_files)

.PHONY: generate
generate: generate-docs generate-docs generate-protobuf

.PHONY: generate-docs
generate-docs:
	go run powerssl.io/powerssl/tools/gendocs

.PHONY: generate-go
generate-go:
	go generate $$(go list ./...)
	@$(MAKE) fmt

.PHONY: generate-protobuf
generate-protobuf:
	$(call delete_files,pkg,*.pb.go)
	$(eval $@_TMP := $(shell mktemp -d))
	mkdir $($@_TMP)/powerssl.io
	ln -s $(abspath .) $($@_TMP)/powerssl.io/powerssl
	for dir in $(call proto_dirs); do \
		protoc \
			-Iapi/protobuf-spec:$(call go_mod_dir,github.com/gogo/googleapis):$(call go_mod_dir,github.com/gogo/protobuf):$(call go_mod_dir,github.com/gogo/protobuf)/protobuf \
			--gogo_out=$(PROTO_MAPPINGS),plugins=grpc:$($@_TMP) \
			$$dir/*.proto; \
	done
	rm -rf $($@_TMP)

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

.PHONY: release-image-%
release-image-%:
	docker push powerssl/$(call strip_powerssl,${*}):latest

.PHONY: run
run: bin/dev-runner bin/powerssl-apiserver bin/powerssl-auth bin/powerssl-controller bin/powerssl-signer bin/powerssl-webapp local/certs/ca-key.pem local/certs/ca.pem local/certs/localhost-key.pem local/certs/localhost.pem local/certs/vault-key.pem local/certs/vault.pem
	@bin/dev-runner

.PHONY: test
test:
	go test $$(go list ./...)

.PHONY: test-%
test-%:
	go test $$(go list ./... | grep $(call strip_powerssl,${*}))

.PHONY: vet
vet:
	go vet $$(go list ./...)
