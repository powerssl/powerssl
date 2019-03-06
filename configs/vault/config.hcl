api_addr = "https://localhost:8200"
cluster_addr = "https://localhost:8201"
default_lease_ttl = "168h"
max_lease_ttl = "720h"
ui = "true"

listener "tcp" {
  address = "localhost:8200"
  tls_cert_file = "examples/certs/localhost.pem"
  tls_key_file = "examples/certs/localhost-key.pem"
}

storage "file" {
  path = "local/vault/file"
}
