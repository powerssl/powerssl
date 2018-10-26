# Quickstart

Run make to generate all required certificates.

`make`

# Manual

Copy `powerctl.yaml` to `~/.powerctl.yaml`.

`cp powerctl.yaml ~/.powerctl.yaml`

Create `/etc/powerssl`.

`sudo mkdir -p /etc/powerssl`

Copy `api` to `/etc/powerssl/api`.

`sudo cp -r api /etc/powerssl`

Copy `controller` to `/etc/powerssl/controller`.

`sudo cp -r controller /etc/powerssl`

Copy `signer` to `/etc/powerssl/signer`.

`sudo cp -r signer /etc/powerssl`

Generate certificates:

`make -C certs`

Copy `certs/ca.pem` to `/etc/powerssl/ca.pem`

`sudo cp certs/ca.pem /etc/powerssl`

Copy `certs/localhost.pem` to `/etc/powerssl/api/cert.pem`.

`sudo cp certs/localhost.pem /etc/powerssl/api/cert.pem`

Copy `certs/localhost-key.pem` to `/etc/powerssl/api/cert-key.pem`.

`sudo cp certs/localhost-key.pem /etc/powerssl/api/cert-key.pem`

Grant read permissions on `/etc/powerssl/api/cert-key.pem`.

`sudo chmod +r /etc/powerssl/api/cert-key.pem`

Copy `certs/localhost.pem` to `/etc/powerssl/controller/cert.pem`.

`sudo cp certs/localhost.pem /etc/powerssl/controller/cert.pem`

Copy `certs/localhost-key.pem` to `/etc/powerssl/controller/cert-key.pem`.

`sudo cp certs/localhost-key.pem /etc/powerssl/controller/cert-key.pem`

Grant read permissions on `/etc/powerssl/controller/cert-key.pem`.

`sudo chmod +r /etc/powerssl/controller/cert-key.pem`

Copy `certs/localhost.pem` to `/etc/powerssl/signer/cert.pem`.

`sudo cp certs/localhost.pem /etc/powerssl/signer/cert.pem`

Copy `certs/localhost-key.pem` to `/etc/powerssl/signer/cert-key.pem`.

`sudo cp certs/localhost-key.pem /etc/powerssl/signer/cert-key.pem`

Grant read permissions on `/etc/powerssl/signer/cert-key.pem`.

`sudo chmod +r /etc/powerssl/signer/cert-key.pem`
