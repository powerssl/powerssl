package server

type Config struct {
	APIAddr    string `flag:"apiAddr;;;server API addr"`
	Addr       string `flag:"addr;;;server addr"`
	AuthURI    string `flag:"authURI;;;server auth URI"`
	CertFile   string `flag:"certFile;;;server cert file"`
	GRPCWebURI string `flag:"grpcWebURI;;;server gRPC web URI"`
	Insecure   bool   `flag:"insecure;;;server insecure"`
	KeyFile    string `flag:"keyFile;;;server key file"`
}
