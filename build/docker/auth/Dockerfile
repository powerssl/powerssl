FROM golang:1.11-alpine3.7 AS builder
RUN apk add --no-cache gcc git musl-dev
WORKDIR /src/powerssl.io
COPY . .
RUN go build -o /usr/local/bin/powerssl-auth powerssl.io/cmd/powerssl-auth

FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=builder /usr/local/bin/powerssl-auth /usr/local/bin/powerssl-auth
COPY --from=builder /src/powerssl.io/pkg/auth/templates /src/powerssl.io/pkg/auth/templates
ENTRYPOINT ["/usr/local/bin/powerssl-auth"]
CMD ["serve"]
EXPOSE 8080/tcp 9090/tcp
