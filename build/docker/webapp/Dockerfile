ARG builder_image=powerssl/builder:latest
FROM ${builder_image} AS builder
COPY . .
RUN FORCE_REBUILD=1 GOOS=linux GOARCH=amd64 make bin/powerssl-webapp

FROM alpine:latest
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/powerssl.io/bin/powerssl-webapp /usr/local/bin/powerssl-webapp
ENTRYPOINT ["/usr/local/bin/powerssl-webapp"]
CMD ["serve"]
EXPOSE 8080/tcp 9090/tcp
