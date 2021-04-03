# PowerSSL [![CircleCI](https://circleci.com/gh/powerssl/powerssl.svg?style=svg&circle-token=572c8a8bf77274579537593224433d5de2a0bf09)](https://circleci.com/gh/powerssl/powerssl)

- Website: https://www.powerssl.io

![PowerSSL Logo](https://docs.powerssl.io/assets/images/powerssl.png)

PowerSSL is a certificate management platform.

Getting Started & Documentation
-------------------------------

All documentation is available on the [PowerSSL documentation](https://docs.powerssl.io).

Developing PowerSSL
-------------------

If you wish to work on PowerSSL itself or any of its built-in systems, you'll
first need [Go](https://www.golang.org) installed on your machine (version
1.16+ is *required*). Ensure *gcc-5* package is installed as its a required build dependency.

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

To install all user tools
```sh
$ make install
...
```

To build a dockerized version of PowerSSL, run `make images`.

```sh
$ make images
...
```

### Run locally

```sh
$ make run
...
```

Visit [localhost:8080](http://localhost:8080) in your browser.
Perform a login on the command line.
Afterwards execute `powerctl login --ca-file $(pwd)/local/certs/ca.pem` from the workspace dir.
