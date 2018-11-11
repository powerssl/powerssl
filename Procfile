apiserver: sh -c "sleep 0.1 && go run powerssl.io/cmd/powerssl-apiserver serve --addr localhost:5100 --controller-addr localhost:5200 --metrics-addr localhost:5190"
controller: go run powerssl.io/cmd/powerssl-controller serve --addr localhost:5200 --metrics-addr localhost:5290
signer: go run powerssl.io/cmd/powerssl-signer serve --addr localhost:5300 --metrics-addr localhost:5390
integration-acme: sh -c "sleep 0.1 && go run powerssl.io/cmd/powerssl-integration-acme run --addr localhost:5200 --metrics-addr localhost:5490"
