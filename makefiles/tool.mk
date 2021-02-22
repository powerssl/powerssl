mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
mkfile_dir := $(abspath $(dir $(mkfile_path)))
include $(mkfile_dir)/common.mk
include $(mkfile_dir)/go.mk

tool ?= $(package)
tool_bin ?= $(tool)

export OUTPUT=$(BIN_PATH)/$(tool_bin)
export PACKAGE = powerssl.dev/tools/$(tool)
