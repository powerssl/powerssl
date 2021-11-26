package transport

type Config struct {
	Addr                  string `flag:"addr;;;client addr" validate:"required,hostname_port"`
	CAFile                string `flag:"caFile;;;client CA file"`
	Insecure              bool   `flag:"insecure;;;client insecure"`
	InsecureSkipTLSVerify bool   `flag:"insecureSkipTLSVerify;;;client insecure skip TLS verify"`
	ServerNameOverride    string `flag:"serverNameOverride;;;client server name override"`
}
