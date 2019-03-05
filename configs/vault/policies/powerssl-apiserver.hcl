path "pki/sign/powerssl-apiserver" {
  capabilities = ["update"]
}

path "transit/keys/*" {
  capabilities = ["update"]
}
