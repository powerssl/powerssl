package server

type Config struct {
	Addr              string `flag:"addr;;;server addr"`
	Insecure          bool   `flag:"insecure;;;server insecure"`
	CertFile          string `flag:"certFile;;;server cert file"`
	KeyFile           string `flag:"keyFile;;;server key file"`
	JWTPrivateKeyFile string `flag:"jwtPrivateKeyFile;;;server JWT private key file"`
	WebappURI         string `flag:"webappURI;;;webapp URI"`
}
