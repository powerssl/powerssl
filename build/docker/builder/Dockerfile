FROM golang:1.11-alpine3.7
RUN apk add --no-cache ca-certificates gcc git musl-dev \
 && GRPC_HEALTH_PROBE_VERSION=v0.2.0 \
 && wget -qO/usr/local/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 \
 && chmod +x /usr/local/bin/grpc_health_probe
WORKDIR /go/src/powerssl.io
ENV GO111MODULE=on
COPY go.mod .
COPY go.sum .
RUN go mod download
