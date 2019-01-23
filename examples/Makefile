.PHONY: quickstart
quickstart:
	mkdir ~/.powerctl
	cp powerctl/config.yaml ~/.powerctl/config.yaml
	mkdir -p /etc/powerssl
	cp -r api /etc/powerssl
	cp -r auth /etc/powerssl
	cp -r controller /etc/powerssl
	cp -r signer /etc/powerssl
	cp -r webapp /etc/powerssl
	make -C certs
	cp certs/ca.pem /etc/powerssl
	cp certs/localhost.pem /etc/powerssl/api/cert.pem
	cp certs/localhost-key.pem /etc/powerssl/api/cert-key.pem
	chmod +r /etc/powerssl/api/cert-key.pem
	cp certs/localhost.pem /etc/powerssl/auth/cert.pem
	cp certs/localhost-key.pem /etc/powerssl/auth/cert-key.pem
	chmod +r /etc/powerssl/auth/cert-key.pem
	cp certs/localhost.pem /etc/powerssl/controller/cert.pem
	cp certs/localhost-key.pem /etc/powerssl/controller/cert-key.pem
	chmod +r /etc/powerssl/controller/cert-key.pem
	cp certs/localhost.pem /etc/powerssl/signer/cert.pem
	cp certs/localhost-key.pem /etc/powerssl/signer/cert-key.pem
	chmod +r /etc/powerssl/signer/cert-key.pem
	cp certs/localhost.pem /etc/powerssl/webapp/cert.pem
	cp certs/localhost-key.pem /etc/powerssl/webapp/cert-key.pem
	chmod +r /etc/powerssl/webapp/cert-key.pem
