---
has_toc: false
layout: default
parent: powerssl-grpcgateway
permalink: /powerssl-grpcgateway/serve
title: serve
---
## powerssl-grpcgateway serve

Serve the gRPC Gateway

```
powerssl-grpcgateway serve [flags]
```

### Options

```
      --api-server-client-addr string                   client addr
      --api-server-client-ca-file string                client CA file
      --api-server-client-insecure                      client insecure
      --api-server-client-insecure-skip-tls-verify      client insecure skip TLS verify
      --api-server-client-server-name-override string   client server name override
  -h, --help                                            help for serve
  -e, --log-env string                                  environment (default "production")
      --server-addr string                              server addr
      --telemetry-meter-addr string                     metrics addr
      --telemetry-meter-exporter string                 metrics exporter (default "prometheus")
      --telemetry-tracer-disabled                       disable tracer
```

### SEE ALSO

* [powerssl-grpcgateway](/powerssl-grpcgateway)	 - powerssl-grpcgateway provides PowerSSL gRPC Gateway
