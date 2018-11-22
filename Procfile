apiserver: go run powerssl.io/cmd/powerssl-apiserver serve --addr localhost:5100 --controller-addr localhost:5300 --metrics-addr localhost:5190 --jwks-url http://localhost:5200/.well-known/jwks.json --controller-auth-token http://localhost:5200/service
auth: go run powerssl.io/cmd/powerssl-auth serve --addr localhost:5200 --metrics-addr localhost:5290
controller: go run powerssl.io/cmd/powerssl-controller serve --addr localhost:5300 --apiserver-addr localhost:5100 --metrics-addr localhost:5390 --jwks-url http://localhost:5200/.well-known/jwks.json --apiserver-auth-token http://localhost:5200/service
signer: go run powerssl.io/cmd/powerssl-signer serve --addr localhost:5400 --metrics-addr localhost:5490
integration-acme: sh -c "sleep 1 && go run powerssl.io/cmd/powerssl-integration-acme run --addr localhost:5300 --metrics-addr localhost:5590"
