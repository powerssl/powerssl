package validator // import "powerssl.dev/common/validator"

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type binding struct {
	viper string
	cobra string
}

func ValidateConfig(validate *validator.Validate, s interface{}) (err error) {
	if err = validate.Struct(s); err == nil {
		return nil
	}
	var errs []string
	for _, fieldError := range err.(validator.ValidationErrors) {
		errs = append(errs, "  "+convertFieldError(fieldError))
	}
	return errors.New("\n" + strings.Join(unique(errs), "\n") + "\n")
}

func convertFieldError(fieldError validator.FieldError) string {
	switch fieldError.Tag() {
	case "gt":
		if fieldError.Kind() == reflect.Slice {
			return fmt.Sprintf("%s needs to have more than %v values", convertStructNamespace(fieldError.StructNamespace()), fieldError.Param())
		}
		fallthrough
	case "hostname_port":
		return fmt.Sprintf("%s with value \"%v\", needs to be an hostname with port", convertStructNamespace(fieldError.StructNamespace()), fieldError.Value())
	case "required":
		return fmt.Sprintf("%s is required", convertStructNamespace(fieldError.StructNamespace()))
	case "uri":
		return fmt.Sprintf("%s with value \"%v\", needs to be an URI", convertStructNamespace(fieldError.StructNamespace()), fieldError.Value())
	case "url":
		return fmt.Sprintf("%s with value \"%v\", needs to be an URL", convertStructNamespace(fieldError.StructNamespace()), fieldError.Value())
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

func convertStructNamespace(structNamespace string) string {
	b, ok := map[string]binding{
		"APIServer.Addr":                              {"apiserver.addr", "apiserver-addr"},
		"APIServerClientConfig.Addr":                  {"apiserver.addr", "apiserver-addr"},
		"APIServerClientConfig.CAFile":                {"apiserver.ca-file", "ca-file"},
		"APIServerClientConfig.Insecure":              {"apiserver.insecure", "apiserver-insecure"},
		"APIServerClientConfig.InsecureSkipTLSVerify": {"apiserver.insecure-skip-tls-verify", "apiserver-insecure-skip-tls-verify"},
		"APIServerClientConfig.ServerNameOverride":    {"apiserver.server-name-override", "apiserver-server-name-override"},
		"Addr":                            {"addr", "addr"},
		"Auth.URI":                        {"auth.uri", "auth-uri"},
		"AuthToken":                       {"auth-token", "auth-token"},
		"ConfigDir":                       {"config-dir", "config-dir"},
		"ControllerClientConfig.Addr":     {"controller.addr", "controller-addr"},
		"ControllerClientConfig.CAFile":   {"controller.ca-file", "ca-file"},
		"ControllerClientConfig.Insecure": {"controller.insecure", "controller-insecure"},
		"ControllerClientConfig.InsecureSkipTLSVerify": {"controller.insecure-skip-tls-verify", "controller-insecure-skip-tls-verify"},
		"ControllerClientConfig.ServerNameOverride":    {"controller.server-name-override", "controller-server-name-override"},
		"Env":                            {"env", "env"},
		"GRPCWeb.URI":                    {"grpcweb.uri", "grpcweb-uri"},
		"JWT.PrivateKeyFile":             {"jwt.private-key-file", "jwt-private-key-file"},
		"Metrics.Addr":                   {"metrics.addr", "metrics-addr"},
		"Services":                       {"services", "service"},
		"TemporalClientConfig.CAFile":    {"temporal.ca-file", "ca-file"},
		"TemporalClientConfig.HostPort":  {"temporal.host-port", "temporal-host-port"},
		"TemporalClientConfig.Namespace": {"temporal.namespace", "temporal-namespace"},
		"Tracer":                         {"tracer", "tracer"},
		"VaultClientConfig.CAFile":       {"vault.ca-file", "ca-file"},
		"VaultClientConfig.Token":        {"vault.token", "vault-token"},
		"VaultClientConfig.URL":          {"vault.url", "vault-url"},
		"WebApp.URI":                     {"webapp.uri", "webapp-uri"},
		"Zone":                           {"zone", "zone"},
	}[strings.TrimPrefix(structNamespace, "Config.")]
	if !ok {
		return structNamespace
	}
	return fmt.Sprintf("--%s", b.cobra)
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
