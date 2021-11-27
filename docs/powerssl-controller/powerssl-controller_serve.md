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
      --api-server-client-auth-token string                    apiserver client addr
      --api-server-client-client-addr string                   client addr
      --api-server-client-client-ca-file string                client CA file
      --api-server-client-client-insecure                      client insecure
      --api-server-client-client-insecure-skip-tls-verify      client insecure skip TLS verify
      --api-server-client-client-server-name-override string   client server name override
  -h, --help                                                   help for serve
  -e, --log-env string                                         environment (default "production")
      --server-addr string                                     server addr
      --server-cert-file string                                server Cert file
      --server-common-name string                              server common name
      --server-insecure                                        server insecure
      --server-key-file string                                 server key file
      --server-vault-app-role-id string                        vault app role ID
      --server-vault-app-role-secret-id string                 vault app role secret ID
      --server-vault-ca-file string                            vault CA file
      --server-vault-token string                              vault token
      --server-vault-url string                                vault URL
      --telemetry-meter-addr string                            metrics addr
      --telemetry-meter-exporter string                        metrics exporter (default "prometheus")
      --telemetry-tracer-disabled                              disable tracer
      --temporal-client-ca-file string                         temporal CA file
      --temporal-client-disable-health-check                   temporal disable health check
      --temporal-client-health-check-timeout int               temporal health check timeout
      --temporal-client-host-port string                       temporal host port
      --temporal-client-identity string                        temporal identity
      --temporal-client-namespace string                       temporal namespace
      --temporal-client-tls-cert-file string                   temporal TLS cert file
      --temporal-client-tls-key-file string                    temporal TLS key file
      --vault-client-app-role-id string                        vault app role ID
      --vault-client-app-role-secret-id string                 vault app role secret ID
      --vault-client-ca-file string                            vault CA file
      --vault-client-token string                              vault token
      --vault-client-url string                                vault URL
```

### SEE ALSO

* [powerssl-controller](/powerssl-controller)	 - powerssl-controller provides PowerSSL Controller
