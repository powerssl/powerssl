---
has_toc: false
layout: default
parent: powerctl
permalink: /powerctl/login
title: login
---
## powerctl login

Login to PowerSSL

```
powerctl login [flags]
```

### Options

```
  -h, --help   help for login
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
