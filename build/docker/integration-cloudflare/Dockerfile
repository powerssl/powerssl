ARG builder_image=powerssl/builder:latest
FROM ${builder_image} AS builder
COPY . .
RUN FORCE_REBUILD=1 GOOS=linux GOARCH=amd64 make bin/powerssl-integration-cloudflare

FROM alpine:latest
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/powerssl.io/bin/powerssl-integration-cloudflare /usr/local/bin/powerssl-integration-cloudflare
ENTRYPOINT ["/usr/local/bin/powerssl-integration-cloudflare"]
CMD ["run"]
EXPOSE 9090/tcp
