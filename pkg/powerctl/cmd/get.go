package cmd

import (
	"reflect"
	"strings"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get resource",
	Long:  `Get resource.`,
	Run: func(cmd *cobra.Command, args []string) {
		s := strings.Split(args[0], "/")
		var resourceType string
		if len(s)%2 == 0 {
			// get
			resourceType = s[len(s)-2]
			Name = args[0]
			switch resourceType {
			case "certificates":
				getCertificate(cmd, args)
			case "issues":
				getCertificateIssue(cmd, args)
			}
		} else {
			// list
			resourceType = s[len(s)-1]
			if len(s) > 1 {
				Parent = strings.Join(s[:len(s)-1], "/")
			}
			switch resourceType {
			case "certificates":
				listCertificate(cmd, args)
			case "issues":
				listCertificateIssue(cmd, args)
			}
		}
	},
}

var typeRegistry = make(map[string]reflect.Type)

func init() {
	rootCmd.AddCommand(getCmd)

	//	typeRegistry["certificate"] = reflect.TypeOf(certificateservice.Service)
}

func makeInstance(name string) interface{} {
	return reflect.New(typeRegistry[name]).Elem().Interface()
}
