---
has_toc: false
layout: default
parent: powerssl-controller
permalink: /powerssl-controller/serve
title: serve
---
## powerssl-controller serve

Serve the Controller

```
powerssl-controller serve [flags]
```

### Options

```
      --addr string                             GRPC Addr (default ":8080")
      --apiserver-addr string                   GRPC address of API server
      --apiserver-insecure                      Use insecure communication
      --apiserver-insecure-skip-tls-verify      Accepts any certificate presented by the server and any host name in that certificate
      --apiserver-server-name-override string   It will override the virtual host name of authority
      --auth-token string                       Authentication token
      --ca-file string                          Certificate authority file
      --common-name string                      API Server common name
  -h, --help                                    help for serve
      --insecure                                Do not use TLS for the server
      --metrics-addr string                     HTTP Addr (default ":9090")
      --no-metrics                              Do not serve metrics
      --no-tracing                              Do not trace
      --temporal-host-port string               Host and port for this client to connect to (default "localhost:7233")
      --temporal-namespace string               Namespace name for this client to work with (default "powerssl")
      --tls-cert-file string                    File containing the default x509 Certificate for GRPC.
      --tls-private-key-file string             File containing the default x509 private key matching --tls-cert-file.
      --tracer string                           Tracing implementation (default "jaeger")
      --vault-token string                      Vault Token
      --vault-url string                        Vault URL
```

### Options inherited from parent commands

```
      --config string   config file (default is /etc/powerssl/controller/config.yaml)
  -v, --verbose         Verbose output
```

### SEE ALSO

* [powerssl-controller](/powerssl-controller)	 - powerssl-controller provides PowerSSL Controller
