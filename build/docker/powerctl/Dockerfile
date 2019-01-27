ARG builder_image=powerssl/builder:latest
FROM ${builder_image} AS builder
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o /usr/local/bin/powerctl powerssl.io/cmd/powerctl

FROM alpine:latest
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/local/bin/powerctl /usr/local/bin/powerctl
ENTRYPOINT ["/usr/local/bin/powerctl"]
