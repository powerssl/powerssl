---
has_toc: false
layout: default
parent: powerssl-grpcgateway
permalink: /powerssl-grpcgateway/serve
title: serve
---
## powerssl-grpcgateway serve

Serve the gRPC Gateway

### Synopsis

Serve the gRPC Gateway

```
powerssl-grpcgateway serve [flags]
```

### Options

```
      --addr string                             Addr (default ":8080")
      --apiserver-addr string                   GRPC address of APIServer
      --apiserver-auth-token string             APIServer authentication token
      --apiserver-insecure                      Use insecure communication
      --apiserver-insecure-skip-tls-verify      Accepts any certificate presented by the server and any host name in that certificate
      --apiserver-server-name-override string   It will override the virtual host name of authority
      --ca-file string                          Certificate authority file
  -h, --help                                    help for serve
      --metrics-addr string                     HTTP Addr (default ":9090")
      --no-metrics                              Do not serve metrics
```

### Options inherited from parent commands

```
      --config string   config file (default is /etc/powerssl/grpcgateway/config.yaml)
  -v, --verbose         Verbose output
```

### SEE ALSO

* [powerssl-grpcgateway](/powerssl-grpcgateway)	 - powerssl-grpcgateway provides PowerSSL gRPC Gateway
