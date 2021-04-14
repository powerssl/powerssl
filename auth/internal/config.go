package internal

import (
	"github.com/go-playground/validator/v10"

	"powerssl.dev/common"
)

type Config struct {
	Addr     string `validate:"required,hostname_port"`
	Auth struct {
		URI string `validate:"required,uri"`
	}
	Insecure bool
	JWT      struct {
		PrivateKeyFile string `mapstructure:"private-key-file" validate:"required"`
	}
	Metrics struct {
		Addr string
	}
	OAuth2 struct {
		GitHub struct {
			ClientID     string `mapstructure:"client-id"`
			ClientSecret string `mapstructure:"client-secret"`
		} `mapstructure:"github"`
	} `mapstructure:"oauth2"`
	TLS struct {
		CertFile       string `mapstructure:"cert-file"`
		PrivateKeyFile string `mapstructure:"private-key-file"`
	}
	WebApp struct {
		URI string `mapstructure:"uri" validate:"required"`
	}
}

func (cfg *Config) Validate() error {
	validate := validator.New()
	return common.ValidateConfig(validate, cfg)
}
