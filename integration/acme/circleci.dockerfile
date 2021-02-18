FROM alpine:latest AS vendor
RUN apk add --no-cache ca-certificates

FROM alpine:latest
COPY --from=vendor /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY bin/powerssl-integration-acme /usr/local/bin/powerssl-integration-acme
ENTRYPOINT "/usr/local/bin/powerssl-integration-acme"
CMD ["run"]
EXPOSE 9090/tcp
