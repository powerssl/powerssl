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
  -h, --help                       help for run
  -e, --log-env string             environment (default "production")
      --server-config-dir string   Config directory to load a set of yaml config files from (default "config")
      --server-env string          Environment is one of the input params ex-development (default "development")
      --server-services strings    Service(s) to start (default [frontend,history,matching,worker])
      --server-zone string         Zone is another input param
```

### SEE ALSO

* [powerssl-temporal](/powerssl-temporal)	 - powerssl-temporal provides PowerSSL Temporal Server
