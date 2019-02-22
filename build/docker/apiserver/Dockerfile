ARG builder_image=powerssl/builder:latest
FROM ${builder_image} AS builder
COPY . .
RUN FORCE_REBUILD=1 GOOS=linux GOARCH=amd64 make bin/powerssl-apiserver

FROM alpine:latest
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/local/bin/grpc_health_probe /usr/local/bin/grpc_health_probe
COPY --from=builder /go/src/powerssl.io/bin/powerssl-apiserver /usr/local/bin/powerssl-apiserver
ENTRYPOINT ["/usr/local/bin/powerssl-apiserver"]
CMD ["serve"]
EXPOSE 8080/tcp 9090/tcp
