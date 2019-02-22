ARG builder_image=powerssl/builder:latest
FROM ${builder_image} AS builder
COPY . .
RUN FORCE_REBUILD=1 GOOS=linux GOARCH=amd64 make bin/powerssl-auth

FROM alpine:latest
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/powerssl.io/bin/powerssl-auth /usr/local/bin/powerssl-auth
ENTRYPOINT ["/usr/local/bin/powerssl-auth"]
CMD ["serve"]
EXPOSE 8080/tcp 9090/tcp
