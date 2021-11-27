---
has_toc: false
layout: default
parent: powerssl-agent
permalink: /powerssl-agent/run
title: run
---
## powerssl-agent run

Run the Agent

```
powerssl-agent run [flags]
```

### Options

```
      --api-server-client-auth-token string                    apiserver client addr
      --api-server-client-client-addr string                   client addr
      --api-server-client-client-ca-file string                client CA file
      --api-server-client-client-insecure                      client insecure
      --api-server-client-client-insecure-skip-tls-verify      client insecure skip TLS verify
      --api-server-client-client-server-name-override string   client server name override
  -h, --help                                                   help for run
  -e, --log-env string                                         environment (default "production")
      --telemetry-meter-addr string                            metrics addr
      --telemetry-meter-exporter string                        metrics exporter (default "prometheus")
      --telemetry-tracer-disabled                              disable tracer
```

### SEE ALSO

* [powerssl-agent](/powerssl-agent)	 - powerssl-agent provides PowerSSL Agent
