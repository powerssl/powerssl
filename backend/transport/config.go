package transport

type Config struct {
	Addr       string `flag:"addr;;;server addr" validate:"required,hostname_port"`
	CAFile     string `flag:"caFile;;;server CA file"`
	CertFile   string `flag:"certFile;;;server Cert file"`
	CommonName string `flag:"commonName;;;server common name"`
	Insecure   bool   `flag:"insecure;;;server insecure"`
	KeyFile    string `flag:"keyFile;;;server key file"`
	VaultRole  string `flag:"vaultRole;;;server vault role"`
	VaultToken string `flag:"vaultToken;;;server vault token"`
	VaultURL   string `flag:"vaultURL;;;server vault URL"`
}
