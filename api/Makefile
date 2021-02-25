include ../makefiles/help.mk

.PHONY: fmt
# Format code
fmt:
	@clang-format -i --style=Google $$(find proto/powerssl -type f -name '*.proto' -print)

.PHONY: generate
# Generate proto stubs
generate:
	@scripts/generate-protobuf.sh
