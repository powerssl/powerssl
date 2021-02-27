---
has_children: true
has_toc: false
layout: default
parent: powerssl-temporal
permalink: /powerssl-temporal/migrate
title: migrate
---
## powerssl-temporal migrate

Run temporal migrations

```
powerssl-temporal migrate [flags]
```

### Options

```
  -h, --help                         help for migrate
      --host string                  DB host
      --password string              DB Password
      --plugin string                DB Plugin
      --port string                  DB Port
      --temporal-database string     Temporal DB (default "temporal")
      --user string                  DB User
      --visibility-database string   Visibility DB (default "temporal_visibility")
```

### Options inherited from parent commands

```
      --config string   config file (default is /etc/powerssl/temporal/config.yaml)
  -v, --verbose         Verbose output
```

### SEE ALSO

* [powerssl-temporal](/powerssl-temporal)	 - powerssl-temporal provides PowerSSL Temporal Server
