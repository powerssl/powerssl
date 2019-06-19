---
has_toc: false
layout: default
parent: powerssl-apiserver
permalink: /powerssl-apiserver/serve
title: serve
---
## powerssl-apiserver serve

Serve the API

### Synopsis

Serve the API

```
powerssl-apiserver serve [flags]
```

### Options

```
      --addr string                              GRPC Addr (default ":8080")
      --auth-token string                        Authentication token
      --ca-file string                           Certificate authority file
      --common-name string                       API Server common name
      --controller-addr string                   GRPC address of Controller
      --controller-insecure                      Use insecure communication
      --controller-insecure-skip-tls-verify      Accepts any certificate presented by the server and any host name in that certificate
      --controller-server-name-override string   It will override the virtual host name of authority
      --db-connection string                     DB connection (default "/tmp/powerssl.sqlie3")
      --db-dialect string                        DB Dialect (default "sqlite3")
  -h, --help                                     help for serve
      --insecure                                 Do not use TLS for the server
      --jwks-url string                          JWKS URL
      --metrics-addr string                      HTTP Addr (default ":9090")
      --no-metrics                               Do not serve metrics
      --no-tracing                               Do not trace
      --tls-cert-file string                     File containing the default x509 Certificate for GRPC
      --tls-private-key-file string              File containing the default x509 private key matching --tls-cert-file
      --tracer string                            Tracing implementation (default "jaeger")
      --vault-token string                       Vault Token
      --vault-url string                         Vault URL
```

### Options inherited from parent commands

```
      --config string   config file (default is /etc/powerssl/api/config.yaml)
  -v, --verbose         Verbose output
```

### SEE ALSO

* [powerssl-apiserver](/powerssl-apiserver)	 - powerssl-apiserver provides PowerSSL API
