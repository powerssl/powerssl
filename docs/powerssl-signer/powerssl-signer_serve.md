---
has_toc: false
layout: default
parent: powerssl-signer
permalink: /powerssl-signer/serve
title: serve
---
## powerssl-signer serve

Serve the Signer

```
powerssl-signer serve [flags]
```

### Options

```
      --addr string                   GRPC Addr (default ":8080")
      --ca-file string                Certificate authority file
      --common-name string            API Server common name
  -h, --help                          help for serve
      --insecure                      Do not use TLS for the server
      --metrics-addr string           HTTP Addr (default ":9090")
      --no-metrics                    Do not serve metrics
      --no-tracing                    Do not trace
      --tls-cert-file string          File containing the default x509 Certificate for GRPC.
      --tls-private-key-file string   File containing the default x509 private key matching --tls-cert-file.
      --tracer string                 Tracing implementation (default "jaeger")
      --vault-token string            Vault Token
      --vault-url string              Vault URL
```

### Options inherited from parent commands

```
      --config string   config file (default is /etc/powerssl/signer/config.yaml)
  -v, --verbose         Verbose output
```

### SEE ALSO

* [powerssl-signer](/powerssl-signer)	 - powerssl-signer provides PowerSSL Signer
