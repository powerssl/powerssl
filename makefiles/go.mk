include $(mkfile_dir)/go_bootstrap.mk
include $(mkfile_dir)/go_generate.mk

.PHONY: build
# Build compiles this package
build:
	@$(SCRIPTS_PATH)/build-go.sh

.PHONY: clean
# Clean removes object files from package source directories
# It also deletes compiled binaries
clean:
	@$(SCRIPTS_PATH)/clean-go.sh

.PHONY: fmt
# Format code
fmt:
	@$(SCRIPTS_PATH)/fmt-go.sh

.PHONY: install
# Install compiles this package and installs it
install:
	@$(SCRIPTS_PATH)/install-go.sh

.PHONY: test
# Test code
test:
	@$(SCRIPTS_PATH)/test-go.sh

.PHONY: vet
# Run go vet
vet:
	@$(SCRIPTS_PATH)/vet-go.sh
