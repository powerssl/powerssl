FROM golang:1.15 AS builder
WORKDIR /workspace
COPY . .
RUN FORCE_REBUILD=1 GOOS=linux GOARCH=amd64 make build-controller \
 && wget -qO/usr/local/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v0.2.2/grpc_health_probe-linux-amd64 \
 && chmod +x /usr/local/bin/grpc_health_probe

FROM alpine:latest
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/local/bin/grpc_health_probe /usr/local/bin/grpc_health_probe
COPY --from=builder /workspace/bin/powerssl-controller /usr/local/bin/powerssl-controller
ENTRYPOINT "/usr/local/bin/powerssl-controller"
CMD ["serve"]
EXPOSE 8080/tcp 9090/tcp
