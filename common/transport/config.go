package transport

type ClientConfig struct {
	Addr                  string `flag:"addr;;;apiserver client addr" validate:"required,hostname_port"`
	CAFile                string `flag:"caFile;;;apiserver client CA file"`
	Insecure              bool   `flag:"insecure;;;apiserver client insecure"`
	InsecureSkipTLSVerify bool   `flag:"insecureSkipTLSVerify;;;apiserver client insecure skip TLS verify"`
	ServerNameOverride    string `flag:"serverNameOverride;;;apiserver client server name override"`
}
