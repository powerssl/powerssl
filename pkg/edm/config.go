package edm

import (
	"regexp"
	"time"

	"github.com/coreos/etcd/clientv3"
	"gopkg.in/go-playground/validator.v9"
)

type Config struct {
	Etcd            *clientv3.Config `validate:"required"`
	DataKeyPrefix   string           `validate:"required,is_valid_config_key_prefix,nefield=LookupKeyPrefix"`
	LookupKeyPrefix string           `validate:"required,is_valid_config_key_prefix,nefield=DataKeyPrefix"`
}

func NewConfig() *Config {
	return &Config{
		Etcd:            NewEtcdConfig(),
		DataKeyPrefix:   "/data",
		LookupKeyPrefix: "/lookup",
	}
}

func NewEtcdConfig() *clientv3.Config {
	return &clientv3.Config{
		DialTimeout: 10 * time.Second,
	}
}

func (c *Config) IsValid() bool {
	return c.Validate() == nil
}

func (c *Config) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("is_valid_config_key_prefix", validateConfigKeyPrefix)

	return validate.Struct(c)
}

func validateConfigKeyPrefix(fl validator.FieldLevel) bool {
	matched, _ := regexp.MatchString(`^\/([a-z0-9]{1,1})([a-z0-9-_]{0,})*[a-z0-9]$`, fl.Field().String())
	return matched
}
