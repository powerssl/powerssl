#!/bin/bash

set -euxo pipefail

go generate "$(go list "$PACKAGE/...")"
