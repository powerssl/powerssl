package component

var Components = []Component{
	{
		Command: "bin/powerssl-apiserver",
		Args:    "serve",
		Env: Environment{
			"POWERSSL_ADDR":                  "localhost:8082",
			"POWERSSL_METRICS_ADDR":          "localhost:9092",
			"POWERSSL_CONTROLLER_ADDR":       "localhost:8083",
			"POWERSSL_CONTROLLER_AUTH_TOKEN": "http://localhost:8081/service",
			"POWERSSL_CA_FILE":               "local/certs/ca.pem",
			"POWERSSL_COMMON_NAME":           "localhost",
			"POWERSSL_DB_DIALECT":            "sqlite3",
			"POWERSSL_DB_CONNECTION":         "local/powerssl.sqlite3",
			"POWERSSL_JWKS_URL":              "http://localhost:8081/.well-known/jwks.json",
			"POWERSSL_VAULT_URL":             "https://localhost:8200",
			"POWERSSL_VAULT_TOKEN":           "powerssl-apiserver",
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
			"POWERSSL_ADDR":                 "localhost:8083",
			"POWERSSL_METRICS_ADDR":         "localhost:9093",
			"POWERSSL_APISERVER_ADDR":       "localhost:8082",
			"POWERSSL_APISERVER_AUTH_TOKEN": "http://localhost:8081/service",
			"POWERSSL_CA_FILE":              "local/certs/ca.pem",
			"POWERSSL_COMMON_NAME":          "localhost",
			"POWERSSL_JWKS_URL":             "http://localhost:8081/.well-known/jwks.json",
			"POWERSSL_VAULT_URL":            "https://localhost:8200",
			"POWERSSL_VAULT_TOKEN":          "powerssl-controller",
		},
	},
	{
		Command: "bin/powerssl-grpcgateway",
		Args:    "serve",
		Env: Environment{
			"POWERSSL_ADDR":                               "localhost:8085",
			"POWERSSL_METRICS_ADDR":                       "localhost:9095",
			"POWERSSL_APISERVER_ADDR":                     "localhost:8082",
			"POWERSSL_CA_FILE":                            "local/certs/ca.pem",
			"POWERSSL_APISERVER_INSECURE_SKIP_TLS_VERIFY": "true", // TODO: Does not work without yet.
		},
	},
	{
		Command: "bin/powerssl-signer",
		Args:    "serve",
		Env: Environment{
			"POWERSSL_ADDR":         "localhost:8084",
			"POWERSSL_METRICS_ADDR": "localhost:9094",
			"POWERSSL_CA_FILE":      "local/certs/ca.pem",
			"POWERSSL_COMMON_NAME":  "localhost",
			"POWERSSL_VAULT_URL":    "https://localhost:8200",
			"POWERSSL_VAULT_TOKEN":  "powerssl-signer",
		},
	},
	{
		Command: "bin/powerssl-webapp",
		Args:    "serve",
		Env: Environment{

			"POWERSSL_ADDR":         "localhost:8080",
			"POWERSSL_API_ADDR":     "localhost:8082",
			"POWERSSL_AUTH_URI":     "http://localhost:8081",
			"POWERSSL_METRICS_ADDR": "localhost:9090",
		},
	},
}
