package server

import (
	"github.com/go-playground/validator/v10"
)

type Config struct {
	APIAddr    string `flag:"apiAddr" flag-desc:"server API addr"`
	Addr       string `flag:"addr" flag-desc:"server addr"`
	AuthURI    string `flag:"authURI" flag-desc:"server auth URI"`
	CertFile   string `flag:"certFile" flag-desc:"server cert file"`
	GRPCWebURI string `flag:"grpcWebURI" flag-desc:"server gRPC web URI"`
	Insecure   bool   `flag:"insecure" flag-desc:"server insecure"`
	KeyFile    string `flag:"keyFile" flag-desc:"server key file"`
}

func (cfg *Config) PreValidate(_ *validator.Validate) {}
