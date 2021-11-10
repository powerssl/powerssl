package config

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func Validate(cfg Config) error {
	cfg.Defaults()
	validate := validator.New()
	if err := validate.Struct(cfg); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}
		var fieldErrors []string
		for _, fieldError := range err.(validator.ValidationErrors) {
			fieldErrors = append(fieldErrors, "  "+convertFieldError(cfg, fieldError))
		}
		return errors.New("\n" + strings.Join(unique(fieldErrors), "\n") + "\n")
	}
	return nil
}

func convertFieldError(cfg Config, fieldError validator.FieldError) string {
	switch fieldError.Tag() {
	case "gt":
		if fieldError.Kind() == reflect.Slice {
			return fmt.Sprintf("%s needs to have more than %v values", convertStructNamespace(cfg, fieldError.StructNamespace()), fieldError.Param())
		}
		fallthrough
	case "hostname_port":
		return fmt.Sprintf("%s with value \"%v\", needs to be an hostname with port", convertStructNamespace(cfg, fieldError.StructNamespace()), fieldError.Value())
	case "required":
		return fmt.Sprintf("%s is required", convertStructNamespace(cfg, fieldError.StructNamespace()))
	case "uri":
		return fmt.Sprintf("%s with value \"%v\", needs to be an URI", convertStructNamespace(cfg, fieldError.StructNamespace()), fieldError.Value())
	case "url":
		return fmt.Sprintf("%s with value \"%v\", needs to be an URL", convertStructNamespace(cfg, fieldError.StructNamespace()), fieldError.Value())
	default:
		return fmt.Sprintf("namespace: %v, field: %v, struct_namespace: %v, struct_field: %v, tag: %v, actual_tag: %v, kind: %v, type: %v, value: %v, param: %v",
			fieldError.Namespace(),
			fieldError.Field(),
			fieldError.StructNamespace(),
			fieldError.StructField(),
			fieldError.Tag(),
			fieldError.ActualTag(),
			fieldError.Kind(),
			fieldError.Type(),
			fieldError.Value(),
			fieldError.Param())
	}
}

func convertStructNamespace(cfg Config, structNamespace string) string {
	paths := strings.Split(structNamespace, ".")
	return "--" + tagInformation(reflect.TypeOf(cfg).Elem(), paths[1:], "")
}

func unique(stringSlice []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
