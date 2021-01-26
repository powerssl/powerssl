package powerutil

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	schemaVersion   = "v96"
	temporalVersion = "1.5.1"
)

func RunTemporalMigrate(docker bool, host, password, plugin, port, temporalDatabase, user, visibilityDatabase string) error {
	var stdoutStderr string
	var err error

	stdoutStderr, err = executeTemporalSQLTool(docker, "--endpoint", host, "--port", port, "--user", user, "--password", password, "--plugin", plugin, "create-database", "--database", temporalDatabase)
	if err != nil && !strings.Contains(stdoutStderr, "already exists") {
		fmt.Println(err)
		os.Exit(1)
	}
	stdoutStderr, err = executeTemporalSQLTool(docker, "--endpoint", host, "--port", port, "--user", user, "--password", password, "--database", temporalDatabase, "--plugin", plugin, "setup-schema", "--version", "0.0")
	if err != nil && !strings.Contains(stdoutStderr, "already exists") {
		fmt.Println(err)
		os.Exit(1)
	}
	_, err = executeTemporalSQLTool(docker, "--endpoint", host, "--port", port, "--user", user, "--password", password, "--database", temporalDatabase, "--plugin", plugin, "update-schema", "--schema-dir", fmt.Sprintf("/etc/temporal/schema/postgresql/%s/temporal/versioned", schemaVersion))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	stdoutStderr, err = executeTemporalSQLTool(docker, "--endpoint", host, "--port", port, "--user", user, "--password", password, "--plugin", plugin, "create-database", "--database", visibilityDatabase)
	if err != nil && !strings.Contains(stdoutStderr, "already exists") {
		fmt.Println(err)
		os.Exit(1)
	}
	stdoutStderr, err = executeTemporalSQLTool(docker, "--endpoint", host, "--port", port, "--user", user, "--password", password, "--database", visibilityDatabase, "--plugin", plugin, "setup-schema", "--version", "0.0")
	if err != nil && !strings.Contains(stdoutStderr, "already exists") {
		fmt.Println(err)
		os.Exit(1)
	}
	_, err = executeTemporalSQLTool(docker, "--endpoint", host, "--port", port, "--user", user, "--password", password, "--database", visibilityDatabase, "--plugin", plugin, "update-schema", "--schema-dir", fmt.Sprintf("/etc/temporal/schema/postgresql/%s/visibility/versioned", schemaVersion))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return nil
}

func RunTemporalRegisterNamespace(docker, tlsEnableHostVerification bool, address, namespace, tlsCertPath, tlsKeyPath, tlsCAPath, tlsServerName string) error {
	args := []string{
		"--address", address,
		"--namespace", namespace,
		"--tls_cert_path", tlsCertPath,
		"--tls_key_path", tlsKeyPath,
		"--tls_ca_path", tlsCAPath,
		"--tls_server_name", tlsServerName,
	}
	if tlsEnableHostVerification {
		args = append(args, "--tls_enable_host_verification")
	}
	args = append(args, "namespace", "register")
	_, err := executeTCTL(docker, args...)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return nil
}

func executeTemporalSQLTool(docker bool, args ...string) (string, error) {
	fmt.Println("temporal-sql-tool", strings.Join(args, " "))
	command := "temporal-sql-tool"
	if docker {
		command = "docker"
		args = append([]string{
			"run", "--rm",
			"--entrypoint", "/usr/local/bin/temporal-sql-tool",
			"--net", "host",
			fmt.Sprintf("temporalio/admin-tools:%s", temporalVersion),
		}, args...)
	}
	cmd := exec.Command(command, args...)
	byt, err := cmd.CombinedOutput()
	stdoutStderr := string(byt)
	fmt.Print(stdoutStderr)
	return stdoutStderr, err
}

func executeTCTL(docker bool, args ...string) (string, error) {
	fmt.Println("tctl", strings.Join(args, " "))
	command := "tctl"
	if docker {
		command = "docker"
		args = append([]string{
			"run", "--rm",
			"--entrypoint", "/usr/local/bin/tctl",
			"--net", "host",
			fmt.Sprintf("temporalio/admin-tools:%s", temporalVersion),
		}, args...)
		fmt.Println(strings.Join(args, " "))
	}
	cmd := exec.Command(command, args...)
	byt, err := cmd.CombinedOutput()
	stdoutStderr := string(byt)
	fmt.Print(stdoutStderr)
	return stdoutStderr, err
}
