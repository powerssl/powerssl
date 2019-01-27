.PHONY: quickstart
quickstart: ca.pem
	$(MAKE) gencert HOSTNAME=localhost
	$(MAKE) gencert HOSTNAME=vault

ca.pem:
	cfssl genkey -initca csr.json | cfssljson --bare ca

.PHONY: initca
initca: ca.pem

.PHONY: gencert
gencert: ca.pem
	cfssl gencert -ca ca.pem -ca-key ca-key.pem -hostname=$$HOSTNAME csr.json | cfssljson -bare $$HOSTNAME

.PHONY: cfssl
cfssl:
	@GO111MODULE=off go get -u github.com/cloudflare/cfssl/cmd/cfssl

.PHONY: cfssljson
cfssljson:
	@GO111MODULE=off go get -u github.com/cloudflare/cfssl/cmd/cfssljson
