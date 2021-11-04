package server

type Config struct {
	Addr              string
	Insecure          bool
	CertFile          string
	KeyFile           string
	JWTPrivateKeyFile string
	WebappURI         string
}
