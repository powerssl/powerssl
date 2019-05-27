#!/bin/bash

set -euxo pipefail

env CGO_ENABLED="${CGO_ENABLED:-0}" go install ${FORCE_REBUILD:+-a} -tags netgo -ldflags "-w -extldflags \"-static\" $(govvv install -flags)" "powerssl.io/powerssl/cmd/$COMPONENT"
