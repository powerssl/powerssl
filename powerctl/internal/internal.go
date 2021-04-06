package internal

import (
	"context"
	"errors"
	"github.com/spf13/viper"

	"powerssl.dev/common/log"
	"powerssl.dev/common/tracing"
	"powerssl.dev/common/transport"
	"powerssl.dev/sdk/apiserver"
)

func NewGRPCClient() (_ *apiserver.Client, err error) {
	addr := viper.GetString("addr")
	authToken := viper.GetString("auth-token")
	caFle := viper.GetString("ca-file")
	insecure := viper.GetBool("insecure")
	insecureSkipTLSVerify := viper.GetBool("insecure-skip-tls-verify")
	serverNameOverride := viper.GetString("server-name-override")
	if addr == "" {
		return nil, errors.New("provide addr")
	}
	if !insecure && !insecureSkipTLSVerify && caFle == "" {
		return nil, errors.New("provide ca-file")
	}
	var logger log.Logger
	if logger, err = log.NewLogger(false); err != nil {
		return nil, err
	}
	// TODO: logger.Sync()
	tracer, _, _ := tracing.NewNoopTracer("powerctl", logger)
	cfg := &transport.ClientConfig{
		Addr:                  addr,
		CAFile:                caFle,
		Insecure:              insecure,
		InsecureSkipTLSVerify: insecureSkipTLSVerify,
		ServerNameOverride:    serverNameOverride,
	}
	return apiserver.NewClient(context.TODO(), cfg, authToken, logger, tracer)
}
