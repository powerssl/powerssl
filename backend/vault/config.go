package vault

type Config struct {
	AppRoleID       string `flag:"appRoleID;;;vault app role ID"`
	AppRoleSecretID string `flag:"appRoleSecretID;;;vault app role secret ID"`
	CAFile          string `flag:"caFile;;;vault CA file"`
	Token           string `flag:"token;;;vault token"`
	URL             string `flag:"url;;;vault URL" validate:"url"`
}
