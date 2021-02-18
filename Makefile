include makefiles/help.mk

.DELETE_ON_ERROR:

makefile_path := $(abspath $(dir $(abspath $(firstword $(MAKEFILE_LIST)))))

bin/%:
	@$(MAKE) build-$(subst powerssl-,,${*})

local/certs/%-key.pem local/certs/%.csr local/certs/%.pem: bin/powerutil local/certs/ca-key.pem local/certs/ca.pem
	cd local/certs && $(makefile_path)/bin/powerutil ca gen --ca ca.pem --ca-key ca-key.pem --hostname ${*}

local/certs/ca-key.pem local/certs/ca.csr local/certs/ca.pem: bin/powerutil
	mkdir -p local/certs
	cd local/certs && $(makefile_path)/bin/powerutil ca init

.PHONY: bootstrap
# Bootstrap development environment
bootstrap:
	GO111MODULE=off go get -u github.com/myitcv/gobin

.PHONY: build
# Build all targets
build: build-agent build-apiserver build-auth build-controller build-dev-runner build-grpcgateway build-powerctl build-powerutil build-signer build-temporalserver build-webapp build-worker

.PHONY: build-%
# Build single target
build-%:
	@$(MAKE) -C ${*} build

.PHONY: build-dev-runner
build-dev-runner:
	cd tools/dev-runner && go build -o ../../bin/dev-runner powerssl.dev/tools/dev-runner

.PHONY: check-scripts
# Check schripts
check-scripts:
	shellcheck scripts/*.sh

.PHONY: clean
# Clean all targets
clean: clean-agent clean-apiserver clean-auth clean-controller clean-dev-runner clean-grpcgateway clean-powerctl clean-powerutil clean-signer clean-temporalserver clean-webapp clean-worker

.PHONY: clean-%
# Clean single target
clean-%:
	@$(MAKE) -C ${*} build

.PHONY: clear
# Clear local resources
clear:
	rm -rf local

.PHONY: docs-%
# Document single target
docs-%:
	@$(MAKE) -C ${*} docs

.PHONY: fmt
# Format code
fmt:
	go fmt $$(go list ./...)
	clang-format -i --style=Google $$(find api/protobuf-spec -type f -name '*.proto' -print)

.PHONY: generate
# Generate all targets
generate: generate-protobuf generate-go generate-docs

.PHONY: generate-docs
# Document all targets
generate-docs: docs-agent docs-apiserver docs-auth docs-controller docs-grpcgateway docs-powerctl docs-powerutil docs-signer docs-temporalserver docs-webapp docs-worker

.PHONY: generate-go
# Generate go code
generate-go:
	go generate $$(go list ./...)

.PHONY: generate-protobuf
# Generate protobuf
generate-protobuf:
	scripts/generate-protobuf.sh

.PHONY: image-%
# Build image for single target
image-%:
	@$(MAKE) -C ${*} image

.PHONY: images
# Build image for all targets
images: image-agent image-apiserver image-auth image-controller image-envoy image-powerctl image-powerutil image-grpcgateway image-signer image-temporalserver image-webapp image-worker

.PHONY: install
# Install all CLI tools
install: install-powerctl install-powerssl-agent install-powerutil

.PHONY: install-%
# Install single CLI tool
install-%:
	@$(MAKE) -C ${*} install

.PHONY: release-image-%
# Release single target image
release-image-%:
	@$(MAKE) -C ${*} release-image

.PHONY: run
# Run development environment
run: bin/dev-runner bin/powerssl-apiserver bin/powerssl-auth bin/powerssl-controller bin/powerssl-grpcgateway bin/powerssl-signer bin/powerssl-temporalserver bin/powerssl-webapp bin/powerssl-worker local/certs/ca-key.pem local/certs/ca.pem local/certs/localhost-key.pem local/certs/localhost.pem local/certs/vault-key.pem local/certs/vault.pem
	@bin/dev-runner

.PHONY: test
# Test code for all targets
test:
	go test $$(go list ./...)

.PHONY: test-%
# Test code for single target
test-%:
	go test $$(go list ./... | grep $(subst powerssl-,,${*}))

.PHONY: vet
# Run go vet for all targets
vet:
	go vet $$(go list ./...)
