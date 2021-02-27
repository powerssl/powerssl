FROM golang:1.15 AS builder
WORKDIR /workspace
COPY . .
RUN FORCE_REBUILD=1 GOOS=linux GOARCH=amd64 make build-temporal

FROM alpine:latest
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /workspace/bin/powerssl-temporal /usr/local/bin/powerssl-temporal
ENTRYPOINT "/usr/local/bin/powerssl-temporal"
CMD ["serve"]
EXPOSE 8080/tcp 9090/tcp
