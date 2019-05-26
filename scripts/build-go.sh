#!/bin/bash

set -euxo pipefail

env CGO_ENABLED="${CGO_ENABLED:-0}" go build ${FORCE_REBUILD:+-a} -tags netgo -ldflags "-w -extldflags \"-static\" $(govvv build -flags -pkg powerssl.io/powerssl/internal/pkg/version)" -o "bin/$COMPONENT" "powerssl.io/powerssl/cmd/$COMPONENT"
