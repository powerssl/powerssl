.DELETE_ON_ERROR:

root_path ?= $(abspath $(mkfile_dir)/..)
export BIN_PATH = $(root_path)/bin
export DOCS_PATH = $(root_path)/docs
export MKFILES_PATH = $(root_path)/makefiles
export SCRIPTS_PATH = $(root_path)/scripts

include $(MKFILES_PATH)/help.mk

package_mkfile_path := $(abspath $(firstword $(MAKEFILE_LIST)))
package_dir := $(abspath $(dir $(package_mkfile_path)))
package := $(notdir $(abspath $(package_dir)))
