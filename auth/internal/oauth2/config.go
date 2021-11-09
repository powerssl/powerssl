package oauth2

type Config struct {
	AuthURI string  `flag:"authURI;;;oAuth2 auth URI"`
	GitHub  GitHub `flag:"github;;;oAuth2 GitHub"`
}

type GitHub struct {
	ClientID     string `flag:"clientID;;;oAuth2 GitHub client ID"`
	ClientSecret string `flag:"clientSecret;;;oAuth2 GitHub client secret"`
}
