ARG builder_image=powerssl/builder:latest
FROM ${builder_image} AS builder
COPY . .
RUN go build -o /usr/local/bin/powerctl powerssl.io/cmd/powerctl

FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=builder /usr/local/bin/powerctl /usr/local/bin/powerctl
ENTRYPOINT ["/usr/local/bin/powerctl"]
