FROM alpine:latest AS vendor
RUN apk add --no-cache ca-certificates \
 && wget -qO/usr/local/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v0.2.2/grpc_health_probe-linux-amd64 \
 && chmod +x /usr/local/bin/grpc_health_probe

FROM alpine:latest
COPY --from=vendor /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=vendor /usr/local/bin/grpc_health_probe /usr/local/bin/grpc_health_probe
COPY bin/powerssl-controller /usr/local/bin/powerssl-controller
ENTRYPOINT "/usr/local/bin/powerssl-controller"
CMD ["serve"]
EXPOSE 8080/tcp 9090/tcp
