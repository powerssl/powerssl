ARG builder_image=powerssl/builder:latest
FROM ${builder_image} AS builder
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o /usr/local/bin/powerssl-webapp powerssl.io/cmd/powerssl-webapp

FROM alpine:latest
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/local/bin/powerssl-webapp /usr/local/bin/powerssl-webapp
COPY --from=builder /go/src/powerssl.io/web/app /go/src/powerssl.io/web/app
ENTRYPOINT ["/usr/local/bin/powerssl-webapp"]
CMD ["serve"]
EXPOSE 8080/tcp 9090/tcp
