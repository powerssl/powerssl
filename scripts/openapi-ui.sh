#!/bin/bash

set -euo pipefail

urls=()
while IFS= read -r -d '' file
do
  urls+=("{url:\"$GRPC_PROXY_URI/${file/api\/openapi\/powerssl\/apiserver/openapi}\",name:\"${file/api\/openapi\/}\"}")
done <   <(find api/openapi/powerssl/apiserver/v1 -type f -name '*.yaml' -print0)
URLS=${urls[*]}
URLS="[${URLS// /,}]"

docker run --rm -e URLS="$URLS" -e VALIDATOR_URL=null -p 80:8080 swaggerapi/swagger-ui
