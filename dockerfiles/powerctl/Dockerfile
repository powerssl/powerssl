FROM golang:1.11-alpine3.7 AS builder
RUN apk add --no-cache gcc git musl-dev
WORKDIR /src/powerssl.io
COPY . .
RUN go build -o /usr/local/bin/powerctl powerssl.io/cmd/powerctl

FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=builder /usr/local/bin/powerctl /usr/local/bin/powerctl
ENTRYPOINT ["/usr/local/bin/powerctl"]
