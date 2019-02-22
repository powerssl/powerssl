ARG builder_image=powerssl/builder:latest
FROM ${builder_image} AS builder
COPY . .
RUN FORCE_REBUILD=1 GOOS=linux GOARCH=amd64 make bin/powerutil

FROM alpine:latest
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/powerssl.io/bin/powerutil /usr/local/bin/powerutil
ENTRYPOINT ["/usr/local/bin/powerutil"]
CMD ["vault"]
