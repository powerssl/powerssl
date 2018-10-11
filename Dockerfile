FROM golang:1.11-alpine3.7 AS builder
RUN apk add --no-cache gcc git musl-dev
ENV GO111MODULE=on
WORKDIR /go/src/powerssl.io
COPY . .
RUN go build -o /usr/local/bin/powerctl powerssl.io/cmd/powerctl \
 && go build -o /usr/local/bin/powerssl-apiserver powerssl.io/cmd/powerssl-apiserver

FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=builder /usr/local/bin/powerctl /usr/local/bin/powerctl
COPY --from=builder /usr/local/bin/powerssl-apiserver /usr/local/bin/powerssl-apiserver
