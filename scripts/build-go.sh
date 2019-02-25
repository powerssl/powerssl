#!/bin/sh

set -eux

EXT=''
if [ "${GOOS:-}" == 'windows' ]; then
  EXT='.exe'
fi

FORCE_REBUILD_ARG=''
if [ -n "${FORCE_REBUILD:-}" ]; then
  FORCE_REBUILD_ARG='-a'
fi

CGO_ENABLED=0 go build ${FORCE_REBUILD_ARG} -tags netgo -ldflags '-w -extldflags "-static"' -o "bin/${COMPONENT}${EXT}" "powerssl.io/cmd/${COMPONENT}"
