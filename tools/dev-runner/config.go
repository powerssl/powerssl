package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"

	"powerssl.io/tools/dev-runner/internal"
)

type environment map[string]string

func (e environment) Env() []string {
	env := os.Environ()
	for k, v := range e {
		env = append(env, fmt.Sprintf("%s=%s", k, v))
	}
	return env
}

type component struct {
	command string
	args    string
	env     environment
}

func (c component) Name() string {
	return strings.TrimPrefix(c.command, "bin/")
}

func (c component) Command(of *internal.Outlet, idx int) (*exec.Cmd, *sync.WaitGroup) {
	cmd := exec.Command(c.command, c.arg()...)
	cmd.Env = c.env.Env()

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		of.ErrorOutput(fmt.Sprintf("error: %s", err))
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		of.ErrorOutput(fmt.Sprintf("error: %s", err))
	}

	pipeWait := new(sync.WaitGroup)
	pipeWait.Add(2)
	go of.LineReader(pipeWait, c.Name(), idx, stdout, false)
	go of.LineReader(pipeWait, c.Name(), idx, stderr, true)

	return cmd, pipeWait
}

func (c component) arg() []string {
	return strings.Fields(c.args)
}

var components = []component{
	{
		command: "bin/powerssl-apiserver",
		args:    "serve",
		env: environment{
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
		command: "bin/powerssl-auth",
		args:    "serve",
		env: environment{
			"POWERSSL_ADDR":                 "localhost:8081",
			"POWERSSL_JWT_PRIVATE_KEY_FILE": "local/certs/ca-key.pem",
			"POWERSSL_METRICS_ADDR":         "localhost:9091",
			"POWERSSL_WEBAPP_URI":           "http://localhost:8080",
		},
	},
	{
		command: "bin/powerssl-controller",
		args:    "serve",
		env: environment{
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
		command: "bin/powerssl-signer",
		args:    "serve",
		env: environment{
			"POWERSSL_ADDR":         "localhost:8084",
			"POWERSSL_METRICS_ADDR": "localhost:9094",
			"POWERSSL_CA_FILE":      "local/certs/ca.pem",
			"POWERSSL_COMMON_NAME":  "localhost",
			"POWERSSL_VAULT_URL":    "https://localhost:8200",
			"POWERSSL_VAULT_TOKEN":  "powerssl-signer",
		},
	},
	{
		command: "bin/powerssl-webapp",
		args:    "serve",
		env: environment{

			"POWERSSL_ADDR":         "localhost:8080",
			"POWERSSL_API_ADDR":     "localhost:8082",
			"POWERSSL_AUTH_URI":     "http://localhost:8081",
			"POWERSSL_METRICS_ADDR": "localhost:9090",
		},
	},
}
