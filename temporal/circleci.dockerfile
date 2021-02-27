FROM alpine:latest AS vendor
RUN apk add --no-cache ca-certificates

FROM alpine:latest
COPY --from=vendor /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY bin/powerssl-temporal /usr/local/bin/powerssl-temporal
ENTRYPOINT "/usr/local/bin/powerssl-temporal"
CMD ["serve"]
EXPOSE 8080/tcp 9090/tcp
