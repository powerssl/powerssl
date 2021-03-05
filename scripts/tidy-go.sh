#!/bin/bash

set -euo pipefail

ROOT=$PWD
FILES=$(find . -type f -name go.mod)
for f in $FILES
do
 cd "$ROOT/$(dirname "$f")"
 go mod tidy
done