---
has_toc: false
layout: default
parent: powerssl-temporalserver
permalink: /powerssl-temporalserver/run
title: run
---
## powerssl-temporalserver run

Run the Temporal Server

```
powerssl-temporalserver run [flags]
```

### Options

```
      --config-dir string     Config directory to load a set of yaml config files from (default "config")
      --env string            Environment is one of the input params ex-development (default "development")
  -h, --help                  help for run
      --service stringArray   Service(s) to start (default [frontend,history,matching,worker])
      --zone string           Zone is another input param
```

### Options inherited from parent commands

```
      --config string   config file (default is /etc/powerssl/temporalserver/config.yaml)
  -v, --verbose         Verbose output
```

### SEE ALSO

* [powerssl-temporalserver](/powerssl-temporalserver)	 - powerssl-temporalserver provides PowerSSL Temporal Server
