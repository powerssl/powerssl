mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
mkfile_dir := $(abspath $(dir $(mkfile_path)))
include $(mkfile_dir)/common.mk
include $(mkfile_dir)/go.mk

component ?= $(package)
component_bin ?= $(component)

export OUTPUT=$(BIN_PATH)/$(component_bin)
export PACKAGE = powerssl.dev/$(component)
export TAG=powerssl/$(component):latest

.PHONY: check-scripts
# Check scripts
check-scripts:
	@$(SCRIPTS_PATH)/check-scripts.sh

.PHONY: docs
# Generate docs
docs:
	@$(SCRIPTS_PATH)/generate-docs.sh

.PHONY: image
# Build image
image:
	@$(SCRIPTS_PATH)/build-image.sh

.PHONY: release-image
# Release image
release-image:
	@$(SCRIPTS_PATH)/release-image.sh
