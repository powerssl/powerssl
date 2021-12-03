package oauth2

import (
	"github.com/go-playground/validator/v10"
)

type Config struct {
	AuthURI string `flag:"authURI" flag-desc:"oAuth2 auth URI"`
	GitHub  GitHub `flag:"github" flag-desc:"oAuth2 GitHub"`
}

func (cfg *Config) PreValidate(validate *validator.Validate) {
	cfg.GitHub.PreValidate(validate)
}

type GitHub struct {
	ClientID     string `flag:"clientID" flag-desc:"oAuth2 GitHub client ID"`
	ClientSecret string `flag:"clientSecret" flag-desc:"oAuth2 GitHub client secret"`
}

func (cfg *GitHub) PreValidate(_ *validator.Validate) {}
