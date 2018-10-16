.PHONY: quickstart
quickstart:
	cp powerctl.yaml ~/.powerctl.yaml
	sudo mkdir -p /etc/powerssl
	sudo cp -r api /etc/powerssl
	sudo cp -r controller /etc/powerssl
	make -C certs
	sudo cp certs/ca.pem /etc/powerssl
	sudo cp certs/localhost.pem /etc/powerssl/api/cert.pem
	sudo cp certs/localhost-key.pem /etc/powerssl/api/cert-key.pem
	sudo chmod +r /etc/powerssl/api/cert-key.pem
