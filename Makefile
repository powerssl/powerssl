.DEFAULT_GOAL := all
.DELETE_ON_ERROR:

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

.PHONY: bootstrap
bootstrap:
	GO111MODULE=off go get -u github.com/myitcv/gobin

.PHONY: build
build: build-dev-runner build-powerctl build-powerssl-agent build-powerssl-apiserver build-powerssl-auth build-powerssl-controller build-powerssl-grpcgateway build-powerssl-integration-acme build-powerssl-integration-cloudflare build-powerssl-signer build-powerssl-temporalserver build-powerssl-webapp build-powerssl-worker build-powerutil

.PHONY: build-%
build-%:
	COMPONENT=${*} scripts/build-go.sh

.PHONY: build-dev-runner
build-dev-runner:
	go build -o bin/dev-runner powerssl.dev/powerssl/tools/dev-runner

.PHONY: check-scripts
check-scripts:
	shellcheck scripts/*.sh

.PHONY: clean
clean: clean-dev-runner clean-powerctl clean-powerssl-agent clean-powerssl-apiserver clean-powerssl-auth clean-powerssl-controller clean-powerssl-grpcgateway clean-powerssl-integration-acme clean-powerssl-integration-cloudflare clean-powerssl-signer clean-powerssl-temporalserver clean-powerssl-webapp clean-powerssl-worker clean-powerutil

.PHONY: clean-powerssl-integration-%
clean-powerssl-integration-%:
	cd integration/${*} && go clean powerssl.dev/integration/${*}/cmd/powerssl-integration-${*}
	rm -f bin/powerssl-integration-${*}

.PHONY: clean-%
clean-%:
	go clean powerssl.dev/powerssl/cmd/${*}
	rm -f bin/${*}

.PHONY: clean-dev-runner
clean-dev-runner:
	go clean powerssl.dev/powerssl/tools/dev-runner
	rm -f bin/dev-runner

.PHONY: clear
clear:
	rm -rf local

.PHONY: fmt
fmt:
	go fmt $$(go list ./...)
	clang-format -i --style=Google $$(find api/protobuf-spec -type f -name '*.proto' -print)

.PHONY: generate
generate: generate-protobuf generate-go generate-docs

.PHONY: generate-docs
generate-docs:
	go run powerssl.dev/powerssl/tools/gendocs

.PHONY: generate-go
generate-go:
	go generate $$(go list ./...)

.PHONY: generate-protobuf
generate-protobuf:
	scripts/generate-protobuf.sh

.PHONY: image-%
image-%:
	COMPONENT=${*} TAG=powerssl/$(call strip_powerssl,${*}):latest scripts/build-image.sh

.PHONY: images
images: image-envoy image-powerctl image-powerssl-agent image-powerssl-apiserver image-powerssl-auth image-powerssl-controller image-powerssl-grpcgateway image-powerssl-integration-acme image-powerssl-integration-cloudflare image-powerssl-signer image-powerssl-temporalserver image-powerssl-webapp image-powerssl-worker image-powerutil

.PHONY: install
install: install-powerctl install-powerssl-agent install-powerutil

.PHONY: install-%
install-%:
	COMPONENT=${*} scripts/install-go.sh

.PHONY: release-image-%
release-image-%:
	docker push powerssl/$(call strip_powerssl,${*}):latest

.PHONY: run
run: bin/dev-runner bin/powerssl-apiserver bin/powerssl-auth bin/powerssl-controller bin/powerssl-grpcgateway bin/powerssl-signer bin/powerssl-temporalserver bin/powerssl-webapp bin/powerssl-worker local/certs/ca-key.pem local/certs/ca.pem local/certs/localhost-key.pem local/certs/localhost.pem local/certs/vault-key.pem local/certs/vault.pem
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
