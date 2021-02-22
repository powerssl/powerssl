mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
mkfile_dir := $(abspath $(dir $(mkfile_path)))
include $(mkfile_dir)/common.mk
include $(mkfile_dir)/go_generate.mk

library ?= $(package)

export PACKAGE = powerssl.dev/$(library)
