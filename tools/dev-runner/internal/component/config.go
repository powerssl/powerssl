package component

var Components = []Component{
	component("apiserver", "serve", &Environment{
		"POWERSSL_SERVER_AUTHTOKEN":             GenerateAuthToken,
		"POWERSSL_SERVER_VAULT_APPROLEID":       GenerateVaultAppRoleID,
		"POWERSSL_SERVER_VAULT_APPROLESECRETID": GenerateVaultAppSecretID,
	}),
	component("auth", "serve", nil),
	component("controller", "serve", &Environment{
		"POWERSSL_SERVER_AUTHTOKEN":             GenerateAuthToken,
		"POWERSSL_SERVER_VAULT_APPROLEID":       GenerateVaultAppRoleID,
		"POWERSSL_SERVER_VAULT_APPROLESECRETID": GenerateVaultAppSecretID,
		"POWERSSL_VAULTCLIENT_APPROLEID":        GenerateVaultAppRoleID,
		"POWERSSL_VAULTCLIENT_APPROLESECRETID":  GenerateVaultAppSecretID,
	}),
	component("grpcgateway", "serve", nil),
	component("temporal", "run", nil),
	component("webapp", "serve", nil),
	component("worker", "run", &Environment{
		"POWERSSL_SERVER_AUTHTOKEN":            GenerateAuthToken,
		"POWERSSL_VAULTCLIENT_APPROLEID":       GenerateVaultAppRoleID,
		"POWERSSL_VAULTCLIENT_APPROLESECRETID": GenerateVaultAppSecretID,
	}),
}
