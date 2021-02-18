FROM golang:1.15 AS builder
WORKDIR /workspace
COPY . .
RUN FORCE_REBUILD=1 GOOS=linux GOARCH=amd64 make build

FROM alpine:latest
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /workspace/bin/powerssl-integration-cloudflare /usr/local/bin/powerssl-integration-cloudflare
ENTRYPOINT "/usr/local/bin/powerssl-integration-cloudflare"
CMD ["run"]
EXPOSE 9090/tcp
