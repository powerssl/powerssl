FROM alpine:latest AS vendor
RUN apk add --no-cache ca-certificates

FROM alpine:latest
COPY --from=vendor /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY bin/powerssl-temporalserver /usr/local/bin/powerssl-temporalserver
ENTRYPOINT "/usr/local/bin/powerssl-temporalserver"
CMD ["serve"]
EXPOSE 8080/tcp 9090/tcp
