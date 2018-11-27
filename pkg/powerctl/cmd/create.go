package cmd

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"github.com/ghodss/yaml"
	"github.com/spf13/cobra"
)

var Filename string

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create resource",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		filename := Filename

		if filename == "" {
			er("Please specify filename")
		}

		in, err := ioutil.ReadFile(filename)
		if err != nil {
			er(err)
		}

		switch filepath.Ext(filename) {
		case ".json":
		case ".yml", ".yaml":
			in, err = yaml.YAMLToJSON(in)
			if err != nil {
				er(err)
			}
		default:
			er("Unknown input format")
		}

		var resources []ResourceLoader
		if json.Unmarshal(in, &resources) != nil {
			var resource ResourceLoader
			if err := json.Unmarshal(in, &resource); err != nil {
				er(err)
			}
			resources = append(resources, resource)
		}

		client, err := NewGRPCClient()
		if err != nil {
			er(err)
		}

		out := make([]*Resource, len(resources))
		for i, resource := range resources {
			resourceHandler, err := Resources.Get(resource.Kind)
			if err != nil {
				er(err)
			}
			spec := resourceHandler.Spec()
			if err := json.Unmarshal(resource.Spec, &spec); err != nil {
				er(err)
			}
			res := Resource{
				Kind: resource.Kind,
				Meta: resource.Meta,
				Spec: spec,
			}
			if out[i], err = res.Create(client); err != nil {
				er(err)
			}
		}

		if len(out) > 1 {
			pr(out)
		} else if len(out) == 1 {
			pr(out[0])
		}
	},
}

func init() {
	createCmd.Flags().StringVarP(&Filename, "filename", "f", "", "Filename to file to use to create the resources")

	rootCmd.AddCommand(createCmd)
}
