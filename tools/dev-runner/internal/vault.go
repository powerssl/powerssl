package internal

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/ghodss/yaml"
	"github.com/hashicorp/vault/api"

	"powerssl.dev/tools/dev-runner/internal/component"
)

func vaultSecret() (config map[string]interface{}, err error) {
	var byt []byte
	if byt, err = ioutil.ReadFile("local/vault/secret.yaml"); err != nil {
		return nil, fmt.Errorf("config error: %s", err)
	}
	if err = yaml.Unmarshal(byt, &config); err != nil {
		return nil, fmt.Errorf("config error: %s", err)
	}
	return config, nil
}

func handleVault(of *Outlet) error {
	if _, err := os.Stat("local/vault/secret.yaml"); os.IsNotExist(err) {
		comp := component.Component{
			Name:    "vault",
			Command: "bin/powerutil",
			Args: []string{
				"vault",
				"--ca", "local/certs/ca.pem",
				"--ca-key", "local/certs/ca-key.pem",
			},
		}
		var cmd *exec.Cmd
		if cmd, _, err = makeCmd(comp, 0, of); err != nil {
			return err
		}
		if err = cmd.Start(); err != nil {
			return fmt.Errorf("failed to start %s: %s", comp.Command, err)
		}
		if err = cmd.Wait(); err != nil {
			return fmt.Errorf("failed to wait %s: %s", comp.Command, err)
		}
	} else {
		var config map[string]interface{}
		if config, err = vaultSecret(); err != nil {
			return err
		}
		var c *api.Client
		if c, err = newVaultClient(); err != nil {
			return err
		}
		for _, key := range config["keys"].([]interface{}) {
			var resp *api.SealStatusResponse
			if resp, err = c.Sys().Unseal(key.(string)); err != nil {
				return err
			}
			if !resp.Sealed {
				break
			}
		}
	}
	return nil
}

func newVaultClient() (*api.Client, error) {
	conf := api.DefaultConfig()
	if err := conf.ConfigureTLS(&api.TLSConfig{
		CAPath: "local/certs/ca.pem",
	}); err != nil {
		return nil, err
	}
	conf.Address = "https://localhost:8200"
	c, err := api.NewClient(conf)
	if err != nil {
		return nil, err
	}
	var config map[string]interface{}
	if config, err = vaultSecret(); err != nil {
		return nil, err
	}
	c.SetToken(config["rootToken"].(string))
	return c, nil
}

func vaultRoleID(name string) (_ string, err error) {
	c, err := newVaultClient()
	if err != nil {
		return "", err
	}
	var resp *api.Secret
	if resp, err = c.Logical().Read(fmt.Sprintf("auth/approle/role/%s/role-id", name)); err != nil {
		return "", err
	}
	roleID, ok := resp.Data["role_id"]
	if !ok {
		return "", fmt.Errorf(`key "role_id" not found`)
	}
	return roleID.(string), nil
}

func vaultSecretID(name string) (_ string, err error) {
	c, err := newVaultClient()
	if err != nil {
		return "", err
	}
	var resp *api.Secret
	if resp, err = c.Logical().Write(fmt.Sprintf("auth/approle/role/%s/secret-id", name), nil); err != nil {
		return "", err
	}
	secretID, ok := resp.Data["secret_id"]
	if !ok {
		return "", fmt.Errorf(`key "secret_id" not found`)
	}
	return secretID.(string), nil
}
