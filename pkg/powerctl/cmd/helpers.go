package cmd

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"

	"powerssl.io/pkg/powerctl"
)

var (
	Name   string
	Parent string
)

func er(msg interface{}) {
	fmt.Println(msg)
	os.Exit(1)
}

func pr(resource interface{}) {
	byt, err := yaml.Marshal(resource)
	if err != nil {
		er(err)
	}
	fmt.Println(string(byt))
}

func newGRPCClient() *powerctl.GRPCClient {
	grpcAddr := viper.GetString("grpc.addr")
	return powerctl.NewGRPCClient(grpcAddr)
}

func createResource(createFunc func() (interface{}, error)) {
	resource, err := createFunc()
	if err != nil {
		er(err)
	}
	pr(resource)
}

func deleteResource(deleteFunc func() error) {
	if err := deleteFunc(); err != nil {
		er(err)
	}
}

func getResource(getFunc func() (interface{}, error)) {
	resource, err := getFunc()
	if err != nil {
		er(err)
	}
	pr(resource)
}

func listResource(listFunc func(pageToken string) (interface{}, string, error)) {
	var (
		pageToken string
		resources []interface{}
	)
	for {
		t, nextPageToken, err := listFunc(pageToken)
		if err != nil {
			er(err)
		}

		switch reflect.TypeOf(t).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(t)

			for i := 0; i < s.Len(); i++ {
				resources = append(resources, s.Index(i).Interface())
			}
		}

		if nextPageToken == "" {
			break
		}
		pageToken = nextPageToken
	}
	pr(resources)
}

func updateResource(updateFunc func() (interface{}, error)) {
	resource, err := updateFunc()
	if err != nil {
		er(err)
	}
	pr(resource)
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
