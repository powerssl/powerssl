ARG builder_image=powerssl/builder:latest
FROM ${builder_image} AS builder
COPY . .
RUN go build -o /usr/local/bin/powerssl-agent powerssl.io/cmd/powerssl-agent

FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=builder /usr/local/bin/powerssl-agent /usr/local/bin/powerssl-agent
ENTRYPOINT ["/usr/local/bin/powerssl-agent"]
CMD ["serve"]
