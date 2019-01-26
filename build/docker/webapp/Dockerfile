ARG builder_image=powerssl/builder:latest
FROM ${builder_image} AS builder
COPY . .
RUN go build -o /usr/local/bin/powerssl-webapp powerssl.io/cmd/powerssl-webapp

FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=builder /usr/local/bin/powerssl-webapp /usr/local/bin/powerssl-webapp
COPY --from=builder /go/src/powerssl.io/pkg/webapp/templates /go/src/powerssl.io/pkg/webapp/templates
ENTRYPOINT ["/usr/local/bin/powerssl-webapp"]
CMD ["serve"]
EXPOSE 8080/tcp 9090/tcp
