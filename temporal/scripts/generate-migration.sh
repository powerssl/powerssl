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
basedir=$(basename "$(dirname "$base")")
basename="$basedir/$(basename "$base")"

trap 'rm -rf "$basedir"' EXIT
mkdir "$basedir"
ln -s "$base" "$basename"

go run github.com/go-bindata/go-bindata/v3/go-bindata@v3.1.3 -ignore '\.go$' -fs -modtime 726710400 -pkg migration -prefix "$basename" "$basename/schema/postgresql/..."
