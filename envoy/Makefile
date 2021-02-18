include ../makefiles/help.mk

.PHONY: image
# Build image
image:
	docker build -f Dockerfile -t powerssl/envoy:latest ..

.PHONY: release-image
# Release image
release-image:
	docker push powerssl/envoy:latest
