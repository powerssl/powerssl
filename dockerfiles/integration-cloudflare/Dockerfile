FROM golang:1.11-alpine3.7 AS builder
RUN apk add --no-cache gcc git musl-dev
WORKDIR /src/powerssl.io
COPY . .
RUN go build -o /usr/local/bin/powerssl-integration-cloudflare powerssl.io/cmd/powerssl-integration-cloudflare

FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=builder /usr/local/bin/powerssl-integration-cloudflare /usr/local/bin/powerssl-integration-cloudflare
ENTRYPOINT ["/usr/local/bin/powerssl-integration-cloudflare"]
CMD ["run"]
EXPOSE 9090/tcp
