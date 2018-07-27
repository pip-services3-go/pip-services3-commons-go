package config

import (
	"strings"

	"github.com/pip-services-go/pip-services-commons-go/data"
)

type ConfigParams struct {
	data.StringValueMap
}

func NewEmptyConfigParams() *ConfigParams {
	return &ConfigParams{
		StringValueMap: *data.NewEmptyStringValueMap(),
	}
}

func NewConfigParams(values map[string]string) *ConfigParams {
	return &ConfigParams{
		StringValueMap: *data.NewStringValueMap(values),
	}
}

func NewConfigParamsFromValue(value interface{}) *ConfigParams {
	return &ConfigParams{
		StringValueMap: *data.NewStringValueMapFromValue(value),
	}
}

func NewConfigParamsFromTuples(tuples ...interface{}) *ConfigParams {
	return &ConfigParams{
		StringValueMap: *data.NewStringValueMapFromTuplesArray(tuples),
	}
}

func NewConfigParamsFromTuplesArray(tuples []interface{}) *ConfigParams {
	return &ConfigParams{
		StringValueMap: *data.NewStringValueMapFromTuplesArray(tuples),
	}
}

func NewConfigParamsFromString(line string) *ConfigParams {
	return &ConfigParams{
		StringValueMap: *data.NewStringValueMapFromString(line),
	}
}

func NewConfigParamsFromMaps(maps ...map[string]string) *ConfigParams {
	return &ConfigParams{
		StringValueMap: *data.NewStringValueMapFromMaps(maps...),
	}
}

func (c *ConfigParams) GetSectionNames() []string {
	sections := []string{}

	for key := range c.Values() {
		pos := strings.Index(key, ".")
		section := key
		if pos > 0 {
			section = key[0:pos]
		}

		// Perform case sensitive search
		found := false
		for index := 0; index < len(sections); index++ {
			if section == sections[index] {
				found = true
				break
			}
		}

		if !found {
			sections = append(sections, section)
		}
	}

	return sections
}

func (c *ConfigParams) GetSection(section string) *ConfigParams {
	result := NewEmptyConfigParams()
	prefix := section + "."

	for key := range (*c).Values() {
		// Prevents exception on the next line
		if len(key) < len(prefix) {
			continue
		}

		// Perform case sensitive match
		keyPrefix := key[0:len(prefix)]
		if keyPrefix == prefix {
			sectionKey := key[len(prefix):]
			result.Put(sectionKey, (*c).Get(key))
		}
	}

	return result
}

func (c *ConfigParams) AddSection(section string, sectionParams *ConfigParams) {
	if section == "" {
		panic("Section name cannot be empty")
	}

	if sectionParams != nil {
		for key := range sectionParams.Values() {
			sectionKey := key

			if len(sectionKey) > 0 {
				sectionKey = section + "." + sectionKey
			} else {
				sectionKey = section
			}

			value := (*sectionParams).Get(key)

			(*c).Put(sectionKey, value)
		}
	}
}

func (c *ConfigParams) Override(configParams *ConfigParams) *ConfigParams {
	return NewConfigParamsFromMaps(c.Values(), configParams.Values())
}

func (c *ConfigParams) SetDefaults(defaults *ConfigParams) *ConfigParams {
	return NewConfigParamsFromMaps(defaults.Values(), c.Values())
}
