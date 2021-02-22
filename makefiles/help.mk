.DEFAULT_GOAL := help

.PHONY: help
# Show this help
help:
	@cat $(MAKEFILE_LIST) | docker run --rm -i xanders/make-help:latest
