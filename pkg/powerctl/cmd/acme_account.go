package cmd

import (
	"context"
	"strings"

	"github.com/spf13/cobra"

	"powerssl.io/pkg/apiserver/api"
)

var (
	Contacts             string
	TermsOfServiceAgreed bool
)

var createACMEAccountCmd = &cobra.Command{
	Use:     "acmeaccount [PARENT]",
	Short:   "Create ACME account",
	Args:    validateParentArg("acmeServer"),
	Example: `  powerctl create acmeaccount acmeServers/42 --agree-terms-of-service --contacts mailto:john.doe@example.com   Create ACME account within ACME server`,
	Run: func(cmd *cobra.Command, args []string) {
		acmeAccount := &api.ACMEAccount{}
		if Filename != "" {
			loadResource(Filename, acmeAccount)
		} else {
			acmeAccount = makeACMEAccount()
		}
		createACMEAccount(args[0], acmeAccount)
	},
}

var deleteACMEAccountCmd = &cobra.Command{
	Use:   "acmeaccount",
	Short: "Delete ACME account",
	Args:  validateNameArg,
	Run: func(cmd *cobra.Command, args []string) {
		deleteACMEAccount(args[0])
	},
}

var getACMEAccountCmd = &cobra.Command{
	Use:     "acmeaccount",
	Aliases: []string{"acmeaccounts"},
	Short:   "Get ACME account",
	Example: `  powerctl get acmeaccount       List all ACME accounts
  powerctl get acmeaccount acmeservers/42    List all ACME accounts of an ACME server
  powerctl get acmeaccount 42                Get an ACME account
  powerctl get acmeaccounts/42               Get an ACME account`,
	Args: cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			if strings.Contains(args[0], "/") {
				getACMEAccount(args[0])
			} else {
				getACMEAccount(nameArg("acmeaccount", args[0]))
			}
		} else {
			listACMEAccount("")
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
		updateACMEAccount(args[0], acmeAccount)
	},
}

func init() {
	createACMEAccountCmd.Flags().BoolVarP(&TermsOfServiceAgreed, "agree-terms-of-service", "", false, "Terms of Service agreed")
	createACMEAccountCmd.Flags().StringVarP(&Contacts, "contacts", "", "", "Contact URLs (e.g. mailto:contact@example.com) (seperated by \",\")")
	createACMEAccountCmd.Flags().StringVarP(&Filename, "filename", "f", "", "Filename to file to use to create the ACME account")

	updateACMEAccountCmd.Flags().StringVarP(&Filename, "filename", "f", "", "Filename to file to use to update the ACME account")
	updateACMEAccountCmd.Flags().StringVarP(&Contacts, "contacts", "", "", "Contact URLs (e.g. mailto:contact@example.com) (seperated by \",\")")

	createCmd.AddCommand(createACMEAccountCmd)
	deleteCmd.AddCommand(deleteACMEAccountCmd)
	getCmd.AddCommand(getACMEAccountCmd)
	updateCmd.AddCommand(updateACMEAccountCmd)
}

func createACMEAccount(parent string, acmeAccount *api.ACMEAccount) {
	client := newGRPCClient()
	createResource(func() (interface{}, error) {
		return client.ACMEAccount.Create(context.Background(), parent, acmeAccount)
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

func listACMEAccount(parent string) {
	client := newGRPCClient()
	listResource(func(pageToken string) (interface{}, string, error) {
		return client.ACMEAccount.List(context.Background(), parent, 0, pageToken)
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
		Contacts:             strings.Split(Contacts, ","),
		TermsOfServiceAgreed: TermsOfServiceAgreed,
	}
}
