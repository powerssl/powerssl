package oauth2

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
)

var oauthStateString = "foo" // TODO: Randomize it

type OAuth2 struct {
	cfg          *Config
	gitHubConfig *oauth2.Config
}

func New(cfg *Config) *OAuth2 {
	o := &OAuth2{
		cfg: cfg,
	}
	o.init()
	return o
}

func (o *OAuth2) AuthCodeURL(provider string) (*string, error) {
	switch provider {
	case providerGitHub:
		return o.gitHubAuthCodeURL(), nil
	default:
		return nil, fmt.Errorf("provider %v not found", provider)
	}
}

func (o *OAuth2) UserInfo(ctx context.Context, provider, state, code string) (*string, error) {
	switch provider {
	case providerGitHub:
		return o.gitHubUserInfo(ctx, state, code)
	default:
		return nil, fmt.Errorf("provider %v not found", provider)
	}
}

func (o *OAuth2) init() {
	if o.cfg.GitHub != nil {
		o.initGitHub()
	}
}
