mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
mkfile_dir := $(abspath $(dir $(mkfile_path)))
include $(mkfile_dir)/component.mk

.PHONY: run
# Run integration
run:
	@$(BIN_PATH)/$(component_bin) run --auth-token $(shell curl -ks https://localhost:8843/service) --ca-file $(root_path)/local/certs/ca.pem --controller-addr localhost:8083 --no-metrics
