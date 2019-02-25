---
has_children: true
has_toc: false
layout: default
parent: powerctl
permalink: /powerctl/create
title: create
---
## powerctl create

Create resource

### Synopsis

Create resource

```
powerctl create [flags]
```

### Examples

```
  # Create a certificate using the data in certificate.json.
  powerctl create -f ./certificate.json

  # Create a certificate based on the JSON passed into stdin.
  cat certificate.json | powerctl create -f -
```

### Options

```
  -f, --filename string   Filename to file to use to create the resources
  -h, --help              help for create
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

* [powerctl](/powerctl)	 - powerctl controls PowerSSL
* [powerctl create acmeaccount](/powerctl/create/acmeaccount)	 - Create ACME account
* [powerctl create acmeserver](/powerctl/create/acmeserver)	 - Create ACME server
* [powerctl create certificate](/powerctl/create/certificate)	 - Create Certificate
* [powerctl create user](/powerctl/create/user)	 - Create ACME server
