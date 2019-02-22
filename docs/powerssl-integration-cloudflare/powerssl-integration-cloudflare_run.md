---
layout: default
parent: powerssl-integration-cloudflare
title: run
---
## powerssl-integration-cloudflare run

Run ACME integration

### Synopsis

Run ACME integration

```
powerssl-integration-cloudflare run [flags]
```

### Options

```
      --addr string                   GRPC address of Controller
      --ca-file string                Certificate authority file
  -h, --help                          help for run
      --insecure                      Use insecure communication
      --insecure-skip-tls-verify      Accepts any certificate presented by the server and any host name in that certificate
      --metrics-addr string           HTTP Addr (default ":9090")
      --no-metrics                    Do not serve metrics
      --no-tracing                    Do not trace
      --server-name-override string   It will override the virtual host name of authority
      --tracer string                 Tracing implementation (default "jaeger")
```

### Options inherited from parent commands

```
      --config string   config file (default is /etc/powerssl/integration-cloudflare/config.yaml)
  -v, --verbose         Verbose output
```

### SEE ALSO

* [powerssl-integration-cloudflare](powerssl-integration-cloudflare.md)	 - powerssl-integration-cloudflare provides PowerSSL Cloudflare integration
