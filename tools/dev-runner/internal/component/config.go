package component

const GenerateAuthToken = "{{GENERATE_AUTH_TOKEN}}"

var Components = []Component{
	{
		Command: "bin/powerssl-apiserver",
		Args:    "serve --config apiserver/config.yaml",
	},
	{
		Command: "bin/powerssl-auth",
		Args:    "serve --config auth/config.yaml",
	},
	{
		Command: "bin/powerssl-controller",
		Args:    "serve --config controller/config.yaml",
		Env: Environment{
			"POWERSSL_AUTH_TOKEN": GenerateAuthToken,
		},
	},
	{
		Command: "bin/powerssl-grpcgateway",
		Args:    "serve --config grpcgateway/config.yaml",
	},
	{
		Command: "bin/powerssl-temporal",
		Args:    "run --config temporal/config.yaml",
	},
	{
		Command: "bin/powerssl-webapp",
		Args:    "serve --config webapp/config.yaml",
	},
	{
		Command: "bin/powerssl-worker",
		Args:    "run --config worker/config.yaml",
		Env: Environment{
			"POWERSSL_AUTH_TOKEN": GenerateAuthToken,
		},
	},
}
