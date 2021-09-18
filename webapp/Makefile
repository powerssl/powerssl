component_bin := powerssl-webapp

include ../makefiles/component.mk

.PHONY: copy
# Copy webapp resources
copy:
	cp -rf $(root_path)/../webapp/dist/* internal/asset
