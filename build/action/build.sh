#!/bin/sh

set -eux

EXT=''
if [ $GOOS == 'windows' ]; then
  EXT='.exe'
fi

ln -s $GITHUB_WORKSPACE /go/src/powerssl.io
cd /go/src/powerssl.io

CGO_ENABLED=0 GO111MODULE=on go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o /go/bin/$POWERSSL_COMPONENT$EXT powerssl.io/cmd/$POWERSSL_COMPONENT
