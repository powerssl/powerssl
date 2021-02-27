---
has_children: true
has_toc: false
layout: default
parent: powerssl-apiserver
permalink: /powerssl-apiserver/migrate
title: migrate
---
## powerssl-apiserver migrate

Migrate

### Options

```
      --database-url string   Database URL
  -h, --help                  help for migrate
```

### Options inherited from parent commands

```
      --config string   config file (default is /etc/powerssl/api/config.yaml)
  -v, --verbose         Verbose output
```

### SEE ALSO

* [powerssl-apiserver](/powerssl-apiserver)	 - powerssl-apiserver provides PowerSSL API
* [powerssl-apiserver migrate down](/powerssl-apiserver/migrate/down)	 - Apply all or N down migrations
* [powerssl-apiserver migrate drop](/powerssl-apiserver/migrate/drop)	 - Drop everything inside database
* [powerssl-apiserver migrate force](/powerssl-apiserver/migrate/force)	 - Set version V but don't run migration (ignores dirty state)
* [powerssl-apiserver migrate goto](/powerssl-apiserver/migrate/goto)	 - Migrate to version V
* [powerssl-apiserver migrate up](/powerssl-apiserver/migrate/up)	 - Apply all or N up migrations
