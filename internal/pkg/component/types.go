package component

import (
	"fmt"
	"strings"
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
	Args    string      `json:"args,omitempty"`
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
