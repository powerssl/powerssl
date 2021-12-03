package vault

import (
	"github.com/go-playground/validator/v10"
)

type Config struct {
	AppRoleID       string `flag:"appRoleID" flag-desc:"vault app role ID"`
	AppRoleSecretID string `flag:"appRoleSecretID" flag-desc:"vault app role secret ID"`
	CAFile          string `flag:"caFile" flag-desc:"vault CA file"`
	Token           string `flag:"token" flag-desc:"vault token"`
	URL             string `flag:"url" flag-desc:"vault URL" validate:"url"`
}

func (cfg *Config) PreValidate(validate *validator.Validate) {}
