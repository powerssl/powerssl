root_path := $(abspath ..)
include ../makefiles/library.mk

.PHONY: fmt
# Format code
fmt:
	@clang-format -i --style=Google $$(find proto/powerssl -type f -name '*.proto' -print)

.PHONY: generate
# Generate proto stubs
generate:
	@scripts/generate-protobuf.sh
	@$(SCRIPTS_PATH)/generate-go.sh
