package server

import (
	"github.com/go-playground/validator/v10"
)

type Config struct {
	Addr              string `flag:"addr" flag-desc:"server addr"`
	Insecure          bool   `flag:"insecure" flag-desc:"server insecure"`
	CertFile          string `flag:"certFile" flag-desc:"server cert file"`
	KeyFile           string `flag:"keyFile" flag-desc:"server key file"`
	JWTPrivateKeyFile string `flag:"jwtPrivateKeyFile" flag-desc:"server JWT private key file"`
	WebappURI         string `flag:"webappURI" flag-desc:"webapp URI"`
}

func (cfg *Config) PreValidate(_ *validator.Validate) {}
