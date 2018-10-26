FROM golang:1.11-alpine3.7 AS builder
RUN apk add --no-cache gcc git musl-dev
WORKDIR /powerssl.io
COPY . .
RUN go build -o bin/powerctl powerssl.io/cmd/powerctl \
 && go build -o bin/powerssl-apiserver powerssl.io/cmd/powerssl-apiserver \
 && go build -o bin/powerssl-controller powerssl.io/cmd/powerssl-controller \
 && go build -o bin/powerssl-integration-acme powerssl.io/cmd/powerssl-integration-acme \
 && go build -o bin/powerssl-integration-cloudflare powerssl.io/cmd/powerssl-integration-cloudflare \
 && go build -o bin/powerssl-signer powerssl.io/cmd/powerssl-signer

FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=builder /powerssl.io/bin/powerctl /powerssl.io/bin/powerssl-apiserver /powerssl.io/bin/powerssl-controller /powerssl.io/bin/powerssl-integration-acme /powerssl.io/bin/powerssl-integration-cloudflare /powerssl.io/bin/powerssl-signer /usr/local/bin/
