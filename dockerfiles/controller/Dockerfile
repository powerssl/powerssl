FROM golang:1.11-alpine3.7 AS builder
RUN apk add --no-cache gcc git musl-dev
WORKDIR /src/powerssl.io
COPY . .
RUN go build -o /usr/local/bin/powerssl-controller powerssl.io/cmd/powerssl-controller

FROM alpine:latest
RUN apk add --no-cache ca-certificates \
 && GRPC_HEALTH_PROBE_VERSION=v0.2.0 \
 && wget -qO/usr/local/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 \
 && chmod +x /usr/local/bin/grpc_health_probe
COPY --from=builder /usr/local/bin/powerssl-controller /usr/local/bin/powerssl-controller
ENTRYPOINT ["/usr/local/bin/powerssl-controller"]
CMD ["serve"]
EXPOSE 8080/tcp 9090/tcp
