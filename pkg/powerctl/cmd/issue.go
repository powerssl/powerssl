package cmd

import (
	"context"
	"strings"

	"github.com/spf13/cobra"

	"powerssl.io/pkg/apiserver/api"
)

var createIssueCmd = &cobra.Command{
	Use:   "issue",
	Short: "Create issue",
	Args:  validateParentArg("certificate"),
	Run: func(cmd *cobra.Command, args []string) {
		issue := &api.CertificateIssue{}
		if Filename != "" {
			loadResource(Filename, issue)
		} else {
			issue = makeIssue()
		}
		createIssue(args[0], issue)
	},
}

var deleteIssueCmd = &cobra.Command{
	Use:   "issue",
	Short: "Delete issue",
	Args:  validateNameArg,
	Run: func(cmd *cobra.Command, args []string) {
		deleteIssue(args[0])
	},
}

var getIssueCmd = &cobra.Command{
	Use:     "issue",
	Aliases: []string{"issues"},
	Short:   "Get issue",
	Example: `  powerctl get issue                   List all issues
  powerctl get issue certificates/42   List all issues of an certificate
  powerctl get issue 42                Get an issue
  powerctl get issues/42               Get an issue`,
	Args: cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			if strings.Contains(args[0], "/") {
				listIssue(args[0])
			} else {
				getIssue(nameArg("issue", args[0]))
			}
		} else {
			listIssue("")
		}
	},
}

var updateIssueCmd = &cobra.Command{
	Use:   "issue",
	Short: "Update issue",
	Args:  validateNameArg,
	Run: func(cmd *cobra.Command, args []string) {
		issue := &api.CertificateIssue{}
		if Filename != "" {
			loadResource(Filename, issue)
		} else {
			issue = makeIssue()
		}
		updateIssue(args[0], issue)
	},
}

func init() {
	createIssueCmd.Flags().StringVarP(&Filename, "filename", "f", "", "Filename to file to use to create the issue")

	updateIssueCmd.Flags().StringVarP(&Filename, "filename", "f", "", "Filename to file to use to create the issue")

	createCmd.AddCommand(createIssueCmd)
	deleteCmd.AddCommand(deleteIssueCmd)
	getCmd.AddCommand(getIssueCmd)
	updateCmd.AddCommand(updateIssueCmd)
}

func createIssue(parent string, issue *api.CertificateIssue) {
	client := newGRPCClient()
	createResource(func() (interface{}, error) {
		return client.CertificateIssue.Create(context.Background(), parent, issue)
	})
}

func deleteIssue(name string) {
	client := newGRPCClient()
	deleteResource(func() error {
		return client.CertificateIssue.Delete(context.Background(), name)
	})
}

func getIssue(name string) {
	client := newGRPCClient()
	getResource(func() (interface{}, error) {
		return client.CertificateIssue.Get(context.Background(), name)
	})
}

func listIssue(parent string) {
	client := newGRPCClient()
	listResource(func(pageToken string) (interface{}, string, error) {
		return client.CertificateIssue.List(context.Background(), parent, 0, pageToken)
	})
}

func updateIssue(name string, issue *api.CertificateIssue) {
	client := newGRPCClient()
	updateResource(func() (interface{}, error) {
		return client.CertificateIssue.Update(context.Background(), name, issue)
	})
}

func makeIssue() *api.CertificateIssue {
	return &api.CertificateIssue{}
}
