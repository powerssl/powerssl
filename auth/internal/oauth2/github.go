package oauth2

import (
	"context"
	"fmt"
	oauth2github "golang.org/x/oauth2/github"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var gitHubOauthConfig *oauth2.Config

func InitGitHubOauth2Config(authURI, clientID, clientSecret string) {
	gitHubOauthConfig = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     oauth2github.Endpoint,
		RedirectURL:  authURI + "/callback",
		Scopes:       []string{"user:email"},
	}
}

func gitHubAuthCodeURL() string {
	return gitHubOauthConfig.AuthCodeURL(gitHubState())
}

func gitHubState() string {
	return oauthStateString + ":github"
}

func gitHubUserInfo(ctx context.Context, state, code string) (string, error) {
	if state != gitHubState() {
		return "", fmt.Errorf("invalid oauth2 state")
	}
	token, err := gitHubOauthConfig.Exchange(ctx, code)
	if err != nil {
		return "", fmt.Errorf("code exchange failed: %s", err.Error())
	}
	ts := oauth2.StaticTokenSource(token)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	var user *github.User
	if user, _, err = client.Users.Get(ctx, ""); err != nil {
		return "", err
	}
	return *user.Login, nil
}
