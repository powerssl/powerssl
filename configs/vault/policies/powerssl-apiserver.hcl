path "pki/issue/powerssl-apiserver" {
  capabilities = ["update"]
}

path "transit/keys/*" {
  capabilities = ["update"]
}
