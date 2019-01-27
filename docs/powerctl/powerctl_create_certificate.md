---
grand_parent: powerctl
layout: default
parent: create
title: certificate
---
## powerctl create certificate

Create Certificate

### Synopsis

Create Certificate

```
powerctl create certificate [flags]
```

### Options

```
      --auto-renew                Auto renew ...
      --digest-algorithm string   Digest algorithm ...
      --dns-names string          DNS name for the certificate (seperated by ",")
  -h, --help                      help for certificate
      --key-algorithm string      Key algorithm ...
      --key-size int              Key size ...
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

* [powerctl create](powerctl_create.md)	 - Create resource

###### Auto generated by spf13/cobra on 17-Dec-2018