FROM golang:1.11-alpine3.7 AS builder
RUN apk add --no-cache gcc git musl-dev
WORKDIR /src/powerssl.io
COPY . .
RUN go build -o /usr/local/bin/powerssl-signer powerssl.io/cmd/powerssl-signer

FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=builder /usr/local/bin/powerssl-signer /usr/local/bin/powerssl-signer
ENTRYPOINT ["/usr/local/bin/powerssl-signer"]
CMD ["serve"]
EXPOSE 8080/tcp 9090/tcp
