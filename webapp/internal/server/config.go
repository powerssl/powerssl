package server

type Config struct {
	APIAddr    string
	Addr       string
	AuthURI    string
	CertFile   string
	GRPCWebURI string
	Insecure   bool
	KeyFile    string
}
