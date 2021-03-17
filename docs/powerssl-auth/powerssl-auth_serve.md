---
has_toc: false
layout: default
parent: powerssl-auth
permalink: /powerssl-auth/serve
title: serve
---
## powerssl-auth serve

Serve the API

```
powerssl-auth serve [flags]
```

### Options

```
      --addr string                   GRPC (default ":8080")
  -h, --help                          help for serve
      --insecure                      Do not use TLS for the server
      --jwt-private-key-file string   JWT private key file
      --metrics-addr string           HTTP Addr (default ":9090")
      --no-metrics                    Do not serve metrics
      --webapp-uri string             WebApp URI
```

### Options inherited from parent commands

```
      --config string   config file (default is /etc/powerssl/auth/config.yaml)
  -v, --verbose         Verbose output
```

### SEE ALSO

* [powerssl-auth](/powerssl-auth)	 - powerssl-auth provides PowerSSL Auth
