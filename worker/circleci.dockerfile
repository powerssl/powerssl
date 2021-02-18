FROM alpine:latest AS vendor
RUN apk add --no-cache ca-certificates

FROM alpine:latest
COPY --from=vendor /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY bin/powerssl-worker /usr/local/bin/powerssl-worker
ENTRYPOINT "/usr/local/bin/powerssl-worker"
CMD ["serve"]
EXPOSE 9090/tcp
