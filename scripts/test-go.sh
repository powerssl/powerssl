#!/bin/bash

set -euxo pipefail

go test "$(go list "$PACKAGE/...")"
