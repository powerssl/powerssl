package internal

import (
	"context"
	"errors"

	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"powerssl.dev/common/log"
	"powerssl.dev/common/tracer"
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
	var logger *zap.SugaredLogger
	if logger, err = log.New(log.Config{
		Env: "production",
	}); err != nil {
		return nil, err
	}
	// TODO: logger.Sync()
	var trace opentracing.Tracer
	if trace, _, err = tracer.New(tracer.Config{
		Component:      "powerctl",
		Implementation: "",
	}, logger); err != nil {
		return nil, err
	}
	return apiserver.NewClient(context.Background(), apiserver.Config{
		AuthToken: authToken,
		Client: transport.Config{
			Addr:                  addr,
			CAFile:                caFle,
			Insecure:              insecure,
			InsecureSkipTLSVerify: insecureSkipTLSVerify,
			ServerNameOverride:    serverNameOverride,
		},
	}, logger, trace)
}
