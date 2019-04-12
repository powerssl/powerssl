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
	Command string      `json:"command,omitempty"`
	Args    string      `json:"args,omitempty"`
	Env     Environment `json:"env,omitempty"`
}

func (c Component) String() string {
	return strings.TrimPrefix(c.Command, "bin/")
}

func (c Component) Image() string {
	return strings.ReplaceAll(c.String(), "-", "/") + ":latest"
}
