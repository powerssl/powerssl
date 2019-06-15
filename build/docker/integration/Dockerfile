FROM golang:1.12 AS builder
ARG INTEGRATION
WORKDIR /workspace
COPY . .
RUN FORCE_REBUILD=1 GOOS=linux GOARCH=amd64 make build-powerssl-integration-${INTEGRATION}

FROM alpine:latest
ARG INTEGRATION
ENV INTEGRATION=${INTEGRATION}
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /workspace/bin/powerssl-integration-${INTEGRATION} /usr/local/bin/powerssl-integration-${INTEGRATION}
ENTRYPOINT "/usr/local/bin/powerssl-integration-${INTEGRATION}"
CMD ["run"]
EXPOSE 9090/tcp
