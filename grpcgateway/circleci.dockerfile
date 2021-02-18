FROM alpine:latest AS vendor
RUN apk add --no-cache ca-certificates

FROM alpine:latest
COPY --from=vendor /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY bin/powerssl-grpcgateway /usr/local/bin/powerssl-grpcgateway
ENTRYPOINT "/usr/local/bin/powerssl-grpcgateway"
CMD ["serve"]
EXPOSE 8080/tcp 9090/tcp
