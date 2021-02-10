---
has_toc: false
layout: default
parent: powerssl-agent
permalink: /powerssl-agent/run
title: run
---
## powerssl-agent run

Run the Agent

```
powerssl-agent run [flags]
```

### Options

```
      --apiserver-addr string                   GRPC address of API server
      --apiserver-insecure                      Use insecure communication
      --apiserver-insecure-skip-tls-verify      Accepts any certificate presented by the server and any host name in that certificate
      --apiserver-server-name-override string   It will override the virtual host name of authority
      --auth-token string                       Auth token
      --ca-file string                          Certificate authority file
  -h, --help                                    help for run
```

### Options inherited from parent commands

```
      --config string   config file (default is /etc/powerssl/agent/config.yaml)
  -v, --verbose         Verbose output
```

### SEE ALSO

* [powerssl-agent](/powerssl-agent)	 - powerssl-agent provides PowerSSL Agent
