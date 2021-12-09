package component

import (
	"fmt"
	"strings"
)

const (
	GenerateAuthToken        = "{{GENERATE_AUTHTOKEN}}"
	GenerateVaultAppRoleID   = "{{GENERATE_VAULTAPPROLEID}}"
	GenerateVaultAppSecretID = "{{GENERATE_VAULTAPPSECRETID}}"
)

type Environment map[string]string

func (e Environment) Environ() []string {
	var env []string
	for k, v := range e {
		env = append(env, fmt.Sprintf("%s=%s", k, v))
	}
	return env
}

type Component struct {
	Name    string      `json:"name,omitempty"`
	Command string      `json:"command,omitempty"`
	Args    []string    `json:"args,omitempty"`
	Env     Environment `json:"env,omitempty"`
}

func (c Component) Image() string {
	return strings.ReplaceAll(c.String(), "-", "/") + ":latest"
}

func (c Component) String() string {
	if c.Name == "" {
		return strings.TrimPrefix(c.Command, "bin/")
	}
	return c.Name
}

func args(component, cmd string) []string {
	return []string{"--config", component + "/config.yaml", cmd}
}

func command(component string) string {
	return "bin/powerssl-" + component
}

func name(component string) string {
	return "powerssl-" + component
}

func component(component, cmd string, env *Environment) Component {
	c := Component{
		Name:    name(component),
		Command: command(component),
		Args:    args(component, cmd),
	}
	if env != nil {
		c.Env = *env
	}
	return c
}
