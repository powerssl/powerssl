---
has_toc: false
layout: default
parent: powerssl-worker
permalink: /powerssl-worker/run
title: run
---
## powerssl-worker run

Run the Worker

```
powerssl-worker run [flags]
```

### Options

```
      --apiserver-addr string                   GRPC address of API server
      --apiserver-insecure                      Use insecure communication
      --apiserver-insecure-skip-tls-verify      Accepts any certificate presented by the server and any host name in that certificate
      --apiserver-server-name-override string   It will override the virtual host name of authority
      --auth-token string                       Authentication token
      --ca-file string                          Certificate authority file
  -h, --help                                    help for run
      --metrics-addr string                     HTTP Addr (default ":9090")
      --no-metrics                              Do not serve metrics
      --no-tracing                              Do not trace
      --temporal-host-port string               Host and port for this client to connect to (default "localhost:7233")
      --temporal-namespace string               Namespace name for this client to work with (default "powerssl")
      --tracer string                           Tracing implementation (default "jaeger")
      --vault-token string                      Vault Token
      --vault-url string                        Vault URL
```

### Options inherited from parent commands

```
      --config string   config file (default is /etc/powerssl/worker/config.yaml)
  -v, --verbose         Verbose output
```

### SEE ALSO

* [powerssl-worker](/powerssl-worker)	 - powerssl-worker provides PowerSSL Worker
