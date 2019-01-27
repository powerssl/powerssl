#!/bin/sh

export CERTS_DIR=${CERTS_DIR:-"examples/certs"}
export VAULT_ADDR=${VAULT_ADDR:-"https://127.0.0.1:8200/"}
export VAULT_CAPATH=${VAULT_CAPATH:-"/etc/powerssl/ca.pem"}
export VAULT_TLS_SERVER_NAME=${VAULT_TLS_SERVER_NAME:-"vault"}
# export VAULT_TOKEN=${VAULT_TOKEN:-"insecure"}

initoutput=$(vault operator init -key-shares=1 -key-threshold=1)
unsealkey=$(echo $initoutput | cut -d ' ' -f4)
roottoken=$(echo $initoutput | cut -d ' ' -f8)
echo "key $unsealkey"
echo "token $roottoken"

export VAULT_TOKEN=$roottoken

vault operator unseal $unsealkey

# Initialize secret engines
vault secrets enable pki
vault secrets enable transit

# Initialize pki engine
vault secrets tune -max-lease-ttl=8760h pki
csr=$(vault write -field=csr pki/intermediate/generate/internal \
	common_name="PowerSSL Intermediate Authority" \
	ttl=8760h)
cert=$(echo "$csr" | cfssl sign -profile intermediate -config $CERTS_DIR/config.json -ca $CERTS_DIR/ca.pem -ca-key $CERTS_DIR/ca-key.pem -csr - | jq -r '.cert')
vault write pki/intermediate/set-signed certificate="$cert"
vault write pki/config/urls \
	crl_distribution_points="${VAULT_ADDR}v1/pki/crl" \
	issuing_certificates="${VAULT_ADDR}v1/pki/ca"

# Initialize pki roles
vault write pki/roles/powerssl-apiserver \
	allowed_domains=apiserver \
	allow_localhost=false \
	allow_bare_domains=true \
	max_ttl=24h
vault write pki/roles/powerssl-controller \
	allowed_domains=controller \
	allow_localhost=false \
	allow_bare_domains=true \
	max_ttl=24h
vault write pki/roles/powerssl-signer \
	allowed_domains=signer \
	allow_localhost=false \
	allow_bare_domains=true \
	max_ttl=24h

# Initialize policies
vault policy write powerssl-apiserver ./examples/vault/policies/powerssl-apiserver.hcl
vault policy write powerssl-controller ./examples/vault/policies/powerssl-controller.hcl
vault policy write powerssl-signer ./examples/vault/policies/powerssl-signer.hcl

# Initialize tokens
vault token create -id=powerssl-apiserver -policy=powerssl-apiserver
vault token create -id=powerssl-controller -policy=powerssl-controller
vault token create -id=powerssl-signer -policy=powerssl-signer
