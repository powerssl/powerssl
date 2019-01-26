ARG builder_image=powerssl/builder:latest
FROM ${builder_image} AS builder
COPY . .
RUN go build -o /usr/local/bin/powerssl-integration-cloudflare powerssl.io/cmd/powerssl-integration-cloudflare

FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=builder /usr/local/bin/powerssl-integration-cloudflare /usr/local/bin/powerssl-integration-cloudflare
ENTRYPOINT ["/usr/local/bin/powerssl-integration-cloudflare"]
CMD ["run"]
EXPOSE 9090/tcp
