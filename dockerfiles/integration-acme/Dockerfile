FROM golang:1.11-alpine3.7 AS builder
RUN apk add --no-cache gcc git musl-dev
WORKDIR /src/powerssl.io
COPY . .
RUN go build -o /usr/local/bin/powerssl-integration-acme powerssl.io/cmd/powerssl-integration-acme

FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=builder /usr/local/bin/powerssl-integration-acme /usr/local/bin/powerssl-integration-acme
ENTRYPOINT ["/usr/local/bin/powerssl-integration-acme"]
CMD ["run"]
EXPOSE 9090/tcp
