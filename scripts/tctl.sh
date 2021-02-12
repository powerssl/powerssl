#!/bin/bash

set -euxo pipefail

tctl --address localhost:7233 --namespace powerssl --tls_ca_path local/certs/ca.pem --tls_server_name localhost
