#!/bin/bash

set -euxo pipefail

go vet "$PACKAGE/..."
