.PHONY: quickstart
quickstart: ca.pem
	$(MAKE) gencert HOSTNAME=localhost
	$(MAKE) gencert HOSTNAME=vault

ca.pem: ../../bin/powerutil
	../../bin/powerutil ca init

.PHONY: initca
initca: ca.pem

.PHONY: gencert
gencert: ../../bin/powerutil ca.pem
	../../bin/powerutil ca gen --ca ca.pem --ca-key ca-key.pem --hostname $$HOSTNAME

../../bin/powerutil:
	$(MAKE) -C ../.. bin/powerutil
