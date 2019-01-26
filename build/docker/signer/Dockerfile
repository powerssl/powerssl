ARG builder_image=powerssl/builder:latest
FROM ${builder_image} AS builder
COPY . .
RUN go build -o /usr/local/bin/powerssl-signer powerssl.io/cmd/powerssl-signer

FROM alpine:latest
RUN apk add --no-cache ca-certificates \
 && GRPC_HEALTH_PROBE_VERSION=v0.2.0 \
 && wget -qO/usr/local/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 \
 && chmod +x /usr/local/bin/grpc_health_probe
COPY --from=builder /usr/local/bin/powerssl-signer /usr/local/bin/powerssl-signer
ENTRYPOINT ["/usr/local/bin/powerssl-signer"]
CMD ["serve"]
EXPOSE 8080/tcp 9090/tcp
