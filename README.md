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
1.12+ is *required*).

You can then download any required build tools by bootstrapping your environment:

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
make prepare-local-dev
...
```

### Run locally without docker

```sh
$ make run 
...
```

Visit [localhost:8080](http://localhost:8080) in your browser.
Perform a login on the command line.
Afterwards execute `powerctl login --ca-file $(pwd)/local/certs/ca.pem` from the workspace dir.

### Run locally with docker compose

```sh
$ docker-compose -d -f deployments/docker-compose.yml --project-directory . up
...
```

Visit [localhost:8080](http://localhost:8080) in your browser.
