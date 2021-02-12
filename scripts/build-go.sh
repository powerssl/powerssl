#!/bin/bash

set -euo pipefail

if [ "${STATIC_ENABLED:=1}" == "0" ]; then
	unset STATIC_ENABLED
fi

case "$COMPONENT" in
	powerssl-integration-*)
	  INTEGRATION=${COMPONENT/powerssl-integration-}

	  set -x

		cd "integration/$INTEGRATION"
    env CGO_ENABLED="${CGO_ENABLED:-0}" go build ${FORCE_REBUILD:+-a} -tags netgo -ldflags "$(gobin -m -run github.com/ahmetb/govvv build -flags -pkg powerssl.dev/powerssl/internal/pkg/version)${DEBUG_ENABLED:+ -w}${STATIC_ENABLED:+ -extldflags \"-static\"}" -o "../../bin/powerssl-integration-$INTEGRATION" "powerssl.dev/integration/$INTEGRATION/cmd/powerssl-integration-$INTEGRATION"
		;;
	*)
		set -x

    env CGO_ENABLED="${CGO_ENABLED:-0}" go build ${FORCE_REBUILD:+-a} -tags netgo -ldflags "$(gobin -m -run github.com/ahmetb/govvv build -flags -pkg powerssl.dev/powerssl/internal/pkg/version)${DEBUG_ENABLED:+ -w}${STATIC_ENABLED:+ -extldflags \"-static\"}" -o "bin/$COMPONENT" "powerssl.dev/powerssl/cmd/$COMPONENT"
		;;
esac
