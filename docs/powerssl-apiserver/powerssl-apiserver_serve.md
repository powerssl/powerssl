---
has_toc: false
layout: default
parent: powerssl-apiserver
permalink: /powerssl-apiserver/serve
title: serve
---
## powerssl-apiserver serve

Serve the API

```
powerssl-apiserver serve [flags]
```

### Options

```
      --addr string                   GRPC Addr (default ":8080")
      --ca-file string                Certificate authority file
      --common-name string            API Server common name
      --db-connection string          DB connection (default "/tmp/powerssl.sqlie3")
      --db-dialect string             DB Dialect (default "sqlite3")
  -h, --help                          help for serve
      --insecure                      Do not use TLS for the server
      --jwks-url string               JWKS URL
      --metrics-addr string           HTTP Addr (default ":9090")
      --no-metrics                    Do not serve metrics
      --no-tracing                    Do not trace
      --temporal-host-port string     Host and port for this client to connect to (default "localhost:7233")
      --temporal-namespace string     Namespace name for this client to work with (default "powerssl")
      --tls-cert-file string          File containing the default x509 Certificate for GRPC
      --tls-private-key-file string   File containing the default x509 private key matching --tls-cert-file
      --tracer string                 Tracing implementation (default "jaeger")
      --vault-token string            Vault Token
      --vault-url string              Vault URL
```

### Options inherited from parent commands

```
      --config string   config file (default is /etc/powerssl/api/config.yaml)
  -v, --verbose         Verbose output
```

### SEE ALSO

* [powerssl-apiserver](/powerssl-apiserver)	 - powerssl-apiserver provides PowerSSL API
