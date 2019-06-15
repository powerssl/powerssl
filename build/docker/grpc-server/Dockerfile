FROM golang:1.12 AS builder
ARG COMPONENT
WORKDIR /workspace
COPY . .
RUN FORCE_REBUILD=1 GOOS=linux GOARCH=amd64 make build-${COMPONENT} \
 && wget -qO/usr/local/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v0.2.2/grpc_health_probe-linux-amd64 \
 && chmod +x /usr/local/bin/grpc_health_probe

FROM alpine:latest
ARG COMPONENT
ENV COMPONENT=${COMPONENT}
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/local/bin/grpc_health_probe /usr/local/bin/grpc_health_probe
COPY --from=builder /workspace/bin/${COMPONENT} /usr/local/bin/${COMPONENT}
ENTRYPOINT "/usr/local/bin/${COMPONENT}"
CMD ["serve"]
EXPOSE 8080/tcp 9090/tcp
