package cmd

import (
	"context"
	"strings"

	"github.com/spf13/cobra"

	"powerssl.io/pkg/apiserver/api"
)

var (
	ACMEServer           string
	Contacts             string
	TermsOfServiceAgreed bool
)

var createACMEAccountCmd = &cobra.Command{
	Use:   "acmeaccount",
	Short: "Create ACME account",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		acmeAccount := &api.ACMEAccount{}
		if Filename != "" {
			loadResource(Filename, acmeAccount)
		} else {
			acmeAccount = makeACMEAccount()
		}
		createACMEAccount(acmeAccount)
	},
}

var deleteACMEAccountCmd = &cobra.Command{
	Use:   "acmeaccount",
	Short: "Delete ACME account",
	Args:  validateNameArg,
	Run: func(cmd *cobra.Command, args []string) {
		deleteACMEAccount(nameArg("acmeaccounts", args[0]))
	},
}

var getACMEAccountCmd = &cobra.Command{
	Use:     "acmeaccount",
	Aliases: []string{"acmeaccounts"},
	Short:   "Get ACME account",
	Example: `  powerctl get acmeaccount       List all ACME accounts
  powerctl get acmeaccount 42    Get an ACME account
  powerctl get acmeaccounts/42   Get an ACME account`,
	Args: cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			getACMEAccount(nameArg("acmeaccounts", args[0]))
		} else {
			listACMEAccount()
		}
	},
}

var updateACMEAccountCmd = &cobra.Command{
	Use:   "acmeaccount",
	Short: "Update ACME account",
	Args:  validateNameArg,
	Run: func(cmd *cobra.Command, args []string) {
		acmeAccount := &api.ACMEAccount{}
		if Filename != "" {
			loadResource(Filename, acmeAccount)
		} else {
			acmeAccount = makeACMEAccount()
		}
		updateACMEAccount(nameArg("acmeaccounts", args[0]), acmeAccount)
	},
}

func init() {
	createACMEAccountCmd.Flags().BoolVarP(&TermsOfServiceAgreed, "agree-terms-of-service", "", false, "Terms of Service agreed")
	createACMEAccountCmd.Flags().StringVarP(&ACMEServer, "acme-server", "", "", "ACME server name")
	createACMEAccountCmd.Flags().StringVarP(&Contacts, "contacts", "", "", "Contact URLs (e.g. mailto:contact@example.com) (seperated by \",\")")
	createACMEAccountCmd.Flags().StringVarP(&Filename, "filename", "f", "", "Filename to file to use to create the ACME account")

	updateACMEAccountCmd.Flags().StringVarP(&Filename, "filename", "f", "", "Filename to file to use to update the ACME account")
	updateACMEAccountCmd.Flags().StringVarP(&Contacts, "contacts", "", "", "Contact URLs (e.g. mailto:contact@example.com) (seperated by \",\")")

	createCmd.AddCommand(createACMEAccountCmd)
	deleteCmd.AddCommand(deleteACMEAccountCmd)
	getCmd.AddCommand(getACMEAccountCmd)
	updateCmd.AddCommand(updateACMEAccountCmd)
}

func createACMEAccount(acmeAccount *api.ACMEAccount) {
	client := newGRPCClient()
	createResource(func() (interface{}, error) {
		return client.ACMEAccount.Create(context.Background(), acmeAccount)
	})
}

func deleteACMEAccount(name string) {
	client := newGRPCClient()
	deleteResource(func() error {
		return client.ACMEAccount.Delete(context.Background(), name)
	})
}

func getACMEAccount(name string) {
	client := newGRPCClient()
	getResource(func() (interface{}, error) {
		return client.ACMEAccount.Get(context.Background(), name)
	})
}

func listACMEAccount() {
	client := newGRPCClient()
	listResource(func(pageToken string) (interface{}, string, error) {
		return client.ACMEAccount.List(context.Background(), 0, pageToken)
	})
}

func updateACMEAccount(name string, acmeAccount *api.ACMEAccount) {
	client := newGRPCClient()
	updateResource(func() (interface{}, error) {
		return client.ACMEAccount.Update(context.Background(), name, acmeAccount)
	})
}

func makeACMEAccount() *api.ACMEAccount {
	return &api.ACMEAccount{
		ACMEServer:           ACMEServer,
		Contacts:             strings.Split(Contacts, ","),
		TermsOfServiceAgreed: TermsOfServiceAgreed,
	}
}
