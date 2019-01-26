ARG builder_image=powerssl/builder:latest
FROM ${builder_image} AS builder
COPY . .
RUN go build -o /usr/local/bin/powerssl-auth powerssl.io/cmd/powerssl-auth

FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=builder /usr/local/bin/powerssl-auth /usr/local/bin/powerssl-auth
COPY --from=builder /go/src/powerssl.io/pkg/auth/templates /go/src/powerssl.io/pkg/auth/templates
ENTRYPOINT ["/usr/local/bin/powerssl-auth"]
CMD ["serve"]
EXPOSE 8080/tcp 9090/tcp
