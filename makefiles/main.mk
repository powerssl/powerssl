.DEFAULT_GOAL := help
.DELETE_ON_ERROR:

component_path := $(dir $(abspath $(firstword $(MAKEFILE_LIST))))
component ?= $(notdir $(abspath $(component_path)))
component_bin ?= $(component)
#sanitized_component := $(subst powerssl-,,$(component))
mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
mkfiles_path := $(abspath $(dir $(abspath $(mkfile_path))))
root_path := $(abspath $(mkfiles_path)/..)
bin_path := $(root_path)/bin
docs_path := $(root_path)/docs
scripts_path := $(root_path)/scripts
#current_dir := $(notdir $(patsubst %/,%,$(dir $(mkfile_path))))

export PACKAGE = powerssl.dev/$(component)

include $(mkfiles_path)/help.mk

$(bin_path)/${component}: build

.PHONY: build
# Build compiles this package
build:
	@OUTPUT=$(bin_path)/$(component_bin) $(scripts_path)/build-go.sh

.PHONY: clean
# Clean removes object files from package source directories
# It also deletes compiled binaries
clean:
	@go clean $(PACKAGE)
	@rm -f $(bin_path)/$(component_bin)

.PHONY: docs
# Generate docs
docs:
	@go run $(PACKAGE)/tools/gendocs --dir $(docs_path)

.PHONY: image
# Build image
image:
	@TAG=powerssl/$(component):latest $(scripts_path)/build-image.sh

.PHONY: release-image
# Release image
release-image:
	@docker push powerssl/$(component):latest
