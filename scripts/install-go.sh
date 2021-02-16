#!/bin/bash

set -euo pipefail

if [ "${STATIC_ENABLED:=1}" == "0" ]; then
	unset STATIC_ENABLED
fi

set -x

env CGO_ENABLED="${CGO_ENABLED:-0}" go install ${FORCE_REBUILD:+-a} -tags netgo -ldflags "$(gobin -m -run github.com/ahmetb/govvv build -flags -pkg powerssl.dev/common/version)${DEBUG_ENABLED:+ -w}${STATIC_ENABLED:+ -extldflags \"-static\"}" "powerssl.dev/powerssl/cmd/$COMPONENT"
