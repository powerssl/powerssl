ARG builder_image=powerssl/builder:latest
FROM ${builder_image} AS builder
COPY . .
RUN go build -o /usr/local/bin/powerssl-integration-acme powerssl.io/cmd/powerssl-integration-acme

FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=builder /usr/local/bin/powerssl-integration-acme /usr/local/bin/powerssl-integration-acme
ENTRYPOINT ["/usr/local/bin/powerssl-integration-acme"]
CMD ["run"]
EXPOSE 9090/tcp
