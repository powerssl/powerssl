#!/bin/bash

set -euo pipefail

DOCKERFILE=${CIRCLECI:+circleci.dockerfile}
DOCKERFILE=${DOCKERFILE:-Dockerfile}

set -x

docker build -f "${DOCKERFILE}" -t "$TAG" ..
