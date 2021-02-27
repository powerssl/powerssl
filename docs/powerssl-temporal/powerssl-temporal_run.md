---
has_toc: false
layout: default
parent: powerssl-temporal
permalink: /powerssl-temporal/run
title: run
---
## powerssl-temporal run

Run the Temporal Server

```
powerssl-temporal run [flags]
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
      --config string   config file (default is /etc/powerssl/temporal/config.yaml)
  -v, --verbose         Verbose output
```

### SEE ALSO

* [powerssl-temporal](/powerssl-temporal)	 - powerssl-temporal provides PowerSSL Temporal Server
