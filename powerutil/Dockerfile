FROM golang:1.15 AS builder
WORKDIR /workspace
COPY . .
RUN FORCE_REBUILD=1 GOOS=linux GOARCH=amd64 make build-powerutil

FROM alpine:latest
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /workspace/bin/powerutil /usr/local/bin/powerutil
ENTRYPOINT "/usr/local/bin/powerutil"
