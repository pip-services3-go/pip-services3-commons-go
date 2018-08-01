package config

import (
	"github.com/pip-services-go/pip-services-commons-go/refer"
)

type TNameResolver struct{}

var NameResolver *TNameResolver = &TNameResolver{}

func (c *TNameResolver) Resolve(config *ConfigParams) string {
	return c.ResolveWithDefault(config, "")
}

func (c *TNameResolver) ResolveWithDefault(config *ConfigParams, defaultName string) string {
	var name = config.GetAsString("name")

	if name == "" {
		name = config.GetAsString("id")
	}

	if name == "" {
		var descriptorStr = config.GetAsString("descriptor")
		if descriptorStr != "" {
			var descriptor = refer.NewDescriptorFromString(descriptorStr)
			if descriptor != nil {
				name = descriptor.Name()
			}
		}
	}

	if name == "" {
		name = defaultName
	}

	return name
}
