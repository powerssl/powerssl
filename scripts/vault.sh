#!/bin/bash

set -euo pipefail

export VAULT_ADDR=https://localhost:8200
export VAULT_CACERT=local/certs/ca.pem
export VAULT_TOKEN=${VAULT_TOKEN:-"$(yq e '.rootToken' local/vault/secret.yaml)"}

vault "$@"
