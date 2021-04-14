package oauth2

import (
	"context"
	"fmt"
)

var oauthStateString = "foo" // TODO: Randomize it

func AuthCodeURL(provider string) (string, error) {
	switch provider {
	case "github":
		return gitHubAuthCodeURL(), nil
	default:
		return "", fmt.Errorf("provider %v not found", provider)
	}
}

func UserInfo(ctx context.Context, provider, state, code string) (string, error) {
	switch provider {
	case "github":
		return gitHubUserInfo(ctx, state, code)
	default:
		return "", fmt.Errorf("provider %v not found", provider)
	}
}
