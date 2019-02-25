# PowerSSL

-	Website: https://www.powerssl.io

<img width="300" alt="PowerSSL Logo" src="https://docs.powerssl.io/assets/images/powerssl.png">

PowerSSL is a certificate management platform... more description goes here.

Getting Started & Documentation
-------------------------------

All documentation is available on the [PowerSSL documentation](https://docs.powerssl.io).

Developing PowerSSL
-------------------

If you wish to work on Vault itself or any of its built-in systems, you'll
first need [Go](https://www.golang.org) installed on your machine (version
1.11+ is *required*).

```sh
$ make bootstrap
...
```

To compile a version of PowerSSL, run `make` or `make build`.
This will put the PowerSSL binaries in the `bin` folder:

```sh
$ make build
...
```

To build a dockerized version of PowerSSL, run `make images`.

```sh
$ make images
...
```

### Prepare for running locally

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
$ docker-compose -d -f deployments/docker-compose.yml --project-directory . up
...
$ deployments/init-vault.sh
...
```

### Run locally without docker

```sh
$ forego start -f deployments/Procfile
```