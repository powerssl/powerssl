package transport // import "powerssl.dev/common/transport"

import (
	"github.com/go-playground/validator/v10"
)

type Config struct {
	Addr                  string `flag:"addr" flag-desc:"client addr" validate:"required,hostname_port"`
	CAFile                string `flag:"caFile" flag-desc:"client CA file"`
	Insecure              bool   `flag:"insecure" flag-desc:"client insecure"`
	InsecureSkipTLSVerify bool   `flag:"insecureSkipTLSVerify" flag-desc:"client insecure skip TLS verify"`
	ServerNameOverride    string `flag:"serverNameOverride" flag-desc:"client server name override"`
}

func (cfg *Config) PreValidate(validate *validator.Validate) {
	validate.RegisterStructValidation(structValidation, Config{})
}

func structValidation(sl validator.StructLevel) {
	cfg := sl.Current().Interface().(Config)
	if !cfg.Insecure && !cfg.InsecureSkipTLSVerify && cfg.CAFile == "" {
		sl.ReportError(cfg.CAFile, "caFile", "CAFile", "", "")
	}
}
