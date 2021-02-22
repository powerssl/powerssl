#!/bin/bash

set -euxo pipefail

go clean "$PACKAGE"
rm -f "$OUTPUT"
