package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"powerssl.io/pkg/apiserver/api"
)

var (
	DirectoryURL       string
	IntegrationName    string
	LetsEncrypt        bool
	LetsEncryptStaging bool
)

var createACMEServerCmd = &cobra.Command{
	Use:   "acmeserver",
	Short: "Create ACME server",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		acmeServer := &api.ACMEServer{}
		if Filename != "" {
			loadResource(Filename, acmeServer)
		} else {
			acmeServer = makeACMEServer()
		}
		createACMEServer(acmeServer)
	},
}

var deleteACMEServerCmd = &cobra.Command{
	Use:   "acmeserver",
	Short: "Delete ACME server",
	Args:  validateNameArg,
	Run: func(cmd *cobra.Command, args []string) {
		deleteACMEServer(nameArg("acmeservers", args[0]))
	},
}

var getACMEServerCmd = &cobra.Command{
	Use:     "acmeserver",
	Aliases: []string{"acmeservers"},
	Short:   "Get ACME server",
	Example: `  powerctl get acmeserver       List all ACME servers
  powerctl get acmeserver 42    Get an ACME server
  powerctl get acmeservers/42   Get an ACME server`,
	Args: cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			getACMEServer(nameArg("acmeservers", args[0]))
		} else {
			listACMEServer()
		}
	},
}

var updateACMEServerCmd = &cobra.Command{
	Use:   "acmeserver",
	Short: "Update ACME server",
	Args:  validateNameArg,
	Run: func(cmd *cobra.Command, args []string) {
		acmeServer := &api.ACMEServer{}
		if Filename != "" {
			loadResource(Filename, acmeServer)
		} else {
			acmeServer = makeACMEServer()
		}
		updateACMEServer(nameArg("acmeservers", args[0]), acmeServer)
	},
}

func init() {
	createACMEServerCmd.Flags().BoolVarP(&LetsEncrypt, "letsencrypt", "", false, "Let's Encrypt defaults")
	createACMEServerCmd.Flags().BoolVarP(&LetsEncryptStaging, "letsencrypt-staging", "", false, "Let's Encrypt staging defaults")
	createACMEServerCmd.Flags().StringVarP(&DisplayName, "display-name", "", "", "Display name")
	createACMEServerCmd.Flags().StringVarP(&DirectoryURL, "directory-url", "", "", "Directory URL")
	createACMEServerCmd.Flags().StringVarP(&Filename, "filename", "f", "", "Filename to file to use to create the ACME server")
	createACMEServerCmd.Flags().StringVarP(&IntegrationName, "integration-name", "", "", "Integration name")

	updateACMEServerCmd.Flags().StringVarP(&DisplayName, "display-name", "", "", "Display name")
	updateACMEServerCmd.Flags().StringVarP(&DirectoryURL, "directory-url", "", "", "Directory URL")
	updateACMEServerCmd.Flags().StringVarP(&Filename, "filename", "f", "", "Filename to file to use to create the ACME server")
	updateACMEServerCmd.Flags().StringVarP(&IntegrationName, "integration-name", "", "", "Integration name")

	createCmd.AddCommand(createACMEServerCmd)
	deleteCmd.AddCommand(deleteACMEServerCmd)
	getCmd.AddCommand(getACMEServerCmd)
	updateCmd.AddCommand(updateACMEServerCmd)
}

func createACMEServer(acmeServer *api.ACMEServer) {
	client := newGRPCClient()
	createResource(func() (interface{}, error) {
		return client.ACMEServer.Create(context.Background(), acmeServer)
	})
}

func deleteACMEServer(name string) {
	client := newGRPCClient()
	deleteResource(func() error {
		return client.ACMEServer.Delete(context.Background(), name)
	})
}

func getACMEServer(name string) {
	client := newGRPCClient()
	getResource(func() (interface{}, error) {
		return client.ACMEServer.Get(context.Background(), name)
	})
}

func listACMEServer() {
	client := newGRPCClient()
	listResource(func(pageToken string) (interface{}, string, error) {
		return client.ACMEServer.List(context.Background(), 0, pageToken)
	})
}

func updateACMEServer(name string, acmeServer *api.ACMEServer) {
	client := newGRPCClient()
	updateResource(func() (interface{}, error) {
		return client.ACMEServer.Update(context.Background(), name, acmeServer)
	})
}

func makeACMEServer() *api.ACMEServer {
	if LetsEncrypt {
		return &api.ACMEServer{
			DirectoryURL:    "https://acme-v02.api.letsencrypt.org/directory",
			DisplayName:     "Let's Encrypt",
			IntegrationName: "acme",
		}
	}
	if LetsEncryptStaging {
		return &api.ACMEServer{
			DirectoryURL:    "https://acme-staging-v02.api.letsencrypt.org/directory",
			DisplayName:     "Let's Encrypt Staging",
			IntegrationName: "acme",
		}
	}
	return &api.ACMEServer{
		DirectoryURL:    DirectoryURL,
		DisplayName:     DisplayName,
		IntegrationName: IntegrationName,
	}
}
