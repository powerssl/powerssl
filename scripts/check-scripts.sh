#!/bin/bash

set -euo pipefail

if [ ! -d scripts ]; then
  exit 0
fi

set -x

shellcheck scripts/*.sh
