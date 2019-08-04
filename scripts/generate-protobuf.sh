#!/bin/bash

set -Eeuo pipefail

tools=(jq protoc swagger2openapi yq)
for tool in "${tools[@]}"
do
  if ! type -p "$tool" &>/dev/null
  then
    echo "Make sure you have '$tool' installed."
    exit 1
  fi
done

mappings=\
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types

googleapis=$(go mod download -json github.com/gogo/googleapis | grep '"Dir"' | cut -d '"' -f 4)
protobuf=$(go mod download -json github.com/gogo/protobuf | grep '"Dir"' | cut -d '"' -f 4)
proto_path="api/protobuf-spec:$googleapis:$protobuf:$protobuf/protobuf"

tmp=$(mktemp -d)
trap 'rm -rf "$tmp"' EXIT
mkdir "$tmp/powerssl.dev"
ln -s "$PWD" "$tmp/powerssl.dev/powerssl"

PATH="$(dirname "$(gobin -m -p github.com/gogo/protobuf/protoc-gen-gogo)"):$PATH"
PATH="$(dirname "$(gobin -m -p github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway)"):$PATH"
PATH="$(dirname "$(gobin -m -p github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger)"):$PATH"
protoc --proto_path="$proto_path" --gogo_out="$mappings,plugins=grpc:$tmp" --grpc-gateway_out="logtostderr=true:$tmp" --swagger_out=logtostderr="true:$tmp" api/protobuf-spec/powerssl/apiserver/**/*.proto
protoc --proto_path="$proto_path" --gogo_out="$mappings,plugins=grpc:$tmp" --grpc-gateway_out="logtostderr=true:$tmp" api/protobuf-spec/powerssl/controller/**/*.proto

while IFS= read -r -d '' file
do
  for path in $(jq -r '.paths | to_entries[] | select(.key | test("=")) | .key' < "$file")
  do
    # shellcheck disable=SC2001
    jq \
      --arg example "$(echo "$path" | sed 's/^[^=]*=\([^\}]*\).*$/\1/')" \
      --arg name "$(echo "$path" | sed 's/.*{\([^=]*\).*/\1/')" \
      --arg path "$path" \
      --arg pattern "^$(echo "$path" | sed 's/^[^=]*=\([^\}]*\).*$/\1/' | sed 's/*/\\w+/g' | sed 's/\//\\\//g')$" \
      --arg replace "$(echo "$path" | sed 's/^\([^=]*\)=[^\}]*\(.*\)$/\1\2/')" \
      '.info.version = "v1" | .paths = (.paths | with_entries(if .key == $path then .key = $replace else . end)) | .paths[$replace] = (.paths[$replace] | to_entries | map(.value = (.value | to_entries | map(.value = if .key == "parameters" then .value | map(. + if .name == $name then {"example":$example,"schema":{"type":"string","pattern":$pattern}} else {} end) else .value end) | from_entries)) | from_entries)' \
      "$file" > "$file.tmp"
    mv "$file.tmp" "$file"
  done

  ofile=${file/.swagger}
  swagger2openapi -o "$ofile" "$file"
  jq '.components.securitySchemes.bearerAuth = {"type":"http","scheme":"bearer","bearerFormat":"JWT"} | .paths[][].security = [{"bearerAuth":[]}]' "$ofile" > "$ofile.tmp"
  mv "$ofile.tmp" "$ofile"
  file=${file/.swagger}
  file=${file/$tmp/api/openapi}
  mkdir -p "$(dirname "$file")"
  yq r "$ofile" > "${file/json/yaml}"
done <   <(find "$tmp" -type f -name '*.swagger.json' -print0)
