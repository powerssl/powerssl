---
has_toc: false
layout: default
parent: powerssl-temporal
permalink: /powerssl-temporal/register-namespace
title: register-namespace
---
## powerssl-temporal register-namespace

Run temporal register namespace

```
powerssl-temporal register-namespace [flags]
```

### Options

```
      --address string                 host:port for Temporal frontend service (default "127.0.0.1:7233")
  -h, --help                           help for register-namespace
      --namespace string               Temporal workflow namespace (default "powerssl")
      --tls-ca-path string             path to server CA certificate
      --tls-cert-path string           path to x509 certificate
      --tls-enable-host-verification   validates hostname of temporal cluster against server certificate
      --tls-key-path string            path to private key
      --tls-server-name string         override for target server name
```

### Options inherited from parent commands

```
      --config string   config file (default is /etc/powerssl/temporal/config.yaml)
  -v, --verbose         Verbose output
```

### SEE ALSO

* [powerssl-temporal](/powerssl-temporal)	 - powerssl-temporal provides PowerSSL Temporal Server
