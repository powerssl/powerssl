package resource

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"

	apiv1 "powerssl.dev/api/apiserver/v1"
	"powerssl.dev/powerctl/internal"
	"powerssl.dev/sdk/apiserver"
)

func NewCmdCreateACMEAccount() *cobra.Command {
	var (
		acmeServerID         string
		contacts             string
		termsOfServiceAgreed bool
	)

	cmd := internal.CmdWithClient(&cobra.Command{
		Use:     "acmeaccount",
		Short:   "Create ACME account",
		Args:    cobra.NoArgs,
		Example: `  powerctl create acmeaccount --agree-terms-of-service --contacts mailto:john.doe@example.com --acmeserver 42   Create ACME account within ACME server`,
	}, func(ctx context.Context, client *apiserver.Client, cmd *cobra.Command, args []string) error {
		apiACMEAccount := &apiv1.ACMEAccount{
			Contacts:             strings.Split(contacts, ","),
			TermsOfServiceAgreed: termsOfServiceAgreed,
		}
		var err error
		if apiACMEAccount, err = client.ACMEAccount.Create(ctx, &apiv1.CreateACMEAccountRequest{
			Parent:      fmt.Sprintf("acmeServers/%s", acmeServerID),
			AcmeAccount: apiACMEAccount,
		}); err != nil {
			return err
		}
		return FormatResource(acmeAccount{}.Encode(apiACMEAccount), cmd.OutOrStdout())
	})

	cmd.Flags().BoolVar(&termsOfServiceAgreed, "agree-terms-of-service", false, "Terms of Service agreed")
	cmd.Flags().StringVar(&contacts, "contacts", "", "Contact URLs (e.g. mailto:contact@example.com) (seperated by \",\")")
	cmd.Flags().StringVar(&acmeServerID, "acmeserver", "", "ACME Server")

	must(cmd.MarkFlagRequired("acmeserver"))

	return cmd
}

type acmeAccount struct{}

func (r acmeAccount) Create(ctx context.Context, client *apiserver.Client, resource *Resource) (*Resource, error) {
	spec := resource.Spec.(*acmeAccountSpec)
	acmeAccount := &apiv1.ACMEAccount{
		TermsOfServiceAgreed: spec.TermsOfServiceAgreed,
		Contacts:             spec.Contacts,
	}
	acmeAccount, err := client.ACMEAccount.Create(ctx, &apiv1.CreateACMEAccountRequest{
		Parent:      fmt.Sprintf("acmeServers/%s", spec.ACMEServer),
		AcmeAccount: acmeAccount,
	})
	if err != nil {
		return nil, err
	}
	return r.Encode(acmeAccount), nil
}

func (r acmeAccount) Delete(ctx context.Context, client *apiserver.Client, name string) error {
	_, err := client.ACMEAccount.Delete(ctx, &apiv1.DeleteACMEAccountRequest{
		Name: fmt.Sprintf("acmeServers/-/acmeAccounts/%s", name),
	})
	return err
}

func (r acmeAccount) Encode(acmeAccount *apiv1.ACMEAccount) *Resource {
	uid := strings.Split(acmeAccount.Name, "/")[3]
	acmeServer := strings.Split(acmeAccount.Name, "/")[1]
	return &Resource{
		Kind: "acmeaccount",
		Meta: &resourceMeta{
			UID:        uid,
			CreateTime: acmeAccount.GetCreateTime().AsTime(),
			UpdateTime: acmeAccount.GetUpdateTime().AsTime(),
		},
		Spec: &acmeAccountSpec{
			DisplayName:          acmeAccount.GetDisplayName(),
			Title:                acmeAccount.GetTitle(),
			Description:          acmeAccount.GetDescription(),
			ACMEServer:           acmeServer,
			TermsOfServiceAgreed: acmeAccount.GetTermsOfServiceAgreed(),
			Contacts:             acmeAccount.GetContacts(),
			AccountURL:           acmeAccount.GetAccountUrl(),
		},
	}
}

func (r acmeAccount) Get(ctx context.Context, client *apiserver.Client, name string) (*Resource, error) {
	acmeAccount, err := client.ACMEAccount.Get(ctx, &apiv1.GetACMEAccountRequest{
		Name: fmt.Sprintf("acmeServers/-/acmeAccounts/%s", name),
	})
	if err != nil {
		return nil, err
	}
	return r.Encode(acmeAccount), nil
}

func (r acmeAccount) List(ctx context.Context, client *apiserver.Client) ([]*Resource, error) {
	return listResource(func(pageToken string) ([]*Resource, string, error) {
		response, err := client.ACMEAccount.List(ctx, &apiv1.ListACMEAccountsRequest{
			Parent:    "parent",
			PageToken: pageToken,
			PageSize:  0,
		})
		if err != nil {
			return nil, "", err
		}
		a := make([]*Resource, len(response.GetAcmeAccounts()))
		for i, acmeAccount := range response.GetAcmeAccounts() {
			a[i] = r.Encode(acmeAccount)
		}
		return a, response.GetNextPageToken(), nil
	})
}

func (r acmeAccount) Spec() interface{} {
	return new(acmeAccountSpec)
}

func (r acmeAccount) Columns(resource *Resource) ([]string, []string) {
	spec := resource.Spec.(*acmeAccountSpec)
	return []string{
			"DISPLAY NAME",
			"DESCRIPTION",
			"ACME SERVER",
			"TOS AGREED",
			"CONTACTS",
		}, []string{
			fmt.Sprint(spec.DisplayName),
			fmt.Sprint(spec.Description),
			fmt.Sprint(spec.ACMEServer),
			fmt.Sprint(spec.TermsOfServiceAgreed),
			strings.Join(spec.Contacts, ","),
		}
}

func (r acmeAccount) Describe(ctx context.Context, client *apiserver.Client, resource *Resource, output io.Writer) (err error) {
	spec := resource.Spec.(*acmeAccountSpec)
	w := tabwriter.NewWriter(output, 0, 0, 1, ' ', tabwriter.TabIndent)
	_, _ = fmt.Fprintln(w, fmt.Sprintf("UID:\t%s", resource.Meta.UID))
	_, _ = fmt.Fprintln(w, fmt.Sprintf("Create Time:\t%s", resource.Meta.CreateTime))
	_, _ = fmt.Fprintln(w, fmt.Sprintf("Update Time:\t%s", resource.Meta.UpdateTime))
	_, _ = fmt.Fprintln(w, fmt.Sprintf("Display Name:\t%s", spec.DisplayName))
	_, _ = fmt.Fprintln(w, fmt.Sprintf("Title:\t%s", spec.Title))
	_, _ = fmt.Fprintln(w, fmt.Sprintf("Description:\t%s", spec.Description))
	_, _ = fmt.Fprintln(w, fmt.Sprintf("TOS Agreed:\t%v", spec.TermsOfServiceAgreed))
	_, _ = fmt.Fprintln(w, fmt.Sprintf("Contacts:\t%s", strings.Join(spec.Contacts, ",")))
	_, _ = fmt.Fprintln(w, fmt.Sprintf("Account URL:\t%s", spec.AccountURL))
	_, _ = fmt.Fprintln(w, "ACME Server:")
	acmeServer := &Resource{
		Kind: "acmeserver",
		Meta: &resourceMeta{
			UID: spec.ACMEServer,
		},
	}
	if acmeServer, err = acmeServer.Get(ctx, client); err != nil {
		return err
	}
	acmeServerDescription := new(bytes.Buffer)
	if err = acmeServer.Describe(ctx, client, acmeServerDescription); err != nil {
		return err
	}
	scanner := bufio.NewScanner(acmeServerDescription)
	for scanner.Scan() {
		_, _ = fmt.Fprintln(w, "", "", scanner.Text())
	}
	return w.Flush()
}

type acmeAccountSpec struct {
	DisplayName          string   `json:"displayName,omitempty"          yaml:"displayName,omitempty"`
	Title                string   `json:"directoryURL,omitempty"         yaml:"directoryURL,omitempty"`
	Description          string   `json:"integrationName,omitempty"      yaml:"integrationName,omitempty"`
	ACMEServer           string   `json:"acmeServer,omitempty"           yaml:"acmeServer,omitempty"`
	TermsOfServiceAgreed bool     `json:"termsOfServiceAgreed,omitempty" yaml:"termsOfServiceAgreed,omitempty"`
	Contacts             []string `json:"contacts,omitempty"             yaml:"contacts,omitempty"`
	AccountURL           string   `json:"accountURL,omitempty"           yaml:"accountURL,omitempty"`
}

func init() {
	resources.Add(acmeAccount{}, "aa")
}
