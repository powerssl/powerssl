FROM golang:1.11-alpine3.7 AS builder
RUN apk add --no-cache gcc git musl-dev
WORKDIR /src/powerssl.io
COPY . .
RUN go build -o /usr/local/bin/powerssl-controller powerssl.io/cmd/powerssl-controller

FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=builder /usr/local/bin/powerssl-controller /usr/local/bin/powerssl-controller
ENTRYPOINT ["/usr/local/bin/powerssl-controller"]
CMD ["serve"]
EXPOSE 8081/tcp
