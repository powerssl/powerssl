package vault

type AppRole struct {
	RoleID   string `flag:"roleID;;;vault app role ID" mapstructure:"role-id"`
	SecretID string `flag:"secretID;;;vault app role secret ID" mapstructure:"secret-id"`
}

type Config struct {
	AppRole AppRole `flag:"appRole;;;vault app role"`
	CAFile  string  `flag:"caFile;;;vault CA file"`
	Token   string  `flag:"token;;;vault token"`
	URL     string  `flag:"url;;;vault URL" validate:"url"`
}
