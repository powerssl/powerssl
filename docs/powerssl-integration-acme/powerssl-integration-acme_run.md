---
has_toc: false
layout: default
parent: powerssl-integration-acme
permalink: /powerssl-integration-acme/run
title: run
---
## powerssl-integration-acme run

Run ACME integration

```
powerssl-integration-acme run [flags]
```

### Options

```
      --auth-token string                        Authentication token
      --ca-file string                           Certificate authority file
      --controller-addr string                   GRPC address of Controller
      --controller-insecure                      Use insecure communication
      --controller-insecure-skip-tls-verify      Accepts any certificate presented by the server and any host name in that certificate
      --controller-server-name-override string   It will override the virtual host name of authority
  -h, --help                                     help for run
      --metrics-addr string                      HTTP Addr (default ":9090")
      --no-metrics                               Do not serve metrics
      --no-tracing                               Do not trace
      --tracer string                            Tracing implementation (default "jaeger")
```

### Options inherited from parent commands

```
      --config string   config file (default is /etc/powerssl/integration-acme/config.yaml)
  -v, --verbose         Verbose output
```

### SEE ALSO

* [powerssl-integration-acme](/powerssl-integration-acme)	 - powerssl-integration-acme provides PowerSSL ACME integration
