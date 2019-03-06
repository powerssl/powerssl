.PHONY: quickstart
quickstart: ca.pem
	$(MAKE) gencert HOSTNAME=localhost
	$(MAKE) gencert HOSTNAME=vault

ca.pem:
	powerutil ca init

.PHONY: initca
initca: ca.pem

.PHONY: gencert
gencert: ca.pem
	powerutil ca gen --ca ca.pem --ca-key ca-key.pem --hostname $$HOSTNAME
