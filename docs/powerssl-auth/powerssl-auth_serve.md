---
has_toc: false
layout: default
parent: powerssl-auth
permalink: /powerssl-auth/serve
title: serve
---
## powerssl-auth serve

Serve Web

```
powerssl-auth serve [flags]
```

### Options

```
  -h, --help                                 help for serve
  -e, --log-env string                       environment (default "production")
      --oauth2-auth-uri string               oAuth2 auth URI
      --oauth2-github-client-id string       oAuth2 GitHub client ID
      --oauth2-github-client-secret string   oAuth2 GitHub client secret
      --server-addr string                   server addr
      --server-cert-file string              server cert file
      --server-insecure                      server insecure
      --server-jwt-private-key-file string   server JWT private key file
      --server-key-file string               server key file
      --server-webapp-uri string             webapp URI
      --telemetry-meter-addr string          metrics addr
      --telemetry-meter-exporter string      metrics exporter (default "prometheus")
      --telemetry-tracer-disabled            disable tracer
```

### SEE ALSO

* [powerssl-auth](/powerssl-auth)	 - powerssl-auth provides PowerSSL Auth
