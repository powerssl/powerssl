#!/bin/bash

set -euxo pipefail

go fmt "$(go list "$PACKAGE/...")"
