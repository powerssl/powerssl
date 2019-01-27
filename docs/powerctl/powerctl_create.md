---
has_children: true
layout: default
parent: powerctl
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

* [powerctl](powerctl.md)	 - powerctl controls PowerSSL
* [powerctl create acmeaccount](powerctl_create_acmeaccount.md)	 - Create ACME account
* [powerctl create acmeserver](powerctl_create_acmeserver.md)	 - Create ACME server
* [powerctl create certificate](powerctl_create_certificate.md)	 - Create Certificate
* [powerctl create user](powerctl_create_user.md)	 - Create ACME server

###### Auto generated by spf13/cobra on 23-Jan-2019