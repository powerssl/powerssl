#!/bin/bash

set -euxo pipefail

go vet "$(go list "$PACKAGE/...")"
