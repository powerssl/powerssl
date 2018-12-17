#!/bin/bash

export CERTS_DIR=${CERTS_DIR:-"examples/certs"}
export VAULT_ADDR=${VAULT_ADDR:-"http://127.0.0.1:8200/"}
export VAULT_TOKEN=${VAULT_TOKEN:-"insecure"}

trap "exit" INT TERM ERR
trap "kill 0" EXIT

vault server -dev -dev-root-token-id=$VAULT_TOKEN &

sleep 1

vault secrets enable pki
vault secrets tune -max-lease-ttl=8760h pki
csr=$(vault write -field=csr pki/intermediate/generate/internal \
	common_name="PowerSSL Intermediate Authority" \
	ttl=8760h)
cert=$(echo "$csr" | cfssl sign -profile intermediate -config $CERTS_DIR/config.json -ca $CERTS_DIR/ca.pem -ca-key $CERTS_DIR/ca-key.pem -csr - | jq -r '.cert')
vault write pki/intermediate/set-signed certificate="$cert"
vault write pki/config/urls \
	crl_distribution_points="${VAULT_ADDR}v1/pki/crl" \
	issuing_certificates="${VAULT_ADDR}v1/pki/ca"
vault write pki/roles/powerssl-apiserver \
	max_ttl=24h
vault write pki/roles/powerssl-controller \
	max_ttl=24h

wait
