FROM golang:1.11-alpine3.7 AS builder-base
RUN apk add --no-cache gcc git musl-dev
WORKDIR /go/src/powerssl.io
ENV GO111MODULE=on
COPY go.mod .
COPY go.sum .
RUN go mod download
