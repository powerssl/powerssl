package oauth2

type Config struct {
	AuthURI string
	GitHub  *GitHub
}

type GitHub struct {
	ClientID     string
	ClientSecret string
}
