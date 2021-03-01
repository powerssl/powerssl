FROM golang:1.16 AS builder
WORKDIR /workspace
COPY . .
RUN FORCE_REBUILD=1 GOOS=linux GOARCH=amd64 make build-webapp

FROM alpine:latest
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /workspace/bin/powerssl-webapp /usr/local/bin/powerssl-webapp
ENTRYPOINT "/usr/local/bin/powerssl-webapp"
CMD ["serve"]
EXPOSE 8080/tcp 9090/tcp
