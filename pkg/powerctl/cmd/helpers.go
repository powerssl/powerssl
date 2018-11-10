package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc/status"
	"gopkg.in/yaml.v2"

	apiserverclient "powerssl.io/pkg/apiserver/client"
	"powerssl.io/pkg/util"
	"powerssl.io/pkg/util/tracing"
)

var Filename string

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

func loadResource(filename string, resource interface{}) {
	in, err := ioutil.ReadFile(filename)
	if err != nil {
		er(err)
	}
	switch filepath.Ext(filename) {
	case ".yml", ".yaml":
		err = yaml.Unmarshal(in, resource)
	case ".json":
		err = json.Unmarshal(in, resource)
	default:
		err = errors.New("Unknown input format")
	}
	if err != nil {
		er(err)
	}
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

func newGRPCClient() *apiserverclient.GRPCClient {
	certFile := viper.GetString("ca-file")
	addr := viper.GetString("addr")
	insecure := viper.GetBool("insecure")
	insecureSkipTLSVerify := viper.GetBool("insecure-skip-tls-verify")
	serverNameOverride := viper.GetString("server-name-override")
	if addr == "" {
		er("Provide addr")
	}
	if !insecure && !insecureSkipTLSVerify && certFile == "" {
		er("Provide ca-file")
	}
	logger := util.NewLogger(os.Stdout)
	tracer, _, _ := tracing.NewNoopTracer("powerctl", logger)
	client, err := apiserverclient.NewGRPCClient(addr, certFile, serverNameOverride, insecure, insecureSkipTLSVerify, logger, tracer)
	if err != nil {
		logger.Log("transport", "gRPC", "during", "Connect", "err", err)
		os.Exit(1)
	}
	return client
}

func nameArg(resourcePlural, arg string) string {
	if strings.HasPrefix(arg, fmt.Sprint(resourcePlural, "/")) {
		return arg
	}
	return fmt.Sprint(resourcePlural, "/", arg)
}

func checkParentArg(parent string, resourcePlural string) error {
	sl := strings.Split(parent, "/")
	if len(sl) != 2 || sl[0] != resourcePlural {
		return errors.New("Invalid parent")
	}
	return nil
}

func validateParentArg(resourcePlural string) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("expected parent arg")
		}
		if len(args) > 1 {
			return errors.New("expected parent arg only")
		}
		return checkParentArg(args[0], resourcePlural)
	}

}

func validateNameArg(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("expected name arg")
	}
	if len(args) > 1 {
		return errors.New("expected name arg only")
	}
	return nil
}
