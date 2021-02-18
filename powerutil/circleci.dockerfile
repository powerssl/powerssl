FROM alpine:latest AS vendor
RUN apk add --no-cache ca-certificates

FROM alpine:latest
COPY --from=vendor /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY bin/powerutil /usr/local/bin/powerutil
ENTRYPOINT "/usr/local/bin/powerutil"
