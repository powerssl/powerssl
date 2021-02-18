FROM alpine:latest AS vendor
RUN apk add --no-cache ca-certificates

FROM alpine:latest
COPY --from=vendor /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY bin/powerssl-agent /usr/local/bin/powerssl-agent
ENTRYPOINT "/usr/local/bin/powerssl-agent"
CMD ["run"]
