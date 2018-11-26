FROM golang:1.11-alpine3.7 AS builder
RUN apk add --no-cache gcc git musl-dev
WORKDIR /src/powerssl.io
COPY . .
RUN go build -o /usr/local/bin/powerssl-agent powerssl.io/cmd/powerssl-agent

FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=builder /usr/local/bin/powerssl-agent /usr/local/bin/powerssl-agent
ENTRYPOINT ["/usr/local/bin/powerssl-agent"]
CMD ["serve"]
