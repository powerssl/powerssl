package cmd

import (
	"strings"

	"github.com/spf13/cobra"
)

var (
	getFuncs  map[string]func(string)
	listFuncs map[string]func(string)
)

func init() {
	getFuncs = make(map[string]func(string))
	getFuncs["acme-accounts"] = getACMEAccount
	getFuncs["acme-servers"] = getACMEServer
	getFuncs["acmeaccounts"] = getACMEAccount
	getFuncs["acmeservers"] = getACMEServer
	getFuncs["certificates"] = getCertificate
	getFuncs["issues"] = getIssue

	listFuncs = make(map[string]func(string))
	listFuncs["issues"] = listIssue
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get resource",
	Example: `  # List certificates
  powerctl get certificate

  # List issues
  powerctl get issue

  # List issues of a certificate
  powerctl get certificates/42/issues

  # Get a certificate
  powerctl get certificate 42
  powerctl get certificates/42

  # Get an issue
  powerctl get issue 42
  powerctl get issues/42
  powerctl get certificates/42/issues/42`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var (
			resourceType string
			name         string
			parent       string
		)
		s := strings.Split(args[0], "/")
		if len(s)%2 == 0 {
			resourceType = s[len(s)-2]
			name = args[0]
		} else {
			resourceType = s[len(s)-1]
			if len(s) > 1 {
				parent = strings.Join(s[:len(s)-1], "/")
			}
		}
		if name != "" {
			if getFuncs[resourceType] == nil {
				er("Unknown resource")
			}
			getFuncs[resourceType](name)
		} else {
			if listFuncs[resourceType] == nil {
				er("Unknown resource")
			}
			listFuncs[resourceType](parent)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
