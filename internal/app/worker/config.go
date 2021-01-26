package worker

import (
	"gopkg.in/go-playground/validator.v9"

	"powerssl.dev/powerssl/internal/pkg/transport"
)

type APIServerClientConfig = transport.ClientConfig

type TemporalClientConfig struct {
	HostPort  string
	Namespace string
}

type VaultClientConfig struct {
	Token  string
	URL    string
	CAFile string
}

type Config struct {
	APIServerClientConfig *APIServerClientConfig
	AuthToken             string
	MetricsAddr           string
	TemporalClientConfig  *TemporalClientConfig
	Tracer                string
	VaultClientConfig     *VaultClientConfig
}

func (cfg *Config) Validate() error {
	validate := validator.New()
	validate.RegisterStructValidation(transport.ClientConfigValidator, transport.ClientConfig{})
	return validate.Struct(cfg)
}
