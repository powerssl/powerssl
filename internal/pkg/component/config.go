package component

var Components = []Component{
	{
		Command: "bin/powerssl-apiserver",
		Args:    "serve",
		Env: Environment{
			"POWERSSL_ADDR":            "localhost:8082",
			"POWERSSL_AUTH_TOKEN":      "http://localhost:8081/service",
			"POWERSSL_CA_FILE":         "local/certs/ca.pem",
			"POWERSSL_COMMON_NAME":     "localhost",
			"POWERSSL_CONTROLLER_ADDR": "localhost:8083",
			"POWERSSL_DB_CONNECTION":   "postgresql://powerssl:powerssl@localhost:5432/powerssl?sslmode=disable",
			"POWERSSL_JWKS_URL":        "http://localhost:8081/.well-known/jwks.json",
			"POWERSSL_METRICS_ADDR":    "localhost:9092",
			"POWERSSL_VAULT_TOKEN":     "powerssl-apiserver",
			"POWERSSL_VAULT_URL":       "https://localhost:8200",
		},
	},
	{
		Command: "bin/powerssl-auth",
		Args:    "serve",
		Env: Environment{
			"POWERSSL_ADDR":                 "localhost:8081",
			"POWERSSL_JWT_PRIVATE_KEY_FILE": "local/certs/ca-key.pem",
			"POWERSSL_METRICS_ADDR":         "localhost:9091",
			"POWERSSL_WEBAPP_URI":           "http://localhost:8080",
		},
	},
	{
		Command: "bin/powerssl-controller",
		Args:    "serve",
		Env: Environment{
			"POWERSSL_ADDR":           "localhost:8083",
			"POWERSSL_APISERVER_ADDR": "localhost:8082",
			"POWERSSL_AUTH_TOKEN":     "http://localhost:8081/service",
			"POWERSSL_CA_FILE":        "local/certs/ca.pem",
			"POWERSSL_COMMON_NAME":    "localhost",
			"POWERSSL_JWKS_URL":       "http://localhost:8081/.well-known/jwks.json",
			"POWERSSL_METRICS_ADDR":   "localhost:9093",
			"POWERSSL_VAULT_TOKEN":    "powerssl-controller",
			"POWERSSL_VAULT_URL":      "https://localhost:8200",
		},
	},
	{
		Command: "bin/powerssl-grpcgateway",
		Args:    "serve",
		Env: Environment{
			"POWERSSL_ADDR":                               "localhost:8085",
			"POWERSSL_APISERVER_ADDR":                     "localhost:8082",
			"POWERSSL_APISERVER_INSECURE_SKIP_TLS_VERIFY": "true", // TODO: Does not work without yet.
			"POWERSSL_CA_FILE":                            "local/certs/ca.pem",
			"POWERSSL_METRICS_ADDR":                       "localhost:9095",
		},
	},
	{
		Command: "bin/powerssl-signer",
		Args:    "serve",
		Env: Environment{
			"POWERSSL_ADDR":         "localhost:8084",
			"POWERSSL_CA_FILE":      "local/certs/ca.pem",
			"POWERSSL_COMMON_NAME":  "localhost",
			"POWERSSL_METRICS_ADDR": "localhost:9094",
			"POWERSSL_VAULT_TOKEN":  "powerssl-signer",
			"POWERSSL_VAULT_URL":    "https://localhost:8200",
		},
	},
	{
		Command: "bin/powerssl-temporalserver",
		Args:    "run",
		Env: Environment{
			"POWERSSL_CONFIG_DIR": "configs/temporal",
			"POWERSSL_ENV":        "development",
		},
	},
	{
		Command: "bin/powerssl-webapp",
		Args:    "serve",
		Env: Environment{

			"POWERSSL_ADDR":         "localhost:8080",
			"POWERSSL_API_ADDR":     "localhost:8082",
			"POWERSSL_AUTH_URI":     "http://localhost:8081",
			"POWERSSL_GRPC_WEB_URI": "https://localhost:8086",
			"POWERSSL_METRICS_ADDR": "localhost:9090",
		},
	},
	{
		Command: "bin/powerssl-worker",
		Args:    "run",
		Env: Environment{
			"POWERSSL_APISERVER_ADDR": "localhost:8082",
			"POWERSSL_AUTH_TOKEN":     "http://localhost:8081/service",
			"POWERSSL_CA_FILE":        "local/certs/ca.pem",
			"POWERSSL_COMMON_NAME":    "localhost",
			"POWERSSL_JWKS_URL":       "http://localhost:8081/.well-known/jwks.json",
			"POWERSSL_METRICS_ADDR":   "localhost:9096",
			"POWERSSL_VAULT_TOKEN":    "powerssl-worker",
			"POWERSSL_VAULT_URL":      "https://localhost:8200",
		},
	},
}
