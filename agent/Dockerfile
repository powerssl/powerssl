FROM golang:1.16 AS builder
WORKDIR /workspace
COPY . .
RUN FORCE_REBUILD=1 GOOS=linux GOARCH=amd64 make build-agent

FROM alpine:latest
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /workspace/bin/powerssl-agent /usr/local/bin/powerssl-agent
ENTRYPOINT "/usr/local/bin/powerssl-agent"
CMD ["run"]
