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
Mgoogle/api/annotations.proto=google.golang.org/genproto/googleapis/api/annotations

proto_path="proto"

tmp=$(mktemp -d)
trap 'rm -rf "$tmp"' EXIT
mkdir "$tmp/powerssl.dev"
ln -s "$PWD" "$tmp/powerssl.dev/api"

go install \
github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.6.0 \
github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.6.0
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1

PATH="$PWD/bin:${GOBIN:-"$HOME/go/bin"}:$PATH"
protoc --proto_path="$proto_path" --go_out="$mappings:$tmp" --go-grpc_out="$mappings:$tmp" --grpc-gateway_out="$tmp" --openapiv2_out="$tmp" proto/powerssl/apiserver/**/*.proto
protoc --proto_path="$proto_path" --go_out="$mappings:$tmp" --go-grpc_out="$mappings:$tmp" --grpc-gateway_out="$tmp" proto/powerssl/controller/**/*.proto
protoc --proto_path="$proto_path" --js_out=import_style=commonjs,binary:"$tmp" --grpc-web_out=import_style=commonjs,mode=grpcweb:"$tmp" proto/powerssl/apiserver/**/*.proto

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
  file=${file/$tmp/openapi}
  mkdir -p "$(dirname "$file")"
  yq eval -P "$ofile" > "${file/json/yaml}"
done <   <(find "$tmp" -type f -name '*.swagger.json' -print0)

jsdir="../../webapp/src"
while IFS= read -r -d '' file
do
  mkdir -p "$jsdir/$(dirname "${file/$tmp\/powerssl\/}")"
  sed -i -e '/var google_api_annotations_pb/ { N; d; }' "$file"
  mv "$file" "$jsdir/${file/$tmp\/powerssl\/}"
done <   <(find "$tmp" -type f -name '*.js' -print0)
