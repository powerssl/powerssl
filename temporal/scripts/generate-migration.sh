#!/bin/bash

tools=(jq)
for tool in "${tools[@]}"
do
  if ! type -p "$tool" &>/dev/null
  then
    echo "Make sure you have '$tool' installed."
    exit 1
  fi
done

base=$(go mod download -json go.temporal.io/server | jq -r '.Dir')

gobin -m -run github.com/go-bindata/go-bindata/go-bindata -ignore '\.go$' -fs -modtime 726710400 -pkg migration -prefix "$base" "$base/schema/postgresql/..."