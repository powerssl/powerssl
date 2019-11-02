package edm

import (
	"path"
	"regexp"

	"gopkg.in/go-playground/validator.v9"
)

// Full key name for the resource which is primarily used to identify resourc entries in etcd.
func ResourceFullKey(r ResourceInterface) string {
	return path.Join(r.GetParent(), r.Kind(), r.GetName())
}

func validateKeyName(fl validator.FieldLevel) bool {
	matched, _ := regexp.MatchString(`^[a-z0-9-_]{1,}$`, fl.Field().String())
	return matched
}
