---
has_children: true
has_toc: false
layout: default
permalink: /powerctl
title: powerctl
---
## powerctl

powerctl controls PowerSSL

### Synopsis

powerctl controls PowerSSL.

### Options

```
      --addr string                   GRPC address of API server
      --auth-token string             Authentication token
      --ca-file string                Certificate authority file
      --config string                 config file (default is $HOME/.powerctl/config.yaml)
  -h, --help                          help for powerctl
      --insecure                      Use insecure communication
      --insecure-skip-tls-verify      Accepts any certificate presented by the server and any host name in that certificate
  -o, --output string                 Output format (default "table")
      --server-name-override string   It will override the virtual host name of authority
  -v, --verbose                       Verbose output
```

### SEE ALSO

* [powerctl create](/powerctl/create)	 - Create resource
* [powerctl delete](/powerctl/delete)	 - Delete resource
* [powerctl describe](/powerctl/describe)	 - Describe resource
* [powerctl get](/powerctl/get)	 - Get resource
* [powerctl login](/powerctl/login)	 - Login to PowerSSL
