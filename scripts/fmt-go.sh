#!/bin/bash

set -euxo pipefail

go fmt "$PACKAGE/..."
