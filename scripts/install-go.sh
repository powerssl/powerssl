#!/bin/bash

set -euo pipefail

"$(dirname "$0")/build-go.sh"

GOPATH=${GOPATH:-$HOME/go}
GOBIN=${GOBIN:-$GOPATH/bin}

set -x

cp -f "$OUTPUT" "$GOBIN"
