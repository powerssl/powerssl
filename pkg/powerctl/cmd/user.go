package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"powerssl.io/pkg/apiserver/api"
	apiserverclient "powerssl.io/pkg/apiserver/client"
)

var User user

type UserSpec struct {
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	UserName    string `json:"userName,omitempty"    yaml:"userName,omitempty"`
}

type user struct{}

func (r user) Create(client *apiserverclient.GRPCClient, resource *Resource) (*Resource, error) {
	spec := resource.Spec.(*UserSpec)
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
		Meta: &ResourceMeta{
			UID:        uid,
			CreateTime: user.CreateTime,
			UpdateTime: user.UpdateTime,
		},
		Spec: &UserSpec{
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
	return new(UserSpec)
}

func (r user) Columns(resource *Resource) ([]string, []string) {
	spec := resource.Spec.(*UserSpec)
	return []string{
			"DISPLAY NAME",
			"USER NAME",
		}, []string{
			fmt.Sprint(spec.DisplayName),
			fmt.Sprint(spec.UserName),
		}
}

var UserName string

var createUserCmd = &cobra.Command{
	Use:     "user",
	Aliases: []string{"user"},
	Short:   "Create ACME server",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := NewGRPCClient()
		if err != nil {
			er(err)
		}
		user := &api.User{
			DisplayName: DisplayName,
			UserName:    UserName,
		}
		user, err = client.User.Create(context.Background(), user)
		if err != nil {
			er(err)
		}
		pr(User.Encode(user))
	},
}

func init() {
	Resources.Add(User, "u")

	createUserCmd.Flags().StringVarP(&DisplayName, "display-name", "", "", "Display name")
	createUserCmd.Flags().StringVarP(&Filename, "filename", "f", "", "Filename to file to use to create the ACME server")
	createUserCmd.Flags().StringVarP(&UserName, "user-name", "", "", "User name")

	createCmd.AddCommand(createUserCmd)
}
