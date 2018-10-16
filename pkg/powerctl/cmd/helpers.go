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
	"gopkg.in/yaml.v2"

	"powerssl.io/pkg/powerctl"
)

var Filename string

func er(msg interface{}) {
	fmt.Println(msg)
	os.Exit(1)
}

func loadResource(filename string, resource interface{}) error {
	in, err := ioutil.ReadFile(filename)
	if err != nil {
		er(err)
	}
	switch filepath.Ext(filename) {
	case ".yml", ".yaml":
		return yaml.Unmarshal(in, resource)
	case ".json":
		return json.Unmarshal(in, resource)
	default:
		er("Unknown input format")
	}
	return nil
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
	fmt.Println(string(out))
}

func newGRPCClient() *powerctl.GRPCClient {
	certFile := viper.GetString("ca-file")
	grpcAddr := viper.GetString("grpc-addr")
	insecure := viper.GetBool("insecure")
	insecureSkipTLSVerify := viper.GetBool("insecure-skip-tls-verify")
	serverNameOverride := viper.GetString("server-name-override")
	if grpcAddr == "" {
		er("Provide grpc-addr")
	}
	if !insecure && !insecureSkipTLSVerify && certFile == "" {
		er("Provide ca-file")
	}
	return powerctl.NewGRPCClient(grpcAddr, certFile, serverNameOverride, insecure, insecureSkipTLSVerify)
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
