#!/bin/sh

set -eux

FORCE_REBUILD_ARG=''
if [ -n "${FORCE_REBUILD:-}" ]; then
  FORCE_REBUILD_ARG='-a'
fi

CGO_ENABLED=0 go install ${FORCE_REBUILD_ARG} -tags netgo -ldflags '-w -extldflags "-static"' "powerssl.io/powerssl/cmd/${COMPONENT}"
