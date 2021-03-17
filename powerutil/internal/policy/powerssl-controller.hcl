path "pki/sign/powerssl-controller" {
  capabilities = ["update"]
}

path "transit/export/signing-key/*" {
  capabilities = ["read"]
}
