package oauth2

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	oauth2github "golang.org/x/oauth2/github"
)

const providerGitHub = "github"

func (o *OAuth2) initGitHub() {
	o.gitHubConfig = &oauth2.Config{
		ClientID:     o.cfg.GitHub.ClientID,
		ClientSecret: o.cfg.GitHub.ClientSecret,
		Endpoint:     oauth2github.Endpoint,
		RedirectURL:  o.cfg.AuthURI + "/callback",
		Scopes:       []string{"user:email"},
	}
}

func (o *OAuth2) gitHubAuthCodeURL() *string {
	authCodeURL := o.gitHubConfig.AuthCodeURL(o.gitHubState())
	return &authCodeURL
}

func (o *OAuth2) gitHubState() string {
	return oauthStateString + ":github"
}

func (o *OAuth2) gitHubUserInfo(ctx context.Context, state, code string) (*string, error) {
	if state != o.gitHubState() {
		return nil, fmt.Errorf("invalid oauth2 state")
	}
	token, err := o.gitHubConfig.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}
	ts := oauth2.StaticTokenSource(token)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	var user *github.User
	if user, _, err = client.Users.Get(ctx, ""); err != nil {
		return nil, err
	}
	return user.Login, nil
}
