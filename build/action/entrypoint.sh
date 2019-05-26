#!/bin/sh

set -eux

cd $GITHUB_WORKSPACE
make "build-${POWERSSL_COMPONENT}"
cd bin

EVENT_DATA=$(cat $GITHUB_EVENT_PATH)
echo $EVENT_DATA | jq .
UPLOAD_URL=$(echo $EVENT_DATA | jq -r .release.upload_url)
UPLOAD_URL=${UPLOAD_URL/\{?name,label\}/}
RELEASE_NAME=$(echo $EVENT_DATA | jq -r .release.tag_name)
NAME="${POWERSSL_COMPONENT}_${RELEASE_NAME}_${GOOS}_${GOARCH}"

EXT=''
if [ "${GOOS:-}" == 'windows' ]; then
  EXT='.exe'
fi

mv ${POWERSSL_COMPONENT} "${POWERSSL_COMPONENT}${EXT}"
tar cvfz tmp.tgz "${POWERSSL_COMPONENT}${EXT}"
CHECKSUM=$(md5sum tmp.tgz | cut -d ' ' -f 1)

curl \
  -X POST \
  --data-binary @tmp.tgz \
  -H 'Content-Type: application/gzip' \
  -H "Authorization: Bearer ${GITHUB_TOKEN}" \
  "${UPLOAD_URL}?name=${NAME}.tar.gz"

curl \
  -X POST \
  --data $CHECKSUM \
  -H 'Content-Type: text/plain' \
  -H "Authorization: Bearer ${GITHUB_TOKEN}" \
  "${UPLOAD_URL}?name=${NAME}_checksum.txt"
