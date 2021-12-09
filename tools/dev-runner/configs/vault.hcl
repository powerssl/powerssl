api_addr          = "https://localhost:8200"
cluster_addr      = "https://localhost:8201"
default_lease_ttl = "168h"
disable_mlock     = "true"
max_lease_ttl     = "720h"
ui                = "true"

listener "tcp" {
  address       = "0.0.0.0:8200"
  tls_cert_file = "/etc/ssl/certs/localhost.pem"
  tls_key_file  = "/etc/ssl/private/localhost-key.pem"
}

storage "postgresql" {
  connection_url = "postgres://powerssl:powerssl@host.docker.internal:5432/vault?sslmode=disable"
}
