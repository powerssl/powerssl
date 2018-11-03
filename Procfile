apiserver: sh -c "sleep 1 && go run powerssl.io/cmd/powerssl-apiserver serve --addr localhost:5100 --controller-addr localhost:5200 --no-metrics"
controller: go run powerssl.io/cmd/powerssl-controller serve --addr localhost:5200 --metrics-addr localhost:9090
integration-acme: sh -c "sleep 1 && go run powerssl.io/cmd/powerssl-integration-acme run --addr localhost:5200"
signer: go run powerssl.io/cmd/powerssl-signer serve --addr localhost:5300 --no-metrics
