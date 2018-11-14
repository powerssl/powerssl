package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/ghodss/yaml"
	"github.com/spf13/viper"
	"google.golang.org/grpc/status"

	apiserverclient "powerssl.io/pkg/apiserver/client"
	"powerssl.io/pkg/util"
	"powerssl.io/pkg/util/tracing"
)

var DisplayName string

func er(msg interface{}) {
	err, ok := msg.(error)
	if ok {
		status, ok := status.FromError(err)
		if ok {
			fmt.Fprintln(os.Stderr, fmt.Sprintf("RPC error:\n  Code:    %s\n  Message: %s\n", status.Code(), status.Message()))
		} else {
			fmt.Fprintln(os.Stderr, err)
		}
	} else {
		fmt.Fprintln(os.Stderr, msg)
	}
	os.Exit(1)
}

func pr(resource interface{}) {
	var (
		err error
		out []byte
	)
	switch Output {
	case "yaml":
		out, err = yaml.Marshal(resource)
	case "json":
		out, err = json.Marshal(resource)
	default:
		er("Unknown output format")
	}
	if err != nil {
		er(err)
	}
	fmt.Fprintln(os.Stdout, string(out))
}

func NewGRPCClient() (*apiserverclient.GRPCClient, error) {
	certFile := viper.GetString("ca-file")
	addr := viper.GetString("addr")
	insecure := viper.GetBool("insecure")
	insecureSkipTLSVerify := viper.GetBool("insecure-skip-tls-verify")
	serverNameOverride := viper.GetString("server-name-override")
	if addr == "" {
		return nil, errors.New("Provide addr")
	}
	if !insecure && !insecureSkipTLSVerify && certFile == "" {
		return nil, errors.New("Provide ca-file")
	}
	logger := util.NewLogger(os.Stdout)
	tracer, _, _ := tracing.NewNoopTracer("powerctl", logger)
	client, err := apiserverclient.NewGRPCClient(addr, certFile, serverNameOverride, insecure, insecureSkipTLSVerify, logger, tracer)
	if err != nil {
		return nil, err
	}
	return client, nil
}
