FROM golang:1.11-alpine

LABEL "com.github.actions.name"="Go Release Binary"
LABEL "com.github.actions.description"="Automate publishing Go build artifacts for GitHub releases"
LABEL "com.github.actions.icon"="cpu"
LABEL "com.github.actions.color"="orange"

RUN apk add --no-cache ca-certificates curl gcc git jq musl-dev

ADD entrypoint.sh /entrypoint.sh
ADD build.sh /build.sh
ENTRYPOINT ["/entrypoint.sh"]
