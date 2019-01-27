ARG builder_image=powerssl/builder:latest
FROM ${builder_image} AS builder
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o /usr/local/bin/powerssl-integration-acme powerssl.io/cmd/powerssl-integration-acme

FROM alpine:latest
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/local/bin/powerssl-integration-acme /usr/local/bin/powerssl-integration-acme
ENTRYPOINT ["/usr/local/bin/powerssl-integration-acme"]
CMD ["run"]
EXPOSE 9090/tcp
