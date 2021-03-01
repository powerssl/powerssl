FROM golang:1.16 AS builder
WORKDIR /workspace
COPY . .
RUN FORCE_REBUILD=1 GOOS=linux GOARCH=amd64 make build-worker

FROM alpine:latest
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /workspace/bin/powerssl-worker /usr/local/bin/powerssl-worker
ENTRYPOINT "/usr/local/bin/powerssl-worker"
CMD ["run"]
EXPOSE 9090/tcp
