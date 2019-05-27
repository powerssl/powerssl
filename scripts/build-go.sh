#!/bin/bash

set -euo pipefail

if [ "${STATIC_ENABLED:=1}" == "0" ]; then
	unset STATIC_ENABLED
fi

set -x

env CGO_ENABLED="${CGO_ENABLED:-0}" go build ${FORCE_REBUILD:+-a} -tags netgo -ldflags "$(govvv build -flags -pkg powerssl.io/powerssl/internal/pkg/version)${DEBUG_ENABLED:+ -w}${STATIC_ENABLED:+ -extldflags \"-static\"}" -o "bin/$COMPONENT" "powerssl.io/powerssl/cmd/$COMPONENT"
