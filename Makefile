root_path := $(abspath .)
include makefiles/common.mk
include makefiles/go_bootstrap.mk

bin/powerssl-dev-runner:
	@$(MAKE) build-tool-dev-runner

bin/%:
	@$(MAKE) build-$(subst powerssl-,,$(*))

local/certs/%-key.pem local/certs/%.csr local/certs/%.pem: bin/powerutil local/certs/ca-key.pem local/certs/ca.pem
	@cd local/certs && $(BIN_PATH)/powerutil ca gen --ca ca.pem --ca-key ca-key.pem --hostname $(*)

local/certs/ca-key.pem local/certs/ca.csr local/certs/ca.pem: bin/powerutil
	@mkdir -p local/certs && cd local/certs && $(BIN_PATH)/powerutil ca init

.PHONY: build
# Build all targets
build: build-agent build-apiserver build-auth build-controller build-grpcgateway build-powerctl build-powerutil build-temporal build-tool-dev-runner build-webapp build-worker

.PHONY: build-tool-%
# Build single tool target
build-tool-%:
	@$(MAKE) -C tools/$(*) build

.PHONY: build-%
# Build single target
build-%:
	@$(MAKE) -C $(*) build

.PHONY: check-scripts
# Check scripts
check-scripts:
	@scripts/check-scripts.sh

.PHONY: clean
# Clean all targets
clean: clean-agent clean-apiserver clean-auth clean-controller clean-grpcgateway clean-powerctl clean-powerutil clean-temporal clean-tool-dev-runner clean-webapp clean-worker

.PHONY: clean-tool-%
# Clean single tool target
clean-tool-%:
	@$(MAKE) -C tools/$(*) clean

.PHONY: clean-%
# Clean single target
clean-%:
	@$(MAKE) -C $(*) build

.PHONY: clear
# Clear local resources
clear:
	@rm -rf local

.PHONY: docs-%
# Document single target
docs-%:
	@$(MAKE) -C $(*) docs

.PHONY: docs
# Document all targets
docs: docs-agent docs-apiserver docs-auth docs-controller docs-grpcgateway docs-powerctl docs-powerutil docs-temporal docs-webapp docs-worker

.PHONY: generate-circleci-config
# Generate CircleCI config
generate-circleci-config:
	@scripts/generate-circleci-config.sh

.PHONY: image-%
# Build image for single target
image-%:
	@$(MAKE) -C $(*) image

.PHONY: images
# Build image for all targets
images: image-agent image-apiserver image-auth image-controller image-envoy image-powerctl image-powerutil image-grpcgateway image-temporal image-webapp image-worker

.PHONY: install
# Install all CLI tools
install: install-powerctl install-agent install-powerutil

.PHONY: install-%
# Install single CLI tool
install-%:
	@$(MAKE) -C $(*) install

.PHONY: release-images
# Release image for all targets
images: release-image-agent release-image-apiserver release-image-auth release-image-controller release-image-envoy release-image-powerctl release-image-powerutil release-image-grpcgateway release-image-temporal release-image-webapp release-image-worker

.PHONY: release-image-%
# Release single target image
release-image-%:
	@$(MAKE) -C $(*) release-image

.PHONY: run
# Run development environment
run: bin/powerssl-dev-runner bin/powerssl-apiserver bin/powerssl-auth bin/powerssl-controller bin/powerssl-grpcgateway bin/powerssl-temporal bin/powerssl-webapp bin/powerssl-worker local/certs/ca-key.pem local/certs/ca.pem local/certs/localhost-key.pem local/certs/localhost.pem local/certs/vault-key.pem local/certs/vault.pem
	@bin/powerssl-dev-runner

.PHONY: tidy-go
# Run go mod tidy in all places
tidy-go:
	@$(SCRIPTS_PATH)/tidy-go.sh
