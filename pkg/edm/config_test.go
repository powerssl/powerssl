package edm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/go-playground/validator.v9"
)

func TestConfigValidate(t *testing.T) {
	t.Run("default NewConfig", func(t *testing.T) {
		config := NewConfig()
		if assert.NoError(t, config.Validate()) {
			assert.True(t, config.IsValid())
		}
	})

	t.Run("Required fields are empty", func(t *testing.T) {
		config := &Config{
			Etcd:            nil,
			DataKeyPrefix:   "",
			LookupKeyPrefix: "",
		}

		err := config.Validate()
		if assert.Error(t, err) {
			validationErrors := err.(validator.ValidationErrors)
			fields := make([]string, len(validationErrors))
			tags := make([]string, len(validationErrors))

			for i, e := range validationErrors {
				fields[i] = e.Field()
				tags[i] = e.Tag()
			}

			assert.Len(t, fields, 3)
			assert.ElementsMatch(t, fields, []string{"Etcd", "DataKeyPrefix", "LookupKeyPrefix"})

			assert.Len(t, tags, 3)
			assert.ElementsMatch(t, tags, []string{"required", "required", "required"})
		}
	})

	t.Run("DataKeyPrefix and LookupKeyPrefix are equal", func(t *testing.T) {
		config := NewConfig()
		config.LookupKeyPrefix = config.DataKeyPrefix

		err := config.Validate()
		if assert.Error(t, err) {
			validationErrors := err.(validator.ValidationErrors)
			fields := make([]string, len(validationErrors))
			tags := make([]string, len(validationErrors))

			for i, e := range validationErrors {
				fields[i] = e.Field()
				tags[i] = e.Tag()
			}

			assert.Len(t, fields, 2)
			assert.ElementsMatch(t, fields, []string{"DataKeyPrefix", "LookupKeyPrefix"})

			assert.Len(t, tags, 2)
			assert.ElementsMatch(t, tags, []string{"nefield", "nefield"})
		}
	})

	t.Run("DataKeyPrefix has wrong format", func(t *testing.T) {
		config := NewConfig()

		var testCases func(config *Config, testStrs []string)
		testCases = func(config *Config, testStrs []string) {
			for _, str := range testStrs {
				config.DataKeyPrefix = str

				err := config.Validate()
				if assert.Error(t, err) {
					validationErrors := err.(validator.ValidationErrors)
					fields := make([]string, len(validationErrors))
					tags := make([]string, len(validationErrors))

					for i, e := range validationErrors {
						fields[i] = e.Field()
						tags[i] = e.Tag()
					}

					assert.Len(t, fields, 1)
					assert.Contains(t, fields, "DataKeyPrefix")
					assert.Len(t, tags, 1)
					assert.Contains(t, tags, "is_valid_config_key_prefix")
				}
			}
		}

		testCases(config, []string{"/", "key", "/key/", "/-key"})
	})

	t.Run("LookupKeyPrefix has wrong format", func(t *testing.T) {
		config := NewConfig()

		var testCases func(config *Config, testStrs []string)
		testCases = func(config *Config, testStrs []string) {
			for _, str := range testStrs {
				config.LookupKeyPrefix = str

				err := config.Validate()
				if assert.Error(t, err) {
					validationErrors := err.(validator.ValidationErrors)
					fields := make([]string, len(validationErrors))
					tags := make([]string, len(validationErrors))

					for i, e := range validationErrors {
						fields[i] = e.Field()
						tags[i] = e.Tag()
					}

					assert.Len(t, fields, 1)
					assert.Contains(t, fields, "LookupKeyPrefix")
					assert.Len(t, tags, 1)
					assert.Contains(t, tags, "is_valid_config_key_prefix")
				}
			}
		}

		testCases(config, []string{"/", "key", "/key/", "/-key"})
	})
}
