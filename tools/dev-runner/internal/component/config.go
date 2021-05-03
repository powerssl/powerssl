package component

const Generate = "{{GENERATE}}"

var Components = []Component{
	{
		Name:    "powerssl-apiserver",
		Command: "bin/powerssl-apiserver",
		Args:    "serve --config apiserver/config.yaml",
		Env: Environment{
			"POWERSSL_VAULT_ROLE_ID":   Generate,
			"POWERSSL_VAULT_SECRET_ID": Generate,
		},
	},
	{
		Name:    "powerssl-auth",
		Command: "bin/powerssl-auth",
		Args:    "serve --config auth/config.yaml",
	},
	{
		Name:    "powerssl-controller",
		Command: "bin/powerssl-controller",
		Args:    "serve --config controller/config.yaml",
		Env: Environment{
			"POWERSSL_AUTH_TOKEN":      Generate,
			"POWERSSL_VAULT_ROLE_ID":   Generate,
			"POWERSSL_VAULT_SECRET_ID": Generate,
		},
	},
	{
		Name:    "powerssl-grpcgateway",
		Command: "bin/powerssl-grpcgateway",
		Args:    "serve --config grpcgateway/config.yaml",
	},
	{
		Name:    "powerssl-temporal",
		Command: "bin/powerssl-temporal",
		Args:    "run --config temporal/config.yaml",
	},
	{
		Name:    "powerssl-webapp",
		Command: "bin/powerssl-webapp",
		Args:    "serve --config webapp/config.yaml",
	},
	{
		Name:    "powerssl-worker",
		Command: "bin/powerssl-worker",
		Args:    "run --config worker/config.yaml",
		Env: Environment{
			"POWERSSL_AUTH_TOKEN":      Generate,
			"POWERSSL_VAULT_ROLE_ID":   Generate,
			"POWERSSL_VAULT_SECRET_ID": Generate,
		},
	},
}
