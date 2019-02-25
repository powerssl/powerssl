---
grand_parent: powerctl
has_toc: false
layout: default
parent: create
permalink: /powerctl/create/acmeserver
title: acmeserver
---
## powerctl create acmeserver

Create ACME server

### Synopsis

Create ACME server

```
powerctl create acmeserver [flags]
```

### Options

```
      --directory-url string      Directory URL
      --display-name string       Display name
  -h, --help                      help for acmeserver
      --integration-name string   Integration name
      --letsencrypt               Let's Encrypt defaults
      --letsencrypt-staging       Let's Encrypt staging defaults
```

### Options inherited from parent commands

```
      --addr string                   GRPC address of API server
      --auth-token string             Authentication token
      --ca-file string                Certificate authority file
      --config string                 config file (default is $HOME/.powerctl/config.yaml)
      --insecure                      Use insecure communication
      --insecure-skip-tls-verify      Accepts any certificate presented by the server and any host name in that certificate
  -o, --output string                 Output format (default "table")
      --server-name-override string   It will override the virtual host name of authority
  -v, --verbose                       Verbose output
```

### SEE ALSO

* [powerctl create](/powerctl/create)	 - Create resource
