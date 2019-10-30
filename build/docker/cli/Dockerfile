FROM golang:1.13 AS builder
ARG COMPONENT
WORKDIR /workspace
COPY . .
RUN FORCE_REBUILD=1 GOOS=linux GOARCH=amd64 make build-${COMPONENT}

FROM alpine:latest
ARG COMPONENT
ENV COMPONENT=${COMPONENT}
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /workspace/bin/${COMPONENT} /usr/local/bin/${COMPONENT}
ENTRYPOINT "/usr/local/bin/${COMPONENT}"
