# PowerSSL

-	Website: https://www.powerssl.io

<img width="300" alt="PowerSSL Logo" src="https://docs.powerssl.io/assets/images/powerssl.png">

PowerSSL is a ....

Getting Started & Documentation
-------------------------------

All documentation is available on the [PowerSSL documentation](https://docs.powerssl.io).

Developing PowerSSL
-------------------

If you wish to work on Vault itself or any of its built-in systems, you'll
first need [Go](https://www.golang.org) installed on your machine (version
1.11+ is *required*).

### Prepare for local development

```sh
$ mkdir /etc/powerssl
$ cd /etc/powerssl
$ powerutil ca init
$ mkdir vault
$ cd vault
$ powerutil ca gen --ca /etc/powerssl/ca.pem --ca-key /etc/powerssl/ca-key.pem --hostname localhost
```

### Run locally with docker compose

```sh
$ docker-compose -f deployments/docker-compose.yml --project-directory . up
```

```sh
$ deployments/init-vault.sh
```

### Run locally without docker

```sh
$ forego start -f deployments/Procfile
```