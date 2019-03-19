#!/bin/sh

set -eux

FORCE_REBUILD_ARG=''
if [ -n "${FORCE_REBUILD:-}" ]; then
  FORCE_REBUILD_ARG='-a'
fi

EXTLDFLAGS='-static'
EXT=''

if [ -z "${GOOS:-}" ]; then
  case "${OSTYPE:-}" in
    darwin*) EXTLDFLAGS='' ;;
  esac
elif [ "${GOOS}" == 'windows' ]; then
  EXT='.exe'
fi

CGO_ENABLED="${CGO_ENABLED:-0}"
go build ${FORCE_REBUILD_ARG} -tags netgo -ldflags "-w -extldflags \"${EXTLDFLAGS}\"" -o "bin/${COMPONENT}${EXT}" "powerssl.io/cmd/${COMPONENT}"
