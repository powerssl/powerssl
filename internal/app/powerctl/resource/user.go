package resource

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"

	"powerssl.io/internal/app/powerctl"
	"powerssl.io/pkg/apiserver/api"
	apiserverclient "powerssl.io/pkg/apiserver/client"
)

type user struct{}

func (r user) Create(client *apiserverclient.GRPCClient, resource *Resource) (*Resource, error) {
	spec := resource.Spec.(*userSpec)
	user := &api.User{
		DisplayName: spec.DisplayName,
		UserName:    spec.UserName,
	}
	user, err := client.User.Create(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return r.Encode(user), nil
}

func (r user) Delete(client *apiserverclient.GRPCClient, name string) error {
	return client.User.Delete(context.Background(), fmt.Sprintf("users/%s", name))
}

func (r user) Encode(user *api.User) *Resource {
	uid := strings.Split(user.Name, "/")[1]
	return &Resource{
		Kind: "user",
		Meta: &resourceMeta{
			UID:        uid,
			CreateTime: user.CreateTime,
			UpdateTime: user.UpdateTime,
		},
		Spec: &userSpec{
			DisplayName: user.DisplayName,
			UserName:    user.UserName,
		},
	}
}

func (r user) Get(client *apiserverclient.GRPCClient, name string) (*Resource, error) {
	user, err := client.User.Get(context.Background(), fmt.Sprintf("users/%s", name))
	if err != nil {
		return nil, err
	}
	return r.Encode(user), nil
}

func (r user) List(client *apiserverclient.GRPCClient) ([]*Resource, error) {
	return listResource(func(pageToken string) ([]*Resource, string, error) {
		users, nextPageToken, err := client.User.List(context.Background(), 0, pageToken)
		if err != nil {
			return nil, nextPageToken, err
		}
		a := make([]*Resource, len(users))
		for i, user := range users {
			a[i] = r.Encode(user)
		}
		return a, nextPageToken, nil
	})
}

func (r user) Spec() interface{} {
	return new(userSpec)
}

func (r user) Columns(resource *Resource) ([]string, []string) {
	spec := resource.Spec.(*userSpec)
	return []string{
			"DISPLAY NAME",
			"USER NAME",
		}, []string{
			fmt.Sprint(spec.DisplayName),
			fmt.Sprint(spec.UserName),
		}
}

func (r user) Describe(client *apiserverclient.GRPCClient, resource *Resource, output io.Writer) (err error) {
	spec := resource.Spec.(*userSpec)
	w := tabwriter.NewWriter(output, 0, 0, 1, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, fmt.Sprintf("UID:\t%s", resource.Meta.UID))
	fmt.Fprintln(w, fmt.Sprintf("Create Time:\t%s", resource.Meta.CreateTime))
	fmt.Fprintln(w, fmt.Sprintf("Update Time:\t%s", resource.Meta.UpdateTime))
	fmt.Fprintln(w, fmt.Sprintf("Display Name:\t%s", spec.DisplayName))
	fmt.Fprintln(w, fmt.Sprintf("User Name:\t%s", spec.UserName))
	w.Flush()
	return nil
}

type userSpec struct {
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	UserName    string `json:"userName,omitempty"    yaml:"userName,omitempty"`
}

func NewCmdCreateUser() *cobra.Command {
	var client *apiserverclient.GRPCClient
	var (
		displayName string
		userName    string
	)

	cmd := &cobra.Command{
		Use:     "user",
		Aliases: []string{"user"},
		Short:   "Create ACME server",
		Args:    cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) (err error) {
			client, err = powerctl.NewGRPCClient()
			return err
		},
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			apiUser := &api.User{
				DisplayName: displayName,
				UserName:    userName,
			}
			if apiUser, err = client.User.Create(context.Background(), apiUser); err != nil {
				return err
			}
			return FormatResource(user{}.Encode(apiUser), os.Stdout)
		},
	}

	cmd.Flags().StringVarP(&displayName, "display-name", "", "", "Display name")
	cmd.Flags().StringVarP(&userName, "user-name", "", "", "User name")

	return cmd
}

func init() {
	resources.Add(user{}, "u")
}
